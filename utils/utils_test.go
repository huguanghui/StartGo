package utils

import (
	"github.com/bmizerany/assert"
	"testing"
	"time"
)

func TestSearchBrowserLauncher(t *testing.T) {
	browser := searchBrowserLauncher("darwin")
	assert.Equal(t, "open", browser)

	browser = searchBrowserLauncher("windows")
	assert.Equal(t, "cmd /c start", browser)
}

func TestConcatPaths(t *testing.T) {
	assert.Equal(t, "foo/bar/baz", ConcatPaths("foo", "bar", "baz"))
}

func TestTimeAgo(t *testing.T) {
	timeNow = func() time.Time {
		return time.Date(2018, 10, 28, 14, 34, 58, 651387237, time.UTC)
	}

	now := timeNow()

	secAgo := now.Add(-1 * time.Second)
	actual := TimeAgo(secAgo)
	assert.Equal(t, "now", actual)

	minAgo := now.Add(-1 * time.Minute)
	actual = TimeAgo(minAgo)
	assert.Equal(t, "1 minute ago", actual)

	minsAgo := now.Add(-5 * time.Minute)
	actual = TimeAgo(minsAgo)
	assert.Equal(t, "5 minutes ago", actual)

	hourAgo := now.Add(-1 * time.Hour)
	actual = TimeAgo(hourAgo)
	assert.Equal(t, "1 hour ago", actual)
}
