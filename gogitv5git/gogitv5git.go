package gogitv5git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-xlan/gogitv5acp"
	"github.com/pkg/errors"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/must"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type Client struct {
	repo *git.Repository
}

func New(root string) (*Client, error) {
	repo, err := gogitv5acp.NewRepo(root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return NewClient(repo)
}

func NewClient(repo *git.Repository) (*Client, error) {
	worktree, err := repo.Worktree()
	if err != nil {
		return nil, erero.Wro(err) //in case of bare repo 裸仓库
	}
	must.OK(worktree) //meaningless

	return &Client{
		repo: repo,
	}, nil
}

func (G *Client) worktree() *git.Worktree {
	return done.VCE(G.repo.Worktree()).Nice()
}

func (G *Client) AddAll() error {
	worktree := G.worktree()
	if err := worktree.AddWithOptions(&git.AddOptions{All: true}); err != nil {
		return erero.Wro(err)
	}
	return nil
}

func (G *Client) Status() (git.Status, error) {
	worktree := G.worktree()
	status, err := worktree.Status() // We can verify the current status of the worktree using the method Status.
	if err != nil {
		return nil, erero.Wro(err)
	}
	return status, nil
}

func (G *Client) Commit(options *CommitOptions) (string, error) {
	msg := options.newMessage()
	zaplog.ZAPS.P1.SUG.Info("commit: ", "msg: ", msg)

	worktree := G.worktree()
	return G.seeCommit(worktree.Commit(msg, &git.CommitOptions{
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

	worktree := G.worktree()
	return G.seeCommit(worktree.Commit(msg, &git.CommitOptions{
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

	if commitItem, err := G.repo.CommitObject(commitHash); err != nil {
		return "", erero.Wro(err)
	} else {
		zaplog.ZAPS.P2.SUG.Info(commitItem)
	}
	return commitHash.String(), nil
}
