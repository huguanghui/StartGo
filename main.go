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
	// os.args
	//fmt.Println(reflect.TypeOf(os.Args))

	commands.CmdRunner.Execute(os.Args)

	pathdir := utils.ConcatPaths("abc", "bbdd")
	fmt.Println(pathdir)
	fmt.Println(runtime.GOOS)

	n := ui.UiTest()
	fmt.Printf("%d\n", n)

}
