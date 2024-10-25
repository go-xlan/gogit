package gogitv5git

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-xlan/gogitv5acp"
	"github.com/go-xlan/gogitv5acp/internal/utils"
	"github.com/pkg/errors"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type Client struct {
	repo     *git.Repository
	worktree *git.Worktree
}

func New(root string) (*Client, error) {
	repo, worktree, err := gogitv5acp.NewRepoWorktree(root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return NewClient(repo, worktree), nil
}

func NewClient(repo *git.Repository, worktree *git.Worktree) *Client {
	return &Client{
		repo:     repo,
		worktree: worktree,
	}
}

func (G *Client) AddAll() error {
	if err := G.worktree.AddWithOptions(&git.AddOptions{All: true}); err != nil {
		return erero.Wro(err)
	}
	return nil
}

func (G *Client) Status() (git.Status, error) {
	status, err := G.worktree.Status() // We can verify the current status of the worktree using the method Status.
	if err != nil {
		return nil, erero.Wro(err)
	}
	return status, nil
}

type CommitOptions struct {
	Name    string
	Emails  string
	Message string
}

func (options *CommitOptions) newName() string {
	return utils.SOrX(options.Name, "gogitv5acp")
}

func (options *CommitOptions) newEmails() string {
	return utils.SOrX(options.Emails, "gogitv5acp@github.com")
}

func (options *CommitOptions) newMessage() string {
	return utils.SOrX(options.Message, fmt.Sprintf(`git commit -m "%s %s"`, "gogitv5acp", time.Now().Format("2006-01-02 15:04:05")))
}

func (options *CommitOptions) newAuthors() *object.Signature {
	return &object.Signature{
		Name:  options.newName(),
		Email: options.newEmails(),
		When:  time.Now(),
	}
}

func (G *Client) Commit(options *CommitOptions) (string, error) {
	msg := options.newMessage()
	zaplog.ZAPS.P1.SUG.Info("commit: ", "msg: ", msg)

	return G.seeCommit(G.worktree.Commit(msg, &git.CommitOptions{
		All:    true, //是否提交已经删除的东西，通常是true的，毕竟不提交删除的场景，基本是没有的
		Author: options.newAuthors(),
	}))
}

func (G *Client) CAmend(options *CommitOptions) (string, error) {
	msg := options.Message
	if msg == "" {
		reference := done.VCE(G.repo.Head()).Nice()
		commit := done.VCE(G.repo.CommitObject(reference.Hash())).Nice()
		msg = commit.Message
	}
	zaplog.ZAPS.P1.SUG.Info("camend: ", "msg: ", msg)

	return G.seeCommit(G.worktree.Commit(msg, &git.CommitOptions{
		Author: options.newAuthors(),
		Amend:  true, //注意这里有"all and amend cannot be used together"的限制，因此前面的"all"不要设置
	}))
}

func (G *Client) seeCommit(commitHash plumbing.Hash, erx error) (string, error) {
	if erx != nil {
		if errors.Is(erx, git.ErrEmptyCommit) {
			return "", nil
		}
		return "", erero.Wro(erx)
	}
	zaplog.ZAPS.P2.LOG.Info("commit", zap.String("hash", commitHash.String()))
	if commitObject, err := G.repo.CommitObject(commitHash); err != nil {
		return "", erero.Wro(err)
	} else {
		zaplog.ZAPS.P2.SUG.Info(commitObject)
	}
	return commitHash.String(), nil
}
