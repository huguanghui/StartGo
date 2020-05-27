package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/huguanghui/StartGo/commands"
	"github.com/huguanghui/StartGo/ui"
	"github.com/huguanghui/StartGo/utils"
)

func main() {

	fmt.Println("Args: ", os.Args)
	err := commands.CmdRunner.Execute(os.Args)
	ui.Println("=========================")
	exitCode := HandleError(err)

	ui.Println("Hello World!")
	pathdir := utils.ConcatPaths("abc", "bbdd")
	fmt.Println(pathdir)
	fmt.Println(runtime.GOOS)

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
