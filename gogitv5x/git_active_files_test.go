package gogitv5x

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestGetActiveFiles(t *testing.T) {
	root := runpath.PARENT.Up(1)

	worktree, err := NewWorktreeWithIgnore(root)
	require.NoError(t, err)

	options := &GetActiveFilesOptions{
		Root:                root,
		IncludeDeletedFiles: false,
		FileExtension:       ".md", //we only need md files
		NoneExtension:       false,
		MatchPathFunc: func(path string) bool {
			return true
		},
		RunOnEachPath: func(path string) error {
			return nil
		},
	}

	activeFiles, err := GetActiveFiles(worktree, options)
	require.NoError(t, err)
	t.Log(neatjsons.S(activeFiles))
}

func TestGetActiveFiles_Execute_GoFormatFile(t *testing.T) {
	root := runpath.PARENT.Up(1)

	worktree, err := NewWorktreeWithIgnore(root)
	require.NoError(t, err)

	options := NewGetActiveFilesOptions(root).
		SetFileExtension(".go"). //we only need go files
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
		}).
		SetRunOnEachPath(func(path string) error {
			t.Log("run_on_path:", path)

			// format go file content with go fmt operation
			if err := formatgo.FormatFile(path); err != nil {
				return erero.WithMessagef(err, "wrong path=%s", path)
			}
			t.Log("run format content success")
			return nil
		})

	activeFiles, err := GetActiveFiles(worktree, options)
	require.NoError(t, err)
	t.Log(neatjsons.S(activeFiles))
}

func TestGetActiveFilesOptions_SetRunOnFilePath(t *testing.T) {
	root := runpath.PARENT.Up(1)

	worktree, err := NewWorktreeWithIgnore(root)
	require.NoError(t, err)

	activeFiles, err := GetActiveFiles(worktree, NewGetActiveFilesOptions(root).SetRunOnFilePath(func(path string) {
		t.Log("run_on_path:", path)
	}))
	require.NoError(t, err)
	t.Log(neatjsons.S(activeFiles))
}
