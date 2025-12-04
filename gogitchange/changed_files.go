// Package gogitchange: Git changed files detection and processing toolkit
// Provides utilities for identifying, filtering, and processing changed files in Git worktree
// Features customizable matching options and file processing capabilities for changed files
//
// gogitchange: Git 变更文件检测和处理工具包
// 提供用于识别、过滤和处理 Git 工作树中变更文件的实用工具
// 具有可定制的匹配选项和针对变更文件的文件处理功能
package gogitchange

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/yyle88/erero"
	"github.com/yyle88/must"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/osexistpath/ossoftexist"
)

// ChangedFileManager manages detection and processing of changed files in Git worktree
// Encapsulates project path and worktree for efficient file change operations
// Provides foundation for changed file filtering and processing workflows
//
// ChangedFileManager 管理 Git 工作树中变更文件的检测和处理
// 封装项目路径和工作树以实现高效的文件变更操作
// 为变更文件过滤和处理工作流提供基础
type ChangedFileManager struct {
	projectPath string        // Project root path for file operations // 用于文件操作的项目根路径
	tree        *git.Worktree // Git worktree for status checking // 用于状态检查的 Git 工作树
}

// NewChangedFileManager creates a new component to handle changed files
// Validates project path and associates worktree enabling file change detection
// Returns configured component usable in changed file processing
//
// NewChangedFileManager 创建用于处理变更文件的新组件
// 验证项目路径并关联工作树以启用文件变更检测
// 返回可用于变更文件处理的配置好的组件
func NewChangedFileManager(projectPath string, worktree *git.Worktree) *ChangedFileManager {
	return &ChangedFileManager{
		projectPath: osmustexist.ROOT(must.Nice(projectPath)),
		tree:        worktree,
	}
}

// Foreach iterates through changed files (excluding deleted) and processes each
// Applies matching options to screen files by type and path criteria
// Executes provided process function on each qualifying changed file
//
// Foreach 遍历变更的文件（排除已删除的）并处理每个文件
// 应用匹配选项按类型和路径条件过滤文件
// 对每个符合条件的变更文件执行提供的处理函数
func (m *ChangedFileManager) Foreach(matchOptions *MatchOptions, process func(path string) error) error {
	statusMap, err := m.tree.Status()
	if err != nil {
		return erero.Wro(err)
	}

	for relativePath, status := range statusMap {
		// Screen out deleted files as these cannot be processed
		// 过滤掉已删除的文件，因为它们无法被处理
		if status.Staging == git.Deleted {
			continue
		}

		// Screen files by status if status matching is specified
		// 如果指定了状态匹配，则按状态过滤文件
		if !matchOptions.hasStatusMatch(status) {
			continue
		}

		// Screen files by extension if type matching is specified
		// 如果指定了类型匹配，则按扩展名过滤文件
		if matchOptions.matchType != "" && filepath.Ext(relativePath) != matchOptions.matchType {
			continue
		}

		// Screen files by path criteria using custom match function
		// 使用自定义匹配器函数按路径条件过滤文件
		if matchOptions.matchPath != nil && !matchOptions.matchPath(filepath.Join(m.projectPath, relativePath)) {
			continue
		}

		// Construct complete file path for processing
		// 构建用于处理的完整文件路径
		path := filepath.Join(m.projectPath, relativePath)

		// Process file when it exists (not deleted or missing)
		// 仅在文件存在时处理（未删除或缺失）
		if ossoftexist.IsFile(path) {
			// Execute custom processing function on the file
			// 对文件执行自定义处理函数
			if err := process(path); err != nil {
				return erero.Wro(err)
			}
		}
	}
	return nil
}

// ListChangedFilePaths returns list of changed file paths matching specified criteria
// Uses Foreach within to collect paths of changed files
// Returns slice of absolute paths for qualifying changed files
//
// ListChangedFilePaths 返回符合指定条件的变更文件路径列表
// 内部使用 Foreach 收集变更文件的路径
// 返回符合条件的变更文件的绝对路径切片
func (m *ChangedFileManager) ListChangedFilePaths(matchOptions *MatchOptions) ([]string, error) {
	var paths = make([]string, 0)
	if err := m.Foreach(matchOptions, func(path string) error {
		paths = append(paths, path)
		return nil
	}); err != nil {
		return nil, erero.Wro(err)
	}
	return paths, nil
}

// FormatChangedGoFiles formats all changed Go code files (features moved to distinct module)
// This function has been relocated to go-mate/go-commit for improved design
//
// FormatChangedGoFiles 格式化所有变化的 Go 代码文件（功能已移至其他模块）
// 此函数已重定位到 go-mate/go-commit 以获得更好的模块化
// func (m *ChangedFileManager) FormatChangedGoFiles(matchOptions *MatchOptions) error {
// code move to -> https://github.com/go-mate/go-commit
// }

// ForeachChangedGoFile iterates through all changed Go files and processes each one
// Screens changed files to include just .go files and applies custom processing
// Convenient Foreach facade targeting Go-specific file operations
//
// ForeachChangedGoFile 遍历所有变化的 Go 文件并处理每一个
// 过滤变更文件仅包含 .go 文件并应用自定义处理
// 针对 Go 特定文件操作的 Foreach 便捷外观
func (m *ChangedFileManager) ForeachChangedGoFile(matchOptions *MatchOptions, process func(path string) error) error {
	if err := m.Foreach(matchOptions, func(path string) error {
		if filepath.Ext(path) != ".go" {
			return nil
		}
		if err := process(path); err != nil {
			return erero.Wro(err)
		}
		return nil
	}); err != nil {
		return erero.Wro(err)
	}
	return nil
}
