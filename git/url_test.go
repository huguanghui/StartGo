package git

import (
	"testing"

	"github.com/huguanghui/StartGo/internal/assert"
)

func createURLParser() *URLParser {
	c := make(SSHConfig)
	c["github.com"] = "ssh.github.com"
	c["gh"] = "github.com"
	c["git.company.com"] = "ssh.git.company.com"

	return &URLParser{c}
}

func TestURLParser_ParserURL_HTTPURL(t *testing.T) {
	p := createURLParser()

	u, err := p.Parse("https://github.com/octokit/go-octokit.git")
	assert.Equal(t, nil, err)
	assert.Equal(t, "github.com", u.Host)
	assert.Equal(t, "https", u.Scheme)
	assert.Equal(t, "/octokit/go-octokit.git", u.Path)
}
