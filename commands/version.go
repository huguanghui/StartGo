package commands

import (
	"github.com/huguanghui/StartGo/ui"
	"github.com/huguanghui/StartGo/version"
)

var cmdVersion = &Command{
	Run:          runVersion,
	Usage:        "version",
	Long:         "shows git version and hub client version.",
	GitExtension: true,
}

func init() {
	CmdRunner.Use(cmdVersion, "--version")
}

func runVersion(cmd *Command, args *Args) {
	versionCmd := args.ToCmd()
	versionCmd.Spawn()
	ui.Printf("hub version %s\n", version.Version)
	args.NoForward()
}
