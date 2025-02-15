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
	"github.com/yyle88/rese"
	"github.com/yyle88/tern/zerotern"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type Client struct {
	repo *git.Repository
	tree *git.Worktree
}

func NewClient(repo *git.Repository, tree *git.Worktree) *Client {
	return &Client{
		repo: repo,
		tree: tree, // It's best to use the tree with ignored files.
	}
}

func New(root string) (*Client, error) {
	repo, tree, err := gogitassist.NewRepoTreeWithIgnore(root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	client := NewClient(repo, tree)
	return client, nil
}

func MustNew(root string) *Client {
	return rese.P1(New(root))
}

func (G *Client) Repo() *git.Repository {
	return G.repo
}

func (G *Client) Tree() *git.Worktree {
	return G.tree
}

// AddAll stages all changes (including deletions).
func (G *Client) AddAll() error {
	if err := G.tree.AddWithOptions(&git.AddOptions{All: true}); err != nil {
		return erero.Wro(err)
	}
	return nil
}

// Status returns the current status of the worktree.
func (G *Client) Status() (git.Status, error) {
	status, err := G.tree.Status() // We can verify the current status of the worktree.
	if err != nil {
		return nil, erero.Wro(err)
	}
	return status, nil
}

// CommitAll commits all staged changes with the provided commit information.
func (G *Client) CommitAll(info *CommitInfo) (string, error) {
	message := info.BuildCommitMessage()
	zaplog.ZAPS.Skip1.SUG.Info("commit-message:", message)

	if commitHash, err := G.tree.Commit(message, &git.CommitOptions{
		All:    true, // Commit deleted files.
		Author: info.GetObjectSignature(),
	}); err != nil {
		if errors.Is(err, git.ErrEmptyCommit) {
			return "", nil
		}
		return "", erero.Wro(err)
	} else {
		zaplog.ZAPS.Skip1.LOG.Info("commit-success", zap.String("hash", commitHash.String()))
		return G.checkCommitHash(commitHash)
	}
}

type AmendConfig struct {
	CommitInfo *CommitInfo
	ForceAmend bool
}

// AmendCommit amends the previous commit.
func (G *Client) AmendCommit(cfg *AmendConfig) (string, error) {
	if !cfg.ForceAmend {
		pushed, err := G.IsLatestCommitPushed()
		if err != nil {
			return "", erero.Wro(err)
		}
		if pushed {
			return "", erero.New("cannot amend a commit that has been pushed")
		}
	}
	message := cfg.CommitInfo.Message
	if message == "" { // If no message is provided, use the latest commit message to amend.
		topReference := done.VCE(G.repo.Head()).Nice()
		commitObject := done.VCE(G.repo.CommitObject(topReference.Hash())).Nice()
		message = zerotern.VF(commitObject.Message, func() string {
			return cfg.CommitInfo.BuildCommitMessage()
		})
	}
	zaplog.ZAPS.Skip1.SUG.Info("amend-message:", message)

	if commitHash, err := G.tree.Commit(message, &git.CommitOptions{
		Author: cfg.CommitInfo.GetObjectSignature(),
		Amend:  true, // Note: "all" and "amend" cannot be used at the same time, so "all" is not set here.
	}); err != nil {
		if errors.Is(err, git.ErrEmptyCommit) {
			return "", nil
		}
		return "", erero.Wro(err)
	} else {
		zaplog.ZAPS.Skip1.LOG.Info("amend-commit-success", zap.String("hash", commitHash.String()))
		return G.checkCommitHash(commitHash)
	}
}

// checkCommitHash checks and logs the commit hash.
func (G *Client) checkCommitHash(commitHash plumbing.Hash) (string, error) {
	object, err := G.repo.CommitObject(commitHash)
	if err != nil {
		return "", erero.Wro(err)
	}
	zaplog.ZAPS.Skip2.SUG.Info(object)
	must.Same(commitHash, object.Hash)
	return object.Hash.String(), nil
}

// IsLatestCommitPushedToRemote 检查当前分支是否已经推送到指定的远程仓库
func (G *Client) IsLatestCommitPushedToRemote(remoteName string) (bool, error) {
	// 获取当前分支引用
	branchReference := done.VCE(G.repo.Head()).Nice()
	// 获取远程引用
	remoteReference, err := G.repo.Reference(plumbing.ReferenceName(fmt.Sprintf("refs/remotes/%s/%s", remoteName, branchReference.Name().Short())), false)
	if err != nil {
		if errors.Is(err, plumbing.ErrReferenceNotFound) {
			return false, nil //报错通常是因为远程引用未找到，这里的错误是有预期的
		}
		return false, erero.Wro(err) // 其它错误
	}
	// 如果当前分支的提交哈希和远程分支相同，说明本地分支已经推送到远程
	return branchReference.Hash() == remoteReference.Hash(), nil
}

// IsLatestCommitPushed 检查当前分支是否已经推送到任何远程仓库
func (G *Client) IsLatestCommitPushed() (bool, error) {
	remotes, err := G.repo.Remotes()
	if err != nil {
		return false, erero.Wro(err)
	}
	for _, remote := range remotes {
		remoteName := remote.Config().Name

		if matched, err := G.IsLatestCommitPushedToRemote(remoteName); err != nil {
			return false, erero.Wro(err)
		} else if matched {
			return true, nil
		}
	}
	return false, nil // 没有任何远程仓库有当前分支的提交
}
