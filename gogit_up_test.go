package gogit_test

import (
	"testing"

	"github.com/go-xlan/gogit"
	"github.com/go-xlan/gogit/gogitassist"
	"github.com/stretchr/testify/require"
)

// TestClient_GetCurrentBranch verifies getting current branch name
// Should return branch name from temp repo
//
// TestClient_GetCurrentBranch 验证获取当前分支名称
// 应该从临时仓库返回分支名称
func TestClient_GetCurrentBranch(t *testing.T) {
	tempDIR := setupTestRepo(t)

	client, err := gogit.New(tempDIR)
	require.NoError(t, err)

	branch, err := client.GetCurrentBranch()
	require.NoError(t, err)
	require.NotEmpty(t, branch)
	t.Log("current branch:", branch)
}

// TestClient_GetLatestCommit verifies getting latest commit object
// Should return commit object with message and author info
//
// TestClient_GetLatestCommit 验证获取最新提交对象
// 应该返回包含消息和作者信息的提交对象
func TestClient_GetLatestCommit(t *testing.T) {
	tempDIR := setupTestRepo(t)

	client, err := gogit.New(tempDIR)
	require.NoError(t, err)

	commit, err := client.GetLatestCommit()
	require.NoError(t, err)
	require.NotNil(t, commit)
	require.NotEmpty(t, commit.Message)
	require.NotEmpty(t, commit.Author.Name)
	t.Log("latest commit:", commit.Hash.String()[:8], commit.Message)
}

// TestClient_HasChanges verifies uncommitted changes detection
// Should return boolean indicating whether changes exist
//
// TestClient_HasChanges 验证检查未提交更改
// 应该返回布尔值表示是否存在更改
func TestClient_HasChanges(t *testing.T) {
	tempDIR := setupTestRepo(t)

	client, err := gogit.New(tempDIR)
	require.NoError(t, err)

	hasChanges, err := client.HasChanges()
	require.NoError(t, err)
	require.False(t, hasChanges) // Fresh repo has no changes
	t.Log("has changes:", hasChanges)
}

// TestClient_GetRemoteURL verifies getting remote URL
// Should return URL matching the specified remote name
//
// TestClient_GetRemoteURL 验证获取远程 URL
// 应该返回与指定远程名称匹配的 URL
func TestClient_GetRemoteURL(t *testing.T) {
	tempDIR := setupTestRepo(t)

	client, err := gogit.New(tempDIR)
	require.NoError(t, err)

	// Add a remote to use in test validation
	// 添加远程用于测试验证
	require.NoError(t, gogitassist.AddRemote(client.Repo(), "origin", "https://github.com/example/repo.git"))

	remoteURL, err := client.GetRemoteURL("origin")
	require.NoError(t, err)
	require.Equal(t, "https://github.com/example/repo.git", remoteURL)
	t.Log("origin URL:", remoteURL)
}
