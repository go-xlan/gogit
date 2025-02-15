package gogitassist

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"github.com/yyle88/erero"
	"github.com/yyle88/osexistpath/osomitexist"
)

func SetIgnorePatterns(worktree *git.Worktree, patterns []gitignore.Pattern) {
	worktree.Excludes = append(worktree.Excludes, patterns...)
}

func LoadProjectIgnorePatterns(root string) ([]gitignore.Pattern, error) {
	if osomitexist.IsRoot(root) {
		patterns, err := LoadIgnorePatternsFromPath(filepath.Join(root, ".gitignore"))
		if err != nil {
			return nil, erero.Wro(err)
		}
		return patterns, nil
	}
	return []gitignore.Pattern{}, nil
}

func LoadIgnorePatternsFromPath(path string) ([]gitignore.Pattern, error) {
	if osomitexist.IsFile(path) {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, erero.Wro(err)
		}
		patterns, err := LoadIgnorePatternsFromText(string(data))
		if err != nil {
			return nil, erero.Wro(err)
		}
		return patterns, nil
	}
	return []gitignore.Pattern{}, nil
}

func LoadIgnorePatternsFromText(text string) ([]gitignore.Pattern, error) {
	var patterns = make([]gitignore.Pattern, 0)
	for _, stx := range strings.Split(text, "\n") {
		if stx = strings.TrimSpace(stx); stx != "" && !strings.HasPrefix(stx, "#") {
			patterns = append(patterns, gitignore.ParsePattern(stx, []string{}))
		}
	}
	return patterns, nil
}
