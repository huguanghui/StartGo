package commands

import (
	//	"fmt"
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
