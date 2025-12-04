package gogitassist_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-xlan/gogit/gogitassist"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
)

// TestInitRepo verifies initializing a new git repo
// Should create .git directory and return valid repo
//
// TestInitRepo 验证初始化新的 git 仓库
// 应该创建 .git 目录并返回有效仓库
func TestInitRepo(t *testing.T) {
	tempDIR := rese.V1(os.MkdirTemp("", "gogit-init-test-*"))
	t.Cleanup(func() {
		must.Done(os.RemoveAll(tempDIR))
	})

	repo, err := gogitassist.InitRepo(tempDIR)
	require.NoError(t, err)
	require.NotNil(t, repo)

	// Verify .git directory exists
	// 验证 .git 目录存在
	gitDIR := filepath.Join(tempDIR, ".git")
	info, err := os.Stat(gitDIR)
	require.NoError(t, err)
	require.True(t, info.IsDir())
}

// TestSetConfigUserName verifies setting user.name in repo config
//
// TestSetConfigUserName 验证设置仓库配置中的 user.name
func TestSetConfigUserName(t *testing.T) {
	tempDIR := rese.V1(os.MkdirTemp("", "gogit-config-test-*"))
	t.Cleanup(func() {
		must.Done(os.RemoveAll(tempDIR))
	})

	repo := rese.P1(gogitassist.InitRepo(tempDIR))

	err := gogitassist.SetConfigUserName(repo, "Test Account")
	require.NoError(t, err)

	cfg, err := repo.Config()
	require.NoError(t, err)
	require.Equal(t, "Test Account", cfg.User.Name)
}

// TestSetConfigUserMailbox verifies setting user.email in repo config
//
// TestSetConfigUserMailbox 验证在仓库配置中设置 user.email 属性
func TestSetConfigUserMailbox(t *testing.T) {
	tempDIR := rese.V1(os.MkdirTemp("", "gogit-config-test-*"))
	t.Cleanup(func() {
		must.Done(os.RemoveAll(tempDIR))
	})

	repo := rese.P1(gogitassist.InitRepo(tempDIR))

	err := gogitassist.SetConfigUserMailbox(repo, "test@example.com")
	require.NoError(t, err)

	cfg, err := repo.Config()
	require.NoError(t, err)
	require.Equal(t, "test@example.com", cfg.User.Email)
}

// TestSetConfigUserInfo verifies setting both user.name and user.email attributes
//
// TestSetConfigUserInfo 验证同时设置 user.name 和 user.email 属性
func TestSetConfigUserInfo(t *testing.T) {
	tempDIR := rese.V1(os.MkdirTemp("", "gogit-config-test-*"))
	t.Cleanup(func() {
		must.Done(os.RemoveAll(tempDIR))
	})

	repo := rese.P1(gogitassist.InitRepo(tempDIR))

	err := gogitassist.SetConfigUserInfo(repo, "Test Account", "test@example.com")
	require.NoError(t, err)

	cfg, err := repo.Config()
	require.NoError(t, err)
	require.Equal(t, "Test Account", cfg.User.Name)
	require.Equal(t, "test@example.com", cfg.User.Email)
}

// TestAddRemote verifies adding a remote to repo
//
// TestAddRemote 验证向仓库添加远程
func TestAddRemote(t *testing.T) {
	tempDIR := rese.V1(os.MkdirTemp("", "gogit-remote-test-*"))
	t.Cleanup(func() {
		must.Done(os.RemoveAll(tempDIR))
	})

	repo := rese.P1(gogitassist.InitRepo(tempDIR))

	err := gogitassist.AddRemote(repo, "origin", "https://github.com/example/repo.git")
	require.NoError(t, err)

	remote, err := repo.Remote("origin")
	require.NoError(t, err)
	require.NotNil(t, remote)
	require.Contains(t, remote.Config().URLs, "https://github.com/example/repo.git")
}

// TestRemoveRemote verifies removing a remote from repo
//
// TestRemoveRemote 验证从仓库删除远程
func TestRemoveRemote(t *testing.T) {
	tempDIR := rese.V1(os.MkdirTemp("", "gogit-remote-test-*"))
	t.Cleanup(func() {
		must.Done(os.RemoveAll(tempDIR))
	})

	repo := rese.P1(gogitassist.InitRepo(tempDIR))

	// Add remote first
	// 先添加远程
	require.NoError(t, gogitassist.AddRemote(repo, "origin", "https://github.com/example/repo.git"))

	// Remove remote
	// 删除远程
	err := gogitassist.RemoveRemote(repo, "origin")
	require.NoError(t, err)

	// Verify remote no longer exists
	// 验证远程不再存在
	_, err = repo.Remote("origin")
	require.Error(t, err)
}

// TestCommit verifies creating a commit
//
// TestCommit 验证创建提交
func TestCommit(t *testing.T) {
	tempDIR := rese.V1(os.MkdirTemp("", "gogit-commit-test-*"))
	t.Cleanup(func() {
		must.Done(os.RemoveAll(tempDIR))
	})

	repo := rese.P1(gogitassist.InitRepo(tempDIR))

	// Create a file to commit
	// 创建要提交的文件
	testFile := filepath.Join(tempDIR, "README.md")
	must.Done(os.WriteFile(testFile, []byte("# Test\n"), 0644))

	hash, err := gogitassist.Commit(repo, "Base commit", "Test Account", "test@example.com")
	require.NoError(t, err)
	require.False(t, hash.IsZero())

	// Verify commit exists
	// 验证提交存在
	commit, err := repo.CommitObject(hash)
	require.NoError(t, err)
	require.Equal(t, "Base commit", commit.Message)
	require.Equal(t, "Test Account", commit.Author.Name)
	require.Equal(t, "test@example.com", commit.Author.Email)
}
