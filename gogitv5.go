package gogit

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-xlan/gogit/gogitv5ops"
	"github.com/pkg/errors"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/tern/zerotern"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type Client struct {
	repository *git.Repository
	worktree   *git.Worktree
}

// New creates a new Client instance.
func New(root string) (*Client, error) {
	repository, worktree, err := gogitv5ops.NewRepoWorktreeWithIgnore(root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return NewClient(repository, worktree)
}

// NewClient creates a Client from an existing repository and worktree.
func NewClient(repository *git.Repository, worktree *git.Worktree) (*Client, error) {
	return &Client{
		repository: repository,
		worktree:   worktree, // It's best to use the worktree with ignored files to avoid unnecessary commits.
	}, nil
}

// AddAll stages all changes (including deletions) for commit.
func (G *Client) AddAll() error {
	if err := G.worktree.AddWithOptions(&git.AddOptions{All: true}); err != nil {
		return erero.Wro(err)
	}
	return nil
}

// Status returns the current status of the worktree.
func (G *Client) Status() (git.Status, error) {
	status, err := G.worktree.Status() // We can verify the current status of the worktree.
	if err != nil {
		return nil, erero.Wro(err)
	}
	return status, nil
}

// CommitAll commits all staged changes with the provided commit information.
func (G *Client) CommitAll(commitInfo *CommitInfo) (string, error) {
	message := commitInfo.GetCommitMessage()
	zaplog.ZAPS.Skip1.SUG.Info("commit: ", "msg: ", message)

	return G.verifyCommitHash(G.worktree.Commit(message, &git.CommitOptions{
		All:    true, // Whether to commit deleted files. Usually true, as it's rare to commit without deleting.
		Author: commitInfo.GetObjectSignature(),
	}))
}

type AmendConfig struct {
	CommitInfo *CommitInfo
	ForceAmend bool
}

// AmendCommit amends the previous commit with the new message, or uses the latest commit message if empty.
func (G *Client) AmendCommit(amendConfig *AmendConfig) (string, error) {
	if !amendConfig.ForceAmend {
		if pushed, err := G.IsPushedToAnyRemote(); err != nil {
			return "", erero.Wro(err)
		} else if pushed {
			return "", erero.New("cannot amend a commit that has been pushed")
		}
	}
	message := amendConfig.CommitInfo.Message
	if message == "" { // If no message is provided, use the latest commit message to amend.
		headReference := done.VCE(G.repository.Head()).Nice()
		commitObject := done.VCE(G.repository.CommitObject(headReference.Hash())).Nice()
		message = zerotern.VF(commitObject.Message, func() string {
			return amendConfig.CommitInfo.GetCommitMessage()
		})
	}
	zaplog.ZAPS.Skip1.SUG.Info("amend: ", "msg: ", message)

	return G.verifyCommitHash(G.worktree.Commit(message, &git.CommitOptions{
		Author: amendConfig.CommitInfo.GetObjectSignature(),
		Amend:  true, // Note: "all" and "amend" cannot be used together, so "all" is not set here.
	}))
}

// verifyCommitHash checks and logs the commit hash.
func (G *Client) verifyCommitHash(commitHash plumbing.Hash, err error) (string, error) {
	if err != nil {
		if errors.Is(err, git.ErrEmptyCommit) {
			return "", nil
		}
		return "", erero.Wro(err)
	}
	zaplog.ZAPS.Skip2.LOG.Info("commit", zap.String("hash", commitHash.String()))

	commitObject, err := G.repository.CommitObject(commitHash)
	if err != nil {
		return "", erero.Wro(err)
	}
	zaplog.ZAPS.Skip2.SUG.Info(commitObject)
	return commitHash.String(), nil
}

// IsHashMatchedRemote 检查当前分支是否已经推送到指定的远程仓库
func (G *Client) IsHashMatchedRemote(remoteName string) (bool, error) {
	// 获取当前分支引用
	headReference := done.VCE(G.repository.Head()).Nice()
	// 获取远程引用
	remoteReference, err := G.repository.Reference(plumbing.ReferenceName(fmt.Sprintf("refs/remotes/%s/%s", remoteName, headReference.Name().Short())), false)
	if err != nil {
		if errors.Is(err, plumbing.ErrReferenceNotFound) {
			return false, nil //报错通常是因为远程引用未找到，这里的错误是有预期的
		}
		return false, erero.Wro(err) // 其它错误
	}
	// 如果当前分支的提交哈希和远程分支相同，说明本地分支已经推送到远程
	return headReference.Hash() == remoteReference.Hash(), nil
}

// IsPushedToAnyRemote 检查当前分支是否已经推送到任何远程仓库
func (G *Client) IsPushedToAnyRemote() (bool, error) {
	remotes, err := G.repository.Remotes()
	if err != nil {
		return false, erero.Wro(err)
	}
	for _, remote := range remotes {
		remoteName := remote.Config().Name

		if matched, err := G.IsHashMatchedRemote(remoteName); err != nil {
			return false, erero.Wro(err)
		} else if matched {
			return true, nil
		}
	}
	return false, nil // 没有任何远程仓库有当前分支的提交
}
