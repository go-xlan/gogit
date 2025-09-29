package gogit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

// TestPackageName verifies package name constant matches current package
// Ensures package naming alignment for commit message generation
//
// TestPackageName 验证包名常量与当前包匹配
// 确保包命名一致性以生成提交消息
func TestPackageName(t *testing.T) {
	require.Equal(t, packageName, syntaxgo.CurrentPackageName())
}

// TestPackagePath verifies package path constant matches current import path
// Ensures package path precision for default commit message generation
//
// TestPackagePath 验证包路径常量与实际导入路径匹配
// 确保包路径准确性以生成默认提交消息
func TestPackagePath(t *testing.T) {
	require.Equal(t, packagePath, syntaxgo_reflect.GetPkgPathV4(&CommitInfo{}))
}

// TestCommitInfo_GetCommitMessage tests basic commit message extraction
// Verifies BuildCommitMessage returns provided message as expected
//
// TestCommitInfo_GetCommitMessage 测试基本的提交消息获取
// 验证 BuildCommitMessage 正确返回提供的消息
func TestCommitInfo_GetCommitMessage(t *testing.T) {
	commitInfo := CommitInfo{Message: "example"}

	require.Equal(t, "example", commitInfo.BuildCommitMessage())
}

// TestCommitInfo_GetSignatureInfo tests Git signature object generation
// Verifies client name and mailbox are set as expected in signature
//
// TestCommitInfo_GetSignatureInfo 测试 Git 签名对象生成
// 验证创建者名称和邮箱在签名中正确设置
func TestCommitInfo_GetSignatureInfo(t *testing.T) {
	commitInfo := CommitInfo{
		Name:    "John Doe",
		Mailbox: "johndoe@example.com",
	}
	authorInfo := commitInfo.GetObjectSignature()

	require.Equal(t, "John Doe", authorInfo.Name)
	require.Equal(t, "johndoe@example.com", authorInfo.Email)

	t.Log(neatjsons.S(authorInfo))
}

// TestCommitInfo_CheckFullMessage tests complete commit info features
// Verifies both signature generation and message handling work in sync
//
// TestCommitInfo_CheckFullMessage 测试完整的提交信息功能
// 验证签名生成和消息处理协同工作
func TestCommitInfo_CheckFullMessage(t *testing.T) {
	commitInfo := CommitInfo{
		Name:    "Jane Doe",
		Mailbox: "janedoe@example.com",
		Message: "example",
	}

	authorInfo := commitInfo.GetObjectSignature()

	require.Equal(t, "Jane Doe", authorInfo.Name)
	require.Equal(t, "janedoe@example.com", authorInfo.Email)
	require.Equal(t, "example", commitInfo.BuildCommitMessage())

	t.Log(neatjsons.S(authorInfo))
}

// TestCommitInfo_CheckNoneMessage tests default value actions
// Verifies blank CommitInfo uses package defaults for name and mailbox
//
// TestCommitInfo_CheckNoneMessage 测试默认值行为
// 验证空 CommitInfo 使用包默认值作为名称和邮箱
func TestCommitInfo_CheckNoneMessage(t *testing.T) {
	commitInfo := CommitInfo{}

	t.Log(commitInfo.BuildCommitMessage())

	authorInfo := commitInfo.GetObjectSignature()

	require.Equal(t, "gogit", authorInfo.Name)
	require.Equal(t, "gogit@github.com/go-xlan/gogit", authorInfo.Email)
	require.Contains(t, commitInfo.BuildCommitMessage(), packagePath)

	t.Log(neatjsons.S(authorInfo))
}

// TestCommitInfo_FluentPattern checks the fluent pattern methods
// Should chain method calls and set fields as expected
//
// TestCommitInfo_BuilderPattern 验证构建器模式方法
// 应该正确链接方法调用并设置字段
func TestCommitInfo_BuilderPattern(t *testing.T) {
	// Test NewCommitInfo and chaining
	// 测试 NewCommitInfo 和链式调用
	commitInfo := NewCommitInfo("First message").
		WithName("Test Account").
		WithMailbox("test@example.com").
		WithMessage("Updated message")

	require.Equal(t, "Test Account", commitInfo.Name)
	require.Equal(t, "test@example.com", commitInfo.Mailbox)
	require.Equal(t, "Updated message", commitInfo.Message)
}

// TestCommitInfo_BuildCommitMessage verifies message generation
// Should use provided message when available, otherwise generate default
//
// TestCommitInfo_BuildCommitMessage 验证消息生成
// 应该使用提供的消息或生成默认值
func TestCommitInfo_BuildCommitMessage(t *testing.T) {
	t.Run("With custom message", func(t *testing.T) {
		commitInfo := &CommitInfo{Message: "Custom message"}
		message := commitInfo.BuildCommitMessage()
		require.Equal(t, "Custom message", message)
	})

	t.Run("With blank message generates default", func(t *testing.T) {
		commitInfo := &CommitInfo{}
		message := commitInfo.BuildCommitMessage()
		require.Contains(t, message, packageName)
		require.Contains(t, message, packagePath)
		require.Contains(t, message, time.Now().Format("2006-01-02"))
	})
}

// TestCommitInfo_GetObjectSignature verifies signature generation
// Should use provided values when available, otherwise use defaults
//
// TestCommitInfo_GetObjectSignature 验证签名生成
// 应该使用提供的值或默认值
func TestCommitInfo_GetObjectSignature(t *testing.T) {
	t.Run("With custom values", func(t *testing.T) {
		commitInfo := &CommitInfo{
			Name:    "Custom Name",
			Mailbox: "custom@example.com",
		}
		signature := commitInfo.GetObjectSignature()
		require.Equal(t, "Custom Name", signature.Name)
		require.Equal(t, "custom@example.com", signature.Email)
		require.WithinDuration(t, time.Now(), signature.When, 1*time.Second)
	})

	t.Run("With default values", func(t *testing.T) {
		commitInfo := &CommitInfo{}
		signature := commitInfo.GetObjectSignature()
		require.Equal(t, packageName, signature.Name)
		require.Equal(t, packageName+"@"+packagePath, signature.Email)
	})
}
