package vanity_test

import (
	"jordanreger.com/web/vanity"
	"testing"
)

func TestGit(t *testing.T) {
	url := "jordanreger.com/bsky"
	git_url := "git.sr.ht/~jordanreger/bsky"
	res := vanity.Git("/", url, git_url)

	t.Log(res)
}
