package assert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Equal(t testing.TB, want, got interface{}, args ...interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		msg := fmt.Sprint(args...)
		t.Errorf("%s\n%s", msg, cmp.Diff(want, got))
	}
}

func NotEqual(t testing.TB, want, got interface{}, args ...interface{}) {
	t.Helper()
	if reflect.DeepEqual(want, got) {
		msg := fmt.Sprint(args...)
		t.Errorf("%s\nUnexpected: <%#v>", msg, want)
	}
}

func T(t testing.TB, ok bool, args ...interface{}) {
	t.Helper()
	if !ok {
		msg := fmt.Sprint(args...)
		t.Errorf("%s\nFailure", msg)
	}
}
