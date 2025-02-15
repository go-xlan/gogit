package gogit

import (
	"github.com/go-git/go-git/v5"
	"github.com/yyle88/sure"
)

type Client88Must struct{ G *Client }

func (G *Client) Must() *Client88Must {
	return &Client88Must{G: G}
}
func (T *Client88Must) Repo() (res *git.Repository) {
	res = T.G.Repo()
	return res
}
func (T *Client88Must) Tree() (res *git.Worktree) {
	res = T.G.Tree()
	return res
}
func (T *Client88Must) AddAll() {
	err := T.G.AddAll()
	sure.Must(err)
}
func (T *Client88Must) Status() (res git.Status) {
	res, err1 := T.G.Status()
	sure.Must(err1)
	return res
}
func (T *Client88Must) CommitAll(info *CommitInfo) (res string) {
	res, err1 := T.G.CommitAll(info)
	sure.Must(err1)
	return res
}
func (T *Client88Must) AmendCommit(cfg *AmendConfig) (res string) {
	res, err1 := T.G.AmendCommit(cfg)
	sure.Must(err1)
	return res
}
func (T *Client88Must) IsLatestCommitPushedToRemote(remoteName string) (res bool) {
	res, err1 := T.G.IsLatestCommitPushedToRemote(remoteName)
	sure.Must(err1)
	return res
}
func (T *Client88Must) IsLatestCommitPushed() (res bool) {
	res, err1 := T.G.IsLatestCommitPushed()
	sure.Must(err1)
	return res
}
