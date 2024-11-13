package gogitv5git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-xlan/gogitv5git/gogitv5x"
	"github.com/go-xlan/gogitv5git/internal/utils"
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
	repo, worktree, err := gogitv5x.NewRepoWorktreeWithIgnore(root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return NewClient(repo, worktree)
}

func NewClient(repo *git.Repository, worktree *git.Worktree) (*Client, error) {
	return &Client{
		repo:     repo,
		worktree: worktree, //这里最好是使用已经设置 ignore 的对象，否则有可能会提交些没用的东西
	}, nil
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

func (G *Client) Commit(options CommitOptions) (string, error) {
	msg := options.CmMessage()
	zaplog.ZAPS.P1.SUG.Info("commit: ", "msg: ", msg)

	return G.seeCommit(G.worktree.Commit(msg, &git.CommitOptions{
		All:    true, //是否提交已经删除的东西，通常是true的，毕竟不提交删除的场景，基本是没有的
		Author: options.Signature(),
	}))
}

func (G *Client) CAmend(options CommitOptions) (string, error) {
	msg := options.Message
	if msg == "" { //当不填的时候就需要从最新一次的commit里拿到信息，这样才叫追加内容的提交
		reference := done.VCE(G.repo.Head()).Nice()
		commit := done.VCE(G.repo.CommitObject(reference.Hash())).Nice()
		msg = utils.SOrR(commit.Message, func() string {
			return options.CmMessage()
		})
	}
	zaplog.ZAPS.P1.SUG.Info("camend: ", "msg: ", msg)

	return G.seeCommit(G.worktree.Commit(msg, &git.CommitOptions{
		Author: options.Signature(),
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
