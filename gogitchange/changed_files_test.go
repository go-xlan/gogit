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

// TestListChangedFilePaths_Markdown tests listing changed Markdown files
// Verifies file path collection for .md files using type matching
//
// TestListChangedFilePaths_Markdown 测试列出变更的 Markdown 文件
// 使用类型匹配验证 .md 文件的文件路径收集
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

// TestListChangedFilePaths_Golang tests listing changed Go source files
// Verifies file path collection for .go files using type matching
//
// TestListChangedFilePaths_Golang 测试列出变更的 Go 源文件
// 使用类型匹配验证 .go 文件的文件路径收集
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

// TestForeachChangedGoFile tests iteration through changed Go files
// Verifies callback execution for each changed .go file
//
// TestForeachChangedGoFile 测试遍历变更的 Go 文件
// 验证对每个变更的 .go 文件执行回调
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

// TestFormatChangedGoFiles tests Go file formatting with advanced filtering
// Verifies formatgo integration with custom path matching and exclusion patterns
//
// TestFormatChangedGoFiles 测试带高级过滤的 Go 文件格式化
// 验证 formatgo 集成以及自定义路径匹配和排除模式
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
			strings.Contains(path, "/generated/data/ent/") { //skip the auto gen code
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
