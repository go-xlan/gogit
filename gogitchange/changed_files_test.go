package gogitchange_test

import (
	"strings"
	"testing"

	"github.com/go-xlan/gogit/gogitassist"
	"github.com/go-xlan/gogit/gogitchange"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestListChangedFilePaths_Markdown(t *testing.T) {
	root := runpath.PARENT.Up(1)

	repo, tree, err := gogitassist.NewRepoTreeWithIgnore(root)
	require.NoError(t, err)
	gogitassist.DebugRepo(repo)

	manager := gogitchange.NewChangedFileManager(root, tree)
	options := gogitchange.NewMatchOptions().MatchType(".md")
	paths, err := manager.ListChangedFilePaths(options)
	require.NoError(t, err)
	t.Log(neatjsons.S(paths))
}

func TestListChangedFilePaths_Golang(t *testing.T) {
	root := runpath.PARENT.Up(1)

	repo, tree, err := gogitassist.NewRepoTreeWithIgnore(root)
	require.NoError(t, err)
	gogitassist.DebugRepo(repo)

	manager := gogitchange.NewChangedFileManager(root, tree)
	options := gogitchange.NewMatchOptions().MatchType(".go")
	paths, err := manager.ListChangedFilePaths(options)
	require.NoError(t, err)
	t.Log(neatjsons.S(paths))
}

func TestForeachChangedGoFile(t *testing.T) {
	projectRoot := runpath.PARENT.Up(1)

	repo, tree, err := gogitassist.NewRepoTreeWithIgnore(projectRoot)
	require.NoError(t, err)
	gogitassist.DebugRepo(repo)

	manager := gogitchange.NewChangedFileManager(projectRoot, tree)
	options := gogitchange.NewMatchOptions().MatchType(".go")
	require.NoError(t, manager.ForeachChangedGoFile(options, func(path string) error {
		t.Log("path:", path)
		return nil
	}))
}

func TestFormatChangedGoFiles(t *testing.T) {
	projectRoot := runpath.PARENT.Up(1)

	repo, tree, err := gogitassist.NewRepoTreeWithIgnore(projectRoot)
	require.NoError(t, err)
	gogitassist.DebugRepo(repo)

	manager := gogitchange.NewChangedFileManager(projectRoot, tree)
	options := gogitchange.NewMatchOptions().MatchType(".go").MatchPath(func(path string) bool {
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
	require.NoError(t, manager.ForeachChangedGoFile(options, func(path string) error {
		t.Log("golang-format-source-file-path:", path)

		if err := formatgo.FormatFile(path); err != nil {
			return erero.Wro(err)
		}
		return nil
	}))
}
