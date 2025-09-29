// Package gogitassist: Git repo initialization and worktree management assistance
// Provides support functions for creating Git repos and configuring worktree with ignore patterns
// Features system-wide, shared, and project-wide gitignore pattern loading and application
//
// gogitassist: Git 仓库初始化和工作树管理辅助包
// 提供用于创建 Git 仓库和配置带忽略模式的工作树的辅助函数
// 具有系统级、全局级和项目级 gitignore 模式加载和应用功能
package gogitassist

import (
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// NewRepo opens an existing Git repo at the specified root path
// Returns configured repo instance prepared for Git operations
// Wraps go-git PlainOpen with error handling
//
// NewRepo 在指定根路径打开现有 Git 仓库
// 返回配置好的仓库实例，准备进行 Git 操作
// 使用错误处理包装 go-git PlainOpen
func NewRepo(root string) (*git.Repository, error) {
	repo, err := git.PlainOpen(root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return repo, nil
}

// NewRepoTreeWithIgnore creates repo and worktree with comprehensive ignore pattern support
// Loads and applies ignore patterns from system, system, and project levels
// Returns both repo and worktree instances configured with ignore patterns
//
// NewRepoTreeWithIgnore 创建带全面忽略模式支持的仓库和工作树
// 从系统、全局和项目级别加载并应用忽略模式
// 返回配置了忽略模式的仓库和工作树实例
func NewRepoTreeWithIgnore(root string) (repo *git.Repository, tree *git.Worktree, err error) {
	repo, err = NewRepo(root)
	if err != nil {
		return nil, nil, erero.Wro(err)
	}
	tree, err = repo.Worktree()
	if err != nil {
		return nil, nil, erero.Wro(err)
	}

	// Apply ignore patterns in sequence: system < system < project
	// 按优先级应用忽略模式：系统 < 全局 < 项目

	// Lowest sequence - system-wide ignore patterns
	// 最低优先级 - 系统级忽略模式
	SetIgnorePatterns(tree, done.VAE(gitignore.LoadSystemPatterns(osfs.New("/"))).Done())

	// Medium sequence - system ignore patterns
	// 中等优先级 - 全局忽略模式
	SetIgnorePatterns(tree, done.VAE(gitignore.LoadGlobalPatterns(osfs.New("/"))).Done())

	// Highest sequence - project-specific ignore patterns
	// 最高优先级 - 项目特定忽略模式
	SetIgnorePatterns(tree, done.VAE(LoadProjectIgnorePatterns(root)).Done())

	return repo, tree, nil
}

// DebugRepo outputs detailed repo info for debugging purposes
// Logs HEAD reference and commit object details
// Handles errors with structured logging
//
// DebugRepo 输出详细的仓库信息用于调试目的
// 记录 HEAD 引用和提交对象详情
// 使用结构化日志优雅处理错误
func DebugRepo(repo *git.Repository) {
	topReference, err := repo.Head()
	if err != nil {
		zaplog.LOG.Error("wrong-when-get-head-reference", zap.Error(err))
		return
	}
	zaplog.SUG.Debugln("commit", topReference.String())

	commitObject, err := repo.CommitObject(topReference.Hash())
	if err != nil {
		zaplog.LOG.Error("wrong-when-get-commit-object", zap.Error(err))
		return
	}
	zaplog.SUG.Debugln(commitObject.String())
	zaplog.SUG.Debugln("return", "debug-repo-function-return")
}
