package gogitv5git

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
func (T *Client88Must) CmtAll(options CommitMessage) (res string) {
	res, err1 := T.G.CmtAll(options)
	sure.Must(err1)
	return res
}
func (T *Client88Must) CAmend(options CommitMessage) (res string) {
	res, err1 := T.G.CAmend(options)
	sure.Must(err1)
	return res
}
