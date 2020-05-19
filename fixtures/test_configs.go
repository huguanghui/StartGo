package fixtures

import (
	"os"
)

type TestConfigs struct {
	Path string
}

func (c *TestConfigs) TearDown() {
	os.Setenv("HUB_CONFIG", "")
	os.RemoveAll(c.Path)
}
