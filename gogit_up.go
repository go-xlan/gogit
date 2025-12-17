package gogit

import (
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/yyle88/erero"
)

// GetCurrentBranch returns the name of the current branch
// Extracts branch name from HEAD to enable convenient access
// Returns short branch name such as "main" and "feature/xxx"
//
// GetCurrentBranch 返回当前分支的名称
// 从 HEAD 提取分支名称以便于访问
// 返回短分支名称，如 "main" 和 "feature/xxx"
func (c *Client) GetCurrentBranch() (string, error) {
	// Get HEAD reference
	// 获取 HEAD 引用
	head, err := c.repo.Head()
	if err != nil {
		return "", erero.Wro(err)
	}
	// Return short branch name
	// 返回短分支名称
	return head.Name().Short(), nil
}

// GetLatestCommit returns the latest commit object from HEAD
// Retrieves HEAD commit to inspect details such as message, signature, and timestamp
// Returns complete commit object with metadata included
//
// GetLatestCommit 返回 HEAD 的最新提交对象
// 获取 HEAD 提交以检查详情，如消息、签名和时间戳
// 返回包含元数据的完整提交对象
func (c *Client) GetLatestCommit() (*object.Commit, error) {
	// Get HEAD reference
	// 获取 HEAD 引用
	head, err := c.repo.Head()
	if err != nil {
		return nil, erero.Wro(err)
	}
	// Get commit object from HEAD hash
	// 从 HEAD 哈希获取提交对象
	commit, err := c.repo.CommitObject(head.Hash())
	if err != nil {
		return nil, erero.Wro(err)
	}
	return commit, nil
}

// HasChanges checks if uncommitted changes exist in the worktree
// Returns true when staged, modified, and untracked files exist
// Enables quick check to determine if commit is needed
//
// HasChanges 检查工作树中是否存在未提交的更改
// 当存在已暂存、已修改和未跟踪的文件时返回 true
// 提供快速检查以确定是否需要提交
func (c *Client) HasChanges() (bool, error) {
	// Get worktree status
	// 获取工作树状态
	status, err := c.tree.Status()
	if err != nil {
		return false, erero.Wro(err)
	}
	// Check if status map contains entries
	// 检查状态映射是否包含条目
	return len(status) > 0, nil
}

// GetRemoteURL returns the URL of the specified remote
// Retrieves URL from remote config using the given name
// Returns error when remote not found or has no URLs configured
//
// GetRemoteURL 返回指定远程的 URL
// 使用给定名称从远程配置获取 URL
// 未找到远程或未配置 URL 时返回错误
func (c *Client) GetRemoteURL(remoteName string) (string, error) {
	// Get remote by name
	// 按名称获取远程
	remote, err := c.repo.Remote(remoteName)
	if err != nil {
		return "", erero.Wro(err)
	}
	// Get URLs from remote config
	// 从远程配置获取 URL
	urls := remote.Config().URLs
	if len(urls) == 0 {
		return "", erero.New("remote has no URLs configured")
	}
	return urls[0], nil
}

// GetFirstRemoteURL returns the URL of the first available remote
// Returns error when no remotes exist or no URLs configured
// Returns error when fetching remotes fails
//
// GetFirstRemoteURL 返回第一个可用远程的 URL
// 当没有远程或未配置 URL 时返回错误
// 当获取远程失败时返回错误
func (c *Client) GetFirstRemoteURL() (string, error) {
	remotes, err := c.repo.Remotes()
	if err != nil {
		return "", erero.Wro(err)
	}
	if len(remotes) == 0 {
		return "", erero.New("remote repo is not configured")
	}
	urls := remotes[0].Config().URLs
	if len(urls) == 0 {
		return "", erero.New("remote has no URLs configured")
	}
	return urls[0], nil
}
