package git

import (
	"fmt"
	"strings"

	"github.com/huguanghui/StartGo/cmd"
)

var GlobalFlags []string

func Version() (string, error) {
	output, err := gitOutput("version")
	if err == nil {
		return output[0], nil
	} else {
		return "", fmt.Errorf("error running git version: %s", err)
	}
}

var cachedDir string

func gitOutput(input ...string) (outputs []string, err error) {
	cmd := gitCmd(input...)

	out, err := cmd.CombinedOutput()
	for _, line := range strings.Split(out, "\n") {
		if strings.TrimSpace(line) != "" {
			outputs = append(outputs, string(line))
		}
	}

	return outputs, err
}

func gitCmd(args ...string) *cmd.Cmd {
	cmd := cmd.New("git")

	for _, v := range GlobalFlags {
		cmd.WithArg(v)
	}

	for _, a := range args {
		cmd.WithArg(a)
	}

	return cmd
}
