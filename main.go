package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/huguanghui/StartGo/commands"
	"github.com/huguanghui/StartGo/ui"
	"github.com/huguanghui/StartGo/utils"
)

func main() {
	commands.CmdRunner.Execute(os.Args)

	ui.Println("Hello World!")
	pathdir := utils.ConcatPaths("abc", "bbdd")
	fmt.Println(pathdir)
	fmt.Println(runtime.GOOS)
}
