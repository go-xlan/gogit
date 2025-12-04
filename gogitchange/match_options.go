package gogitchange

import (
	"github.com/go-git/go-git/v5"
)

// MatchOptions configures file matching criteria for changed file processing
// Provides flexible filtering by file extension, path, and file status
// Supports fluent configuration pattern for convenient setup
//
// MatchOptions 配置用于变更文件处理的文件匹配条件
// 提供通过文件扩展名、路径和文件状态的灵活过滤
// 支持流畅配置模式以便于设置
type MatchOptions struct {
	matchType     string            // File extension screen like ".go", ".txt" // 文件扩展名过滤器，如 ".go", ".txt"
	matchPath     func(string) bool // Custom path matching function // 自定义路径匹配函数
	matchStatuses []git.StatusCode  // File status codes to match // 要匹配的文件状态码
}

// NewMatchOptions creates a new instance with default blank matching criteria
// Returns MatchOptions that supports fluent configuration chaining
// Use with MatchType and MatchPath to achieve complete setup
//
// NewMatchOptions 创建带默认空匹配条件的新实例
// 返回支持流畅配置链式调用的 MatchOptions
// 与 MatchType 和 MatchPath 一起使用以完成设置
func NewMatchOptions() *MatchOptions {
	return &MatchOptions{}
}

// MatchType sets file extension screen and returns updated MatchOptions
// Enables filtering files by extension like ".go", ".txt", ".md"
// Supports fluent configuration pattern enabling method chaining
//
// MatchType 设置文件扩展名过滤器并返回更新的 MatchOptions
// 支持按扩展名过滤文件，如 ".go", ".txt", ".md"
// 支持流畅配置模式以启用方法链式调用
func (m *MatchOptions) MatchType(fileExtension string) *MatchOptions {
	m.matchType = fileExtension
	return m
}

// MatchPath sets custom path matching function and returns updated MatchOptions
// Enables advanced path filtering using custom logic and conditions
// Function receives absolute path and returns true when files match
//
// MatchPath 设置自定义路径匹配函数并返回更新的 MatchOptions
// 支持使用自定义逻辑和条件进行高级路径过滤
// 函数接收绝对路径并在文件匹配时返回 true
func (m *MatchOptions) MatchPath(matchPath func(path string) bool) *MatchOptions {
	m.matchPath = matchPath
	return m
}

// MatchStatus sets file status codes to filter and returns updated MatchOptions
// Enables filtering files by Git status like Added, Modified, Deleted, Untracked
// Common status codes: git.Added ('A'), git.Modified ('M'), git.Deleted ('D'), git.Untracked ('?')
//
// MatchStatus 设置要过滤的文件状态码并返回更新的 MatchOptions
// 支持按 Git 状态过滤文件，如 Added、Modified、Deleted、Untracked
// 常用状态码：git.Added ('A')、git.Modified ('M')、git.Deleted ('D')、git.Untracked ('?')
func (m *MatchOptions) MatchStatus(statuses ...git.StatusCode) *MatchOptions {
	m.matchStatuses = statuses
	return m
}

// hasStatusMatch checks if the given file status matches any of the configured status codes
// Returns true if no status filter is set or if the file status matches any configured code
//
// hasStatusMatch 检查给定的文件状态是否匹配任何配置的状态码
// 如果未设置状态过滤器或文件状态匹配任何配置的代码则返回 true
func (m *MatchOptions) hasStatusMatch(fileStatus *git.FileStatus) bool {
	// No status constraints means accept everything
	// 没有状态约束意味着接受所有
	if len(m.matchStatuses) == 0 {
		return true
	}
	// Check if staging or worktree status matches any configured status
	// 检查暂存区或工作树状态是否匹配任何配置的状态
	for _, status := range m.matchStatuses {
		if fileStatus.Staging == status || fileStatus.Worktree == status {
			return true
		}
	}
	return false
}
