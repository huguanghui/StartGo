package commands

import (
	"fmt"
	"strings"
)

type Runner struct {
	commands map[string]*Command
}

func NewRunner() *Runner {
	return &Runner{
		commands: make(map[string]*Command),
	}
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
	if forceFail {
		fmt.Println("CmdName:", cmdName)
	}

	return nil
}
