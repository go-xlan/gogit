package gogitv5acp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestNewFmtActiveFilesOptions(t *testing.T) {
	root := runpath.PARENT.Path()

	worktree, err := NewWorktreeWithIgnore(root)
	require.NoError(t, err)

	options := NewFmtActiveFilesOptions(root).
		SetMatchPathFunc(func(path string) bool {
			t.Log("match_path:", path)

			if strings.HasSuffix(path, ".pb.go") || //skip the pb code
				strings.HasSuffix(path, "/wire_gen.go") || //skip the wire code
				strings.Contains(path, "/internal/data/ent/") { //skip the auto gen code
				t.Log("skip")
				return false
			}

			t.Log("pass")
			return true
		})

	activeFiles, err := GetActiveFiles(worktree, options)
	require.NoError(t, err)
	t.Log(neatjsons.S(activeFiles))
}
