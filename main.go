package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/huguanghui/StartGo/commands"
	"github.com/huguanghui/StartGo/ui"
)

func main() {
	// os.args
	fmt.Println(reflect.TypeOf(os.Args))

	commands.StartCommands()

	n := ui.UiTest()
	fmt.Printf("%d\n", n)
}
