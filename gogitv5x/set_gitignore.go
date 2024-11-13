package gogitv5x

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

func SetIgnorePatterns(worktree *git.Worktree, patterns []gitignore.Pattern) {
	worktree.Excludes = append(worktree.Excludes, patterns...)
}
