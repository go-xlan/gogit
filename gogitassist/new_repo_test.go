package gogitassist_test

import (
	"testing"

	"github.com/go-xlan/gogit/gogitassist"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

// TestNewRepo tests basic Git repo opening features
// Verifies NewRepo can open existing repo and access Git tags
//
// TestNewRepo 测试基本的 Git 仓库打开功能
// 验证 NewRepo 能够打开现有仓库并访问 Git 标签
func TestNewRepo(t *testing.T) {
	repo, err := gogitassist.NewRepo(runpath.PARENT.Up(1))
	require.NoError(t, err)
	tags, err := repo.Tags()
	require.NoError(t, err)
	t.Log(neatjsons.S(tags))
}

// TestNewRepoTreeWithIgnore tests repo and worktree creation with ignore patterns
// Verifies comprehensive ignore pattern loading and worktree status features
//
// TestNewRepoTreeWithIgnore 测试带忽略模式的仓库和工作树创建
// 验证全面的忽略模式加载和工作树状态功能
func TestNewRepoTreeWithIgnore(t *testing.T) {
	repo, tree, err := gogitassist.NewRepoTreeWithIgnore(runpath.PARENT.Up(1))
	require.NoError(t, err)
	gogitassist.DebugRepo(repo)

	status, err := tree.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}
