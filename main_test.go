package vanity_test

import (
	"jordanreger.com/vanity"
	"testing"
)

func TestGit(t *testing.T) {
	url := "jordanreger.com/vanity"
	git_url := "git.sr.ht/~jordanreger/vanity"
	res := vanity.Git("/", url, git_url)

	t.Log(res)
}
