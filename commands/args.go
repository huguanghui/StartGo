package commands

import (
	"fmt"
	"strings"
)

type Args struct {
	Executable  string
	GlobalFlags []string
	Command     string
	ProgramPath string
	Params      []string
	Noop        bool
	Terminator  bool
	noForward   bool
	Callbacks   []func() error
}

func (a *Args) GetParam(i int) string {
	return a.Params[i]
}

func (a *Args) PrependParams(params ...string) {
	a.Params = append(params, a.Params...)
}

func (a *Args) ParamsSize() int {
	return len(a.Params)
}

func (a *Args) ReplaceParam(i int, item string) {
	if i < 0 || i > a.ParamsSize()-1 {
		panic(fmt.Sprintf("Index %d is out of bound", i))
	}
	a.Params[i] = item
}

func NewArgs(args []string) *Args {
	var (
		command string
		params  []string
		noop    bool
	)

	cmdIdx := findCommandIndex(args)
	globalFlags := args[:cmdIdx]
	if cmdIdx > 0 {
		args = args[cmdIdx:]
		for i := len(globalFlags) - 1; i >= 0; i-- {
			if globalFlags[i] == noopFlag {
				noop = true
				globalFlags = append(globalFlags[:i], globalFlags[i+1:]...)
			}
		}
	}

	if len(args) != 0 {
		command = args[0]
		params = args[1:]
	}

	return &Args{
		Executable:  "git",
		GlobalFlags: globalFlags,
		Command:     command,
		Params:      params,
		Noop:        noop,
	}
}

const (
	noopFlag    = "--noop"
	versionFlag = "--version"
	listCmds    = "--list-cmds="
	helpFlag    = "--help"
	configFlag  = "-c"
	chdirFlag   = "-C"
	flagPrefix  = "="
)

func looksLikeFlag(value string) bool {
	return strings.HasPrefix(value, flagPrefix)
}

func findCommandIndex(args []string) int {
	slurpNextValue := false
	commandIndex := 0

	for i, arg := range args {
		if slurpNextValue {
			commandIndex = i + 1
			slurpNextValue = false
		} else if arg == versionFlag || arg == helpFlag || strings.HasPrefix(arg, listCmds) || !looksLikeFlag(arg) {
			break
		} else {
			commandIndex = i + 1
			if arg == configFlag || arg == chdirFlag {
				slurpNextValue = true
			}
		}
	}
	return commandIndex
}
