package gogitv5ops

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestCollectChangedFiles_Markdown(t *testing.T) {
	root := runpath.PARENT.Up(1)

	repository, worktree, err := NewRepoWorktreeWithIgnore(root)
	require.NoError(t, err)

	checkRepoHeadReference(t, repository)

	options := &ProcessingOptions{
		projectRoot:        root,
		matchWithExtension: ".md", //we only need md files
		matchNoneExtension: false,
		matchPath: func(path string) bool {
			return true
		},
	}

	changedFiles, err := options.CollectChangedFiles(worktree)
	require.NoError(t, err)
	t.Log(neatjsons.S(changedFiles))
}

func TestCollectChangedFiles_Golang(t *testing.T) {
	root := runpath.PARENT.Up(1)

	repository, worktree, err := NewRepoWorktreeWithIgnore(root)
	require.NoError(t, err)

	checkRepoHeadReference(t, repository)

	options := NewProcessingOptions(root).WithFileExtension(".go")
	changedFiles, err := options.CollectChangedFiles(worktree)
	require.NoError(t, err)
	t.Log(neatjsons.S(changedFiles))
}

func TestFormatModifiedGoFiles(t *testing.T) {
	projectRoot := runpath.PARENT.Up(1)

	repository, worktree, err := NewRepoWorktreeWithIgnore(projectRoot)
	require.NoError(t, err)

	checkRepoHeadReference(t, repository)

	options := NewProcessingOptions(projectRoot).
		WithFileExtension(".go").
		MatchPath(func(path string) bool {
			t.Log("path:", path)

			if strings.HasSuffix(path, ".pb.go") || //skip the pb code
				strings.HasSuffix(path, "/wire_gen.go") || //skip the wire code
				strings.Contains(path, "/internal/data/ent/") { //skip the auto gen code
				t.Log("skip:", path)
				return false
			}

			t.Log("pass:", path)
			return true
		})

	require.NoError(t, options.FormatModifiedGoFiles(worktree))
}
