package gogit

import (
	"github.com/go-git/go-git/v5"
	"github.com/yyle88/sure"
)

type Client88Must struct{ G *Client }

func (G *Client) Must() *Client88Must {
	return &Client88Must{G: G}
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
func (T *Client88Must) CommitAll(commitInfo *CommitInfo) (res string) {
	res, err1 := T.G.CommitAll(commitInfo)
	sure.Must(err1)
	return res
}
func (T *Client88Must) AmendCommit(amendConfig *AmendConfig) (res string) {
	res, err1 := T.G.AmendCommit(amendConfig)
	sure.Must(err1)
	return res
}
func (T *Client88Must) IsHashMatchedRemote(remoteName string) (res bool) {
	res, err1 := T.G.IsHashMatchedRemote(remoteName)
	sure.Must(err1)
	return res
}
func (T *Client88Must) IsPushedToAnyRemote() (res bool) {
	res, err1 := T.G.IsPushedToAnyRemote()
	sure.Must(err1)
	return res
}
