package main

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/huguanghui/StartGo/commands"
	"github.com/huguanghui/StartGo/ui"
)

func main() {

	err := commands.CmdRunner.Execute(os.Args)
	ui.Println("=========================")
	exitCode := HandleError(err)

	os.Exit(exitCode)
}

func HandleError(err error) int {
	if err == nil {
		return 0
	}
	switch e := err.(type) {
	case *exec.ExitError:
		if status, ok := e.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus()
		} else {
			return 1
		}
	default:
		if errString := err.Error(); errString != "" {
			ui.Errorln(err)
		}
		return 1
	}
}
