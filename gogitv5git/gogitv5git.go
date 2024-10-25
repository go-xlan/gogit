package gogitv5git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-xlan/gogitv5acp"
	"github.com/yyle88/erero"
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
	return &Client{repo: repo, worktree: worktree}
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
