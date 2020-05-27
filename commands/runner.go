package commands

import (
	"fmt"
	"strings"

	"github.com/huguanghui/StartGo/git"
	"github.com/kballard/go-shellquote"
)

type Runner struct {
	commands map[string]*Command
}

func NewRunner() *Runner {
	return &Runner{
		commands: make(map[string]*Command),
	}
}

func (r *Runner) Lookup(name string) *Command {
	return r.commands[name]
}

func (r *Runner) All() map[string]*Command {
	return r.commands
}

func (r *Runner) Execute(cliArgs []string) error {
	args := NewArgs(cliArgs[1:])
	args.ProgramPath = cliArgs[0]
	forceFail := false

	if args.Command == "" && len(args.GlobalFlags) == 0 {
		args.Command = "help"
		forceFail = true
	}

	cmdName := args.Command
	if strings.Contains(cmdName, "=") {
		cmdName = strings.SplitN(cmdName, "=", 2)[0]
	}

	git.GlobalFlags = args.GlobalFlags
	if !isBuiltInHubCommand(cmdName) {
		expandAlias(args)
		cmdName = args.Command
	}

	// 使 <cmd> --help 等价于 help <cmd>
	if args.ParamsSize() == 1 && args.GetParam(0) == helpFlag {
		if c := r.Lookup(cmdName); c != nil && !c.GitExtension {
			args.ReplaceParam(0, cmdName)
			args.Command = "help"
			cmdName = args.Command
		}
	}

	if forceFail {
		fmt.Println("CmdName:", cmdName)
	}

	return nil
}

func expandAlias(args *Args) {
	cmd := args.Command
	if cmd == "" {
		return
	}
	expandCmd, err := git.Alias(cmd)
	if err == nil && expandCmd != "" && !git.IsBuiltInGitCommand(cmd) {
		words, e := splitAliasCmd(expandCmd)
		if e == nil {
			args.Command = words[0]
			args.PrependParams(words[1:]...)
		}
	}
}

func isBuiltInHubCommand(command string) bool {
	for hubCommand := range CmdRunner.All() {
		if hubCommand == command {
			return true
		}
	}
	return false
}

func splitAliasCmd(cmd string) ([]string, error) {
	if cmd == "" {
		return nil, fmt.Errorf("alias can't be empty'")
	}

	if strings.HasPrefix(cmd, "!") {
		return nil, fmt.Errorf("alias starting with ! can't be split")
	}

	words, err := shellquote.Split(cmd)
	if err != nil {
		return nil, err
	}

	return words, nil
}
