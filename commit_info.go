package gogit

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/yyle88/eroticgo"
	"github.com/yyle88/tern/zerotern"
	"github.com/yyle88/zaplog"
)

// Package constants for default commit message generation
// 用于默认提交消息生成的包常量
const packageName = "gogit"                    // Package name for commit messages // 用于提交消息的包名
const packagePath = "github.com/go-xlan/gogit" // Package path for commit messages // 用于提交消息的包路径

// CommitInfo represents Git commit author info and message content
// Contains author details and commit message for Git operations
// Supports fluent pattern for convenient configuration
//
// CommitInfo 代表 Git 提交作者信息和消息内容
// 包含作者详情和用于 Git 操作的提交消息
// 支持构建器模式以便于配置
type CommitInfo struct {
	Name    string // Author name for Git commits // 用于 Git 提交的作者姓名
	Mailbox string // Author mailbox for Git commits // 用于 Git 提交的作者邮箱地址
	Message string // Commit message content // 提交消息内容
}

// NewCommitInfo creates a new CommitInfo instance with specified message
// Initializes commit info with provided message, leaving author fields blank for setup
// Returns CommitInfo instance set for fluent pattern chaining
//
// NewCommitInfo 使用指定消息创建新的 CommitInfo 实例
// 使用提供的消息初始化提交信息，留空作者字段以便后续配置
// 返回准备进行构建器模式链式调用的 CommitInfo 实例
func NewCommitInfo(message string) *CommitInfo {
	return &CommitInfo{
		Message: message,
	}
}

// WithName sets the author name and returns the updated CommitInfo instance
// Enables fluent method pattern for convenient configuration chaining
// Name represents the Git author name displayed in commit records
//
// WithName 设置作者姓名并返回更新的 CommitInfo 实例
// 支持流畅的构建器模式以便于配置链式调用
// Name 代表在提交历史中显示的 Git 作者姓名
func (c *CommitInfo) WithName(name string) *CommitInfo {
	c.Name = name
	return c
}

// WithMailbox sets the author mailbox and returns the updated CommitInfo instance
// Configures mailbox for Git commit author signature
// Supports fluent pattern chaining for convenient configuration
//
// WithMailbox 设置作者邮箱地址并返回更新的 CommitInfo 实例
// 为 Git 提交作者签名配置邮箱地址
// 支持构建器模式链式调用以便于配置
func (c *CommitInfo) WithMailbox(mailbox string) *CommitInfo {
	c.Mailbox = mailbox
	return c
}

// WithMessage sets the commit message content and returns the updated CommitInfo instance
// Replaces any existing message with the provided content
// Enables fluent configuration through method pattern chaining
//
// WithMessage 设置提交消息内容并返回更新的 CommitInfo 实例
// 使用提供的内容替换任何现有消息
// 通过构建器模式链式调用实现流畅配置
func (c *CommitInfo) WithMessage(message string) *CommitInfo {
	c.Message = message
	return c
}

// BuildCommitMessage generates commit message using provided content or creates default message
// Returns custom message if available, otherwise generates timestamp-based default message
// Default format includes package name, path, and current timestamp
//
// BuildCommitMessage 使用提供的内容生成提交消息或创建默认消息
// 如果有自定义消息则返回，否则生成基于时间戳的默认消息
// 默认格式包含包名、路径和当前时间戳
func (c *CommitInfo) BuildCommitMessage() string {
	// Use custom message if provided, otherwise generate default message
	// 如果提供了自定义消息则使用，否则生成默认消息
	return zerotern.VF(c.Message, func() string {
		// Generate default commit message with package info and timestamp
		// 生成包含包信息和时间戳的默认提交消息
		message := fmt.Sprintf(`git commit -m "[%s](%s) %s"`, packageName, packagePath, time.Now().Format("2006-01-02 15:04:05"))
		// Log generated message for debugging
		// 记录生成的消息用于调试
		zaplog.SUG.Debugln(eroticgo.BLUE.Sprint(fmt.Sprintf(`git commit -m "%s"`, message)))
		return message
	})
}

// GetObjectSignature creates Git signature object for commit operations
// Builds object.Signature with author name, mailbox, and current timestamp
// Uses package defaults when author info is not provided
//
// GetObjectSignature 为提交操作创建 Git 签名对象
// 使用作者姓名、邮箱地址和当前时间戳构建 object.Signature
// 在未提供作者信息时使用包默认值
func (c *CommitInfo) GetObjectSignature() *object.Signature {
	// Create signature with provided or default values
	// 使用提供的或默认值创建签名
	return &object.Signature{
		Name:  zerotern.VV(c.Name, packageName),                                       // Use provided name or package default // 使用提供的姓名或包默认值
		Email: zerotern.VV(c.Mailbox, fmt.Sprintf("%s@%s", packageName, packagePath)), // Use provided mailbox or package default // 使用提供的邮箱或包默认值
		When:  time.Now(),                                                             // Current timestamp for commit // 当前时间戳用于提交
	}
}
