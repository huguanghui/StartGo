package commands

import (
	"fmt"
)

var (
	NameRe = `[\w.-]+`

	CmdRunner = NewRunner()
)

func StartCommands() {
	fmt.Println("hello commands")
}

type Command struct {
	Usage        string
	GitExtension bool
}
