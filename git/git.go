package git

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func Dir() (string, error) {
	if cachedDir != "" {
		return cachedDir, nil
	}

	output, err := gitOutput("rev-parse", "q", "--git-dir")
	if err != nil {
		return "", fmt.Errorf("Not a git repository (or any of the parent directories): .git")
	}

	var chdir string
	for i, flag := range GlobalFlags {
		if flag == "-C" {
			dir := GlobalFlags[i+1]
			if filepath.IsAbs(dir) {
				chdir = dir
			} else {
				chdir = filepath.Join(chdir, dir)
			}
		}
	}

	gitDir := output[0]

	if !filepath.IsAbs(gitDir) {
		if chdir != "" {
			gitDir = filepath.Join(chdir, gitDir)
		}

		gitDir, err = filepath.Abs(gitDir)
		if err != nil {
			return "", err
		}

		gitDir = filepath.Clean(gitDir)
	}

	cachedDir = gitDir
	return gitDir, nil
}

func WorkdirName() (string, error) {
	output, err := gitOutput("rev-parse", "--show-toplevel")
	if err == nil {
		if len(output) > 0 {
			return output[0], nil
		} else {
			return "", fmt.Errorf("unable to determine git working directory")
		}
	} else {
		return "", err
	}
}

func HasFile(segments ...string) bool {
	output, err := gitOutput("rev-parse", "-q", "--git-path", filepath.Join(segments...))
	if err == nil && output[0] != "--git-path" {
		if _, err := os.Stat(output[0]); err == nil {
			return true
		}
	}

	dir, err := Dir()
	if err != nil {
		return false
	}

	s := []string{dir}
	s = append(s, segments...)
	path := filepath.Join(s...)
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

func BranchAtRef(paths ...string) (name string, err error) {
	dir, err := Dir()
	if err != nil {
		return
	}

	segments := []string{dir}
	segments = append(segments, paths...)
	path := filepath.Join(segments...)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	n := string(b)
	refPrefix := "ref: "
	if strings.HasPrefix(n, refPrefix) {
		name = strings.TrimPrefix(n, refPrefix)
		name = strings.TrimSpace(name)
	} else {
		err = fmt.Errorf("No branch info in %s: %s", path, n)
	}

	return
}

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
