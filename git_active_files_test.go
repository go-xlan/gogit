package gogitv5acp

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
	root := runpath.PARENT.Path()

	worktree, err := NewWorktree(root)
	require.NoError(t, err)

	options := &GetActiveFilesOptions{
		Root:                root,
		IncludeDeletedFiles: false,
		FileExtension:       ".md", //we only need md files
		NoneExtension:       false,
		RunOnEachPath: func(path string) error {
			return nil
		},
	}

	activeFiles, err := GetActiveFiles(worktree, options)
	require.NoError(t, err)
	t.Log(neatjsons.S(activeFiles))
}

func TestGetActiveFiles_Execute_GoFormatFile(t *testing.T) {
	root := runpath.PARENT.Path()

	worktree, err := NewWorktree(root)
	require.NoError(t, err)

	options := NewGetActiveFilesOptions(root).
		SetFileExtension(".go"). //we only need go files
		SetRunOnEachPath(func(path string) error {
			t.Log("run_on_path:", path)

			// want skip some go file you can use this:
			if strings.HasSuffix(path, ".pb.go") || //skip the pb code
				strings.HasSuffix(path, "/wire_gen.go") || //skip the wire code
				strings.Contains(path, "/internal/data/ent/") { //skip the auto gen code
				return nil
			}

			// format go file content with go fmt operation
			if err := formatgo.FormatFile(path); err != nil {
				return erero.WithMessagef(err, "wrong path=%s", path)
			}
			t.Log("run format go file content success")
			return nil
		})

	activeFiles, err := GetActiveFiles(worktree, options)
	require.NoError(t, err)
	t.Log(neatjsons.S(activeFiles))
}
