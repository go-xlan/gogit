package gogitv5git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-xlan/gogitv5git/gogitv5ops"
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
	zaplog.ZAPS.P1.SUG.Info("commit: ", "msg: ", message)

	return G.verifyCommitHash(G.worktree.Commit(message, &git.CommitOptions{
		All:    true, // Whether to commit deleted files. Usually true, as it's rare to commit without deleting.
		Author: commitInfo.GetSignatureInfo(),
	}))
}

// AmendCommit amends the previous commit with the new message, or uses the latest commit message if empty.
func (G *Client) AmendCommit(commitInfo *CommitInfo) (string, error) {
	message := commitInfo.Message
	if message == "" { // If no message is provided, use the latest commit message to amend.
		preReference := done.VCE(G.repository.Head()).Nice()
		commitObject := done.VCE(G.repository.CommitObject(preReference.Hash())).Nice()
		message = zerotern.VF(commitObject.Message, func() string {
			return commitInfo.GetCommitMessage()
		})
	}
	zaplog.ZAPS.P1.SUG.Info("amend: ", "msg: ", message)

	return G.verifyCommitHash(G.worktree.Commit(message, &git.CommitOptions{
		Author: commitInfo.GetSignatureInfo(),
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
	zaplog.ZAPS.P2.LOG.Info("commit", zap.String("hash", commitHash.String()))

	commitItem, err := G.repository.CommitObject(commitHash)
	if err != nil {
		return "", erero.Wro(err)
	}
	zaplog.ZAPS.P2.SUG.Info(commitItem)
	return commitHash.String(), nil
}
