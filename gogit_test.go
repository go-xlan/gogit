package gogit_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-xlan/gogit"
	"github.com/go-xlan/gogit/gogitassist"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
)

// TestNew verifies client creation from temp repo
// Tests basic client creation and status checking features
//
// TestNew 验证从临时仓库创建客户端
// 测试基本客户端创建和状态检查功能
func TestNew(t *testing.T) {
	tempDIR := setupTestRepo(t)

	client, err := gogit.New(tempDIR)
	require.NoError(t, err)
	gogitassist.DebugRepo(client.Repo())

	status, err := client.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}

// TestNewClient tests creation using rese patterns as convenience wrappers
// Verifies that rese.P1 provides consistent handling and panic-on-issue semantics
//
// TestNewClient 使用 rese 辅助模式测试客户端创建
// 验证 rese.P1 提供一致的处理和遇错即崩语义
func TestNewClient(t *testing.T) {
	tempDIR := setupTestRepo(t)

	client := rese.P1(gogit.New(tempDIR))
	gogitassist.DebugRepo(client.Repo())

	status, err := client.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}

// TestClient_IsLatestCommitPushedToRemote tests remote push status detection
// Verifies the function that checks if HEAD commit exists in specified remote
// Uses production repo since it needs a functioning remote connection
//
// TestClient_IsLatestCommitPushedToRemote 测试远程推送状态检测
// 验证检查 HEAD 提交是否存在于指定远程的函数
// 使用真实仓库因为需要有效的远程连接
func TestClient_IsLatestCommitPushedToRemote(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)

	matched, err := client.IsLatestCommitPushedToRemote("origin")
	require.NoError(t, err)
	t.Log(matched)
}

// TestClient_IsLatestCommitPushedToRemote_NotExist tests response with non-existent remote
// Verifies that the function returns false when remote does not exist
//
// TestClient_IsLatestCommitPushedToRemote_NotExist 测试不存在远程的行为
// 验证远程不存在时函数返回 false
func TestClient_IsLatestCommitPushedToRemote_NotExist(t *testing.T) {
	tempDIR := setupTestRepo(t)

	client, err := gogit.New(tempDIR)
	require.NoError(t, err)

	matched, err := client.IsLatestCommitPushedToRemote("origin2")
	require.NoError(t, err)
	t.Log(matched)
	require.False(t, matched)
}

// TestClient_IsLatestCommitPushed tests push status across configured remotes
// Verifies detection of commit push state when checking against all remotes
// Uses production repo since it needs a functioning remote connection
//
// TestClient_IsLatestCommitPushed 测试所有配置远程的推送状态
// 验证针对所有远程检查时的提交推送状态检测
// 使用真实仓库因为需要有效的远程连接
func TestClient_IsLatestCommitPushed(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)

	pushed, err := client.IsLatestCommitPushed()
	require.NoError(t, err)
	t.Log(pushed)
}

// setupTestRepo creates a temp git repo with t.Cleanup enabling auto-dispose
// Environment setup must succeed, so we use rese/must to handle all tasks
// Returns the temp DIR path
//
// setupTestRepo 使用 t.Cleanup 创建临时 git 仓库以便自动删除
// 环境设置必须成功，因此我们使用 rese/must 处理所有任务
// 返回临时 DIR 路径
func setupTestRepo(t *testing.T) string {
	// Create temp DIR - must succeed
	// 创建临时 DIR - 必须成功
	tempDIR := rese.V1(os.MkdirTemp("", "gogit-test-*"))

	// Set up t.Cleanup to auto-dispose temp resources when test ends
	// 设置 t.Cleanup 在测试结束时自动清理临时资源
	t.Cleanup(func() {
		must.Done(os.RemoveAll(tempDIR))
	})

	// Initialize git repo using gogitassist - must succeed
	// 使用 gogitassist 初始化 git 仓库 - 必须成功
	repo := rese.P1(gogitassist.InitRepo(tempDIR))

	// Create base commit making it a functioning repo - must succeed
	// 创建基础提交使其成为有效仓库 - 必须成功
	testFile := filepath.Join(tempDIR, "README.md")
	must.Done(os.WriteFile(testFile, []byte("# Test Project\n"), 0644))

	// Create base commit using gogitassist - must succeed
	// 使用 gogitassist 创建基础提交 - 必须成功
	rese.V1(gogitassist.Commit(repo, "Base commit", "Test Account", "test@example.com"))

	return tempDIR
}

// TestClient_CommitAll_EmptyCommit verifies action when no changes exist
// Should return blank string without fault when commits are vacant
//
// TestClient_CommitAll_EmptyCommit 验证没有更改时的行为
// 对于空提交应该返回空字符串而不报错
func TestClient_CommitAll_EmptyCommit(t *testing.T) {
	tempDIR := setupTestRepo(t)

	client, err := gogit.New(tempDIR)
	require.NoError(t, err)
	require.NotNil(t, client)

	// Try to commit when no changes exist
	// 尝试在没有更改时提交
	commitInfo := &gogit.CommitInfo{
		Name:    "Test Account",
		Mailbox: "test@example.com",
		Message: "Vacant commit attempt",
	}

	hash, err := client.CommitAll(commitInfo)
	require.NoError(t, err)
	require.Empty(t, hash) // Should return blank string when commit is vacant
}

// TestClient_CommitAll_WithNewFile verifies committing a new file
// Should commit with success and return hash
//
// TestClient_CommitAll_WithNewFile 验证提交新文件
// 应该成功提交并返回哈希
func TestClient_CommitAll_WithNewFile(t *testing.T) {
	tempDIR := setupTestRepo(t)

	client, err := gogit.New(tempDIR)
	require.NoError(t, err)
	require.NotNil(t, client)

	// Create a new file - test setup
	// 创建新文件 - 测试设置
	newFile := filepath.Join(tempDIR, "test.txt")
	require.NoError(t, os.WriteFile(newFile, []byte("test content"), 0644))

	// Add new file to staging area first - new files need to be added with explicit intent
	// 先将新文件添加到暂存区 - 新文件需要显式添加
	require.NoError(t, client.AddAll())

	commitInfo := &gogit.CommitInfo{
		Name:    "Test Account",
		Mailbox: "test@example.com",
		Message: "Add test file",
	}

	hash, err := client.CommitAll(commitInfo)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	// Check file was committed - test validation uses require
	// 验证文件已提交 - 测试验证使用 require
	status, err := client.Status()
	require.NoError(t, err)
	require.NotNil(t, status)
	require.Empty(t, status)
}

// TestClient_CommitAll_WithModifiedFile verifies committing modifications
// Should commit changes to existing file with success
//
// TestClient_CommitAll_WithModifiedFile 验证提交修改
// 应该成功提交对现有文件的更改
func TestClient_CommitAll_WithModifiedFile(t *testing.T) {
	tempDIR := setupTestRepo(t)

	client, err := gogit.New(tempDIR)
	require.NoError(t, err)
	require.NotNil(t, client)

	// Change the README file - test setup
	// 修改 README 文件 - 测试设置
	readmeFile := filepath.Join(tempDIR, "README.md")
	require.NoError(t, os.WriteFile(readmeFile, []byte("# Updated Test Project\nNew content\n"), 0644))

	commitInfo := &gogit.CommitInfo{
		Name:    "Test Account",
		Mailbox: "test@example.com",
		Message: "Update README",
	}

	hash, err := client.CommitAll(commitInfo)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	// Check changes were committed
	// 验证更改已提交
	status, err := client.Status()
	require.NoError(t, err)
	require.NotNil(t, status)
	require.Empty(t, status)
}

// TestClient_AmendCommit_Success verifies amending previous commit
// Should amend with success when not pushed
//
// TestClient_AmendCommit_Success 验证修改先前的提交
// 未推送时应该成功修改
func TestClient_AmendCommit_Success(t *testing.T) {
	tempDIR := setupTestRepo(t)

	client, err := gogit.New(tempDIR)
	require.NoError(t, err)
	require.NotNil(t, client)

	// Create and commit a file first - test setup
	// 首先创建并提交文件 - 测试设置
	testFile := filepath.Join(tempDIR, "amend-test.txt")
	require.NoError(t, os.WriteFile(testFile, []byte("first content"), 0644))

	// Add new file to staging area first
	// 先将新文件添加到暂存区
	require.NoError(t, client.AddAll())

	firstCommit := &gogit.CommitInfo{
		Name:    "First Person",
		Mailbox: "first@example.com",
		Message: "First commit message",
	}
	firstHash, err := client.CommitAll(firstCommit)
	require.NoError(t, err)
	require.NotEmpty(t, firstHash)

	// Change the file to facilitate amendment - test setup
	// 修改文件以便修正 - 测试设置
	require.NoError(t, os.WriteFile(testFile, []byte("amended content"), 0644))
	require.NoError(t, client.AddAll())

	amendConfig := &gogit.AmendConfig{
		CommitInfo: &gogit.CommitInfo{
			Name:    "Amended Person",
			Mailbox: "amended@example.com",
			Message: "Amended commit message",
		},
		ForceAmend: false,
	}

	amendedHash, err := client.AmendCommit(amendConfig)
	require.NoError(t, err)
	require.NotEmpty(t, amendedHash)

	// Check amendment succeeded using go-git API
	// 使用 go-git API 验证修正成功
	repo, err := git.PlainOpen(tempDIR)
	require.NoError(t, err)
	require.NotNil(t, repo)
	headRef, err := repo.Head()
	require.NoError(t, err)
	require.NotNil(t, headRef)
	commitObj, err := repo.CommitObject(headRef.Hash())
	require.NoError(t, err)
	require.NotNil(t, commitObj)

	require.Equal(t, "Amended commit message", commitObj.Message)
	require.Equal(t, "Amended Person", commitObj.Author.Name)
}
