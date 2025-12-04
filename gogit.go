// Package gogit: Enhanced Git operations toolkit with streamlined repo management
// Delivers intelligent Git functions with comprehensive commit, amend, and remote sync capabilities
// Automatic staging, status tracking, and commit hash validation with remote push detection included
// Built on go-git package, provides cross-platform Git operations without CLI dependencies
//
// gogit: 增强的 Git 操作工具包，提供简化的仓库管理
// 提供智能 Git 功能，包含全面的提交、修正和远程同步能力
// 包含自动暂存、状态跟踪和提交哈希验证，支持远程推送检测
// 基于 go-git 库构建，提供跨平台 Git 操作，无需 CLI 依赖
package gogit

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-xlan/gogit/gogitassist"
	"github.com/pkg/errors"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/must"
	"github.com/yyle88/tern/zerotern"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Client represents a Git repo client with enhanced operations
// Encapsulates repo and worktree to streamline Git management
// Provides advanced interface with robust exception handling
//
// Client 代表具有增强操作的 Git 仓库客户端
// 封装仓库和工作树以简化 Git 管理
// 提供高级接口，带有健壮的异常处理
type Client struct {
	repo *git.Repository // Git repo instance // Git 仓库实例
	tree *git.Worktree   // Working tree with ignore file support // 支持忽略文件的工作树
}

// NewClient creates a new Git client with specified repo and worktree
// Combines repo and worktree to enable comprehensive Git operations
// Worktree should include suitable ignore file settings to optimize speed
//
// NewClient 使用指定的仓库和工作树创建新的 Git 客户端
// 结合仓库和工作树以启用全面的 Git 操作
// 工作树应包含适当的忽略文件配置以优化速度
func NewClient(repo *git.Repository, tree *git.Worktree) *Client {
	return &Client{
		repo: repo,
		tree: tree, // Use worktree with ignore file support // 使用支持忽略文件的工作树
	}
}

// New initializes a Git client at the specified project root path
// Opens existing repo and creates worktree with ignore file support
// Returns configured client available to use
//
// New 在指定的项目根路径初始化 Git 客户端
// 打开现有仓库并创建支持忽略文件的工作树
// 返回配置好的客户端，可以直接使用
func New(root string) (*Client, error) {
	// Initialize repo and worktree with ignore file support
	// 初始化仓库和工作树，支持忽略文件
	repo, tree, err := gogitassist.NewRepoTreeWithIgnore(root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	// Create client with configured repo and worktree
	// 使用配置的仓库和工作树创建客户端
	client := NewClient(repo, tree)
	return client, nil
}

// Repo returns the underlying Git repo instance
// Enables access to basic repo operations when needed
//
// Repo 返回底层的 Git 仓库实例
// 在需要时提供对基础仓库操作的访问
func (c *Client) Repo() *git.Repository {
	return c.repo
}

// Tree returns the Git worktree used in file operations
// Enables direct worktree manipulation when needed
//
// Tree 返回用于文件操作的 Git 工作树
// 在需要时支持直接操作工作树
func (c *Client) Tree() *git.Worktree {
	return c.tree
}

// AddAll stages changes including new files, modifications and deletions
// Equivalent to 'git add --all' operation with comprehensive change detection
// Fails with an issue when the staging operation encounters problems
//
// AddAll 暂存更改，包括新文件、修改和删除
// 等同于 'git add --all' 操作，具有全面的更改检测
// 当暂存操作遇到问题时返回错误
func (c *Client) AddAll() error {
	if err := c.tree.AddWithOptions(&git.AddOptions{All: true}); err != nil {
		return erero.Wro(err)
	}
	return nil
}

// Status returns the current status of the worktree with file change information
// Provides comprehensive status including staged, modified, and untracked files
// Returns git.Status map with file paths and corresponding change status
//
// Status 返回工作树的当前状态及文件更改信息
// 提供全面的状态，包括已暂存、已修改和未跟踪的文件
// 返回包含文件路径及其相应更改状态的 git.Status 映射
func (c *Client) Status() (git.Status, error) {
	// Get current worktree status with comprehensive file change detection
	// 获取当前工作树状态，全面检测文件更改
	status, err := c.tree.Status()
	if err != nil {
		return nil, erero.Wro(err)
	}
	return status, nil
}

// CommitAll commits staged changes with the provided commit info
// Creates a new commit with staged files and applies specified signature
// Returns commit hash string, blank string when no changes exist
//
// CommitAll 使用提供的提交信息提交已暂存的更改
// 创建包含已暂存文件的新提交并应用指定的签名
// 返回提交哈希字符串，无更改时返回空字符串
func (c *Client) CommitAll(info *CommitInfo) (string, error) {
	// Build commit message from commit info
	// 从提交信息构建提交消息
	message := info.BuildCommitMessage()
	zaplog.ZAPS.Skip1.SUG.Info("commit-message:", message)

	commitHash, err := c.tree.Commit(message, &git.CommitOptions{
		All:    true, // Commit deleted files.
		Author: info.GetObjectSignature(),
	})
	if err != nil {
		if errors.Is(err, git.ErrEmptyCommit) {
			return "", nil
		}
		return "", erero.Wro(err)
	}
	// Log completed commit operation
	// 记录成功的提交操作
	zaplog.ZAPS.Skip1.LOG.Info("commit-success", zap.String("hash", commitHash.String()))
	return c.checkCommitHash(commitHash)
}

// AmendConfig represents settings used when amending commits
// Contains commit info and amend-pushed flag
//
// AmendConfig 代表修正提交时使用的配置
// 包含提交信息和强制修正已推送提交的标志
type AmendConfig struct {
	CommitInfo *CommitInfo // New commit info for amend operation // 修正操作的新提交信息
	ForceAmend bool        // Allow amend even if commit was pushed // 即使提交已推送也允许修正
}

// AmendCommit amends the previous commit with new info
// Modifies the last commit with updated message, signature, and staged changes
// Blocks amending pushed commits unless ForceAmend is enabled
//
// AmendCommit 使用新信息修正上一个提交
// 使用更新的消息、签名和已暂存的更改修改最后一个提交
// 除非启用 ForceAmend，否则阻止修正已推送的提交
func (c *Client) AmendCommit(cfg *AmendConfig) (string, error) {
	// Check if commit was pushed before allowing amend (unless forced)
	// 检查提交是否已推送，在允许修正前（除非强制）
	if !cfg.ForceAmend {
		// Validate HEAD has been pushed to some remote
		// 验证 HEAD 是否已推送到某个远程
		pushed, err := c.IsLatestCommitPushed()
		if err != nil {
			return "", erero.Wro(err)
		}
		if pushed {
			return "", erero.New("cannot amend a commit that has been pushed")
		}
	}
	// Determine commit message: use provided message, else reuse existing one
	// 确定提交消息：使用提供的消息，否则重用现有消息
	message := cfg.CommitInfo.Message
	if message == "" { // Use latest commit message when no new message provided // 未提供新消息时使用最新提交消息
		// Get latest commit reference and message
		// 获取最新提交引用和消息
		topReference := done.VCE(c.repo.Head()).Nice()
		commitObject := done.VCE(c.repo.CommitObject(topReference.Hash())).Nice()
		message = zerotern.VF(commitObject.Message, func() string {
			return cfg.CommitInfo.BuildCommitMessage()
		})
	}
	zaplog.ZAPS.Skip1.SUG.Info("amend-message:", message)

	// Execute amend operation with new signature
	// 使用新签名执行 amend 操作
	commitHash, err := c.tree.Commit(message, &git.CommitOptions{
		Author: cfg.CommitInfo.GetObjectSignature(),
		Amend:  true, // Note: "all" and "amend" are exclusive // 注意："all" 和 "amend" 不能同时使用
	})
	if err != nil {
		// Handle blank amend case
		// 处理空 amend 情况
		if errors.Is(err, git.ErrEmptyCommit) {
			return "", nil
		}
		return "", erero.Wro(err)
	}
	// Log completed amend operation
	// 记录成功的 amend 操作
	zaplog.ZAPS.Skip1.LOG.Info("amend-commit-success", zap.String("hash", commitHash.String()))
	return c.checkCommitHash(commitHash)
}

// checkCommitHash validates commit hash and retrieves commit object as validation
// Ensures commit wholeness via hash matching and logs commit details
// Returns commit hash string that represents the completed operation
//
// checkCommitHash 验证提交哈希并检索提交对象进行验证
// 通过哈希匹配确保提交完整性并记录提交详情
// 返回代表已完成操作的提交哈希字符串
func (c *Client) checkCommitHash(commitHash plumbing.Hash) (string, error) {
	// Retrieve commit object from repo
	// 从仓库检索提交对象
	object, err := c.repo.CommitObject(commitHash)
	if err != nil {
		return "", erero.Wro(err)
	}
	// Log commit object details and check hash matching
	// 记录提交对象详情并验证哈希一致性
	zaplog.ZAPS.Skip2.SUG.Info(object)
	must.Same(commitHash, object.Hash)
	return object.Hash.String(), nil
}

// IsLatestCommitPushedToRemote checks if HEAD has been pushed to specified remote
// Compares HEAD hash with remote branch hash to decide push status
// Returns true when hashes match, false when remote branch not found
//
// IsLatestCommitPushedToRemote 检查 HEAD 是否已推送到指定的远程
// 比较 HEAD 哈希与远程分支哈希来判断推送状态
// 哈希匹配时返回 true，未找到远程分支时返回 false
func (c *Client) IsLatestCommitPushedToRemote(remoteName string) (bool, error) {
	// Get current branch reference (HEAD)
	// 获取当前分支引用（HEAD）
	branchReference := done.VCE(c.repo.Head()).Nice()
	// Get remote branch hash to compare
	// 获取远程分支哈希进行比较
	remoteReference, err := c.repo.Reference(plumbing.ReferenceName(fmt.Sprintf("refs/remotes/%s/%s", remoteName, branchReference.Name().Short())), false)
	if err != nil {
		// Remote reference not found (branch not pushed yet)
		// 远程引用未找到（分支尚未推送）
		if errors.Is(err, plumbing.ErrReferenceNotFound) {
			return false, nil // Expected error when remote branch doesn't exist // 远程分支不存在时的预期错误
		}
		return false, erero.Wro(err) // Unexpected issue encountered // 遇到意外问题
	}
	// Compare current and remote commit hashes to determine push status
	// 比较本地和远程提交哈希来确定推送状态
	return branchReference.Hash() == remoteReference.Hash(), nil
}

// IsLatestCommitPushed checks if HEAD has been pushed to configured remotes
// Iterates through remotes and checks matching commit hashes
// Returns true when HEAD exists in some remote, false when not pushed
//
// IsLatestCommitPushed 检查 HEAD 是否已推送到配置的远程
// 遍历远程并检查匹配的提交哈希
// 当 HEAD 存在于某个远程时返回 true，未推送时返回 false
func (c *Client) IsLatestCommitPushed() (bool, error) {
	// Get all configured remote repos
	// 获取所有配置的远程仓库
	remotes, err := c.repo.Remotes()
	if err != nil {
		return false, erero.Wro(err)
	}
	// Check each remote repo to find matching commits
	// 检查每个远程仓库以查找匹配的提交
	for _, remote := range remotes {
		remoteName := remote.Config().Name

		// Check if current commit exists in this remote
		// 检查当前提交是否存在于此远程
		if matched, err := c.IsLatestCommitPushedToRemote(remoteName); err != nil {
			return false, erero.Wro(err)
		} else if matched {
			return true, nil // Found matching commit in remote // 在远程中找到匹配的提交
		}
	}
	return false, nil // No remote repo contains current branch commit // 没有远程仓库包含当前分支提交
}
