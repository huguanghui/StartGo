package utils

import (
	"reflect"
	"testing"
)

func equal(t *testing.T, expected, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected: %#v, got: %#v", expected, got)
	}
}

func TestArgsParser(t *testing.T) {
	p := NewArgsParser()
	p.RegisterValue("--hello", "-e")
	p.RegisterValue("--origin", "-o")
	args := []string{"--hello", "world", "one", "--", "--two"}
	rest, err := p.Parse(args)
	equal(t, nil, err)
	equal(t, []string{"one", "--two"}, rest)
	equal(t, "world", p.Value("--hello"))
	equal(t, true, p.HasReceived("--hello"))
	equal(t, "", p.Value("-e"))
	equal(t, false, p.HasReceived("-e"))
	equal(t, "", p.Value("--origin"))
	equal(t, false, p.HasReceived("--origin"))
	equal(t, []int{2, 4}, p.PositionalIndices)
}

func TestArgsParser_RepeatedInvocation(t *testing.T) {
	p := NewArgsParser()
	p.RegisterValue("--hello", "-e")
	p.RegisterValue("--origin", "-o")

	rest, err := p.Parse([]string{"--hello", "world", "--", "one"})
	equal(t, nil, err)
	equal(t, []string{"one"}, rest)
	equal(t, []int{3}, p.PositionalIndices)
	equal(t, true, p.HasReceived("--hello"))
	equal(t, "world", p.Value("--hello"))
	equal(t, false, p.HasReceived("--origin"))
	equal(t, true, p.HasTerminated)

	rest, err = p.Parse([]string{"two", "-oupstream"})
	equal(t, nil, err)
	equal(t, []string{"two"}, rest)
	equal(t, []int{0}, p.PositionalIndices)
	equal(t, false, p.HasReceived("--hello"))
	equal(t, true, p.HasReceived("--origin"))
	equal(t, "upstream", p.Value("--origin"))
	equal(t, false, p.HasTerminated)

}
