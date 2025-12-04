package gogitassist

import (
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/yyle88/erero"
)

// InitRepo initializes a new Git repo at the specified path
// Creates .git directory and sets up repo structure
// Returns the initialized repo instance
//
// InitRepo 在指定路径初始化新的 Git 仓库
// 创建 .git 目录并设置仓库结构
// 返回初始化的仓库实例
func InitRepo(path string) (*git.Repository, error) {
	repo, err := git.PlainInit(path, false)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return repo, nil
}

// SetConfigUserName sets the user.name in repo config
// Configures the name used as commit authorship in this repo
//
// SetConfigUserName 设置仓库配置中的 user.name
// 配置此仓库中用于提交署名的名称
func SetConfigUserName(repo *git.Repository, name string) error {
	cfg, err := repo.Config()
	if err != nil {
		return erero.Wro(err)
	}
	cfg.User.Name = name
	if err := repo.SetConfig(cfg); err != nil {
		return erero.Wro(err)
	}
	return nil
}

// SetConfigUserMailbox sets the user.email in repo config
// Configures the mailbox used as commit authorship in this repo
//
// SetConfigUserMailbox 在仓库配置中设置 user.email 属性
// 配置此仓库中用于提交署名的邮箱
func SetConfigUserMailbox(repo *git.Repository, mailbox string) error {
	cfg, err := repo.Config()
	if err != nil {
		return erero.Wro(err)
	}
	cfg.User.Email = mailbox
	if err := repo.SetConfig(cfg); err != nil {
		return erero.Wro(err)
	}
	return nil
}

// SetConfigUserInfo sets both user.name and user.email in repo config
// Configures the authorship info used when creating commits in this repo
//
// SetConfigUserInfo 在仓库配置中设置 user.name 和 user.email 属性
// 配置此仓库中创建提交时使用的署名信息
func SetConfigUserInfo(repo *git.Repository, username, mailbox string) error {
	cfg, err := repo.Config()
	if err != nil {
		return erero.Wro(err)
	}
	cfg.User.Name = username
	cfg.User.Email = mailbox
	if err := repo.SetConfig(cfg); err != nil {
		return erero.Wro(err)
	}
	return nil
}

// AddRemote adds a new remote to the repo
// Creates named reference to external repo location
//
// AddRemote 向仓库添加新的远程
// 创建对外部仓库位置的命名引用
func AddRemote(repo *git.Repository, name, remoteURL string) error {
	_, err := repo.CreateRemote(&config.RemoteConfig{
		Name: name,
		URLs: []string{remoteURL},
	})
	if err != nil {
		return erero.Wro(err)
	}
	return nil
}

// RemoveRemote removes an existing remote from the repo
// Deletes named remote reference from configuration
//
// RemoveRemote 从仓库删除现有的远程
// 从配置中删除命名的远程引用
func RemoveRemote(repo *git.Repository, name string) error {
	if err := repo.DeleteRemote(name); err != nil {
		return erero.Wro(err)
	}
	return nil
}

// Commit stages all files and creates a commit with provided message and authorship
// Returns the commit hash
//
// Commit 暂存所有文件并使用提供的消息和署名创建提交
// 返回提交哈希
func Commit(repo *git.Repository, message, username, mailbox string) (plumbing.Hash, error) {
	tree, err := repo.Worktree()
	if err != nil {
		return plumbing.ZeroHash, erero.Wro(err)
	}
	// Stage all files
	// 暂存所有文件
	if err := tree.AddWithOptions(&git.AddOptions{All: true}); err != nil {
		return plumbing.ZeroHash, erero.Wro(err)
	}
	// Create commit
	// 创建提交
	hash, err := tree.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  username,
			Email: mailbox,
			When:  time.Now(),
		},
	})
	if err != nil {
		return plumbing.ZeroHash, erero.Wro(err)
	}
	return hash, nil
}
