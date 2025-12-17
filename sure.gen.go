// Code generated using sure/sure_cls_gen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/sure
// Generated from: sure.gen_test.go:35 -> gogit.TestGen
// ========== SURE:DO-NOT-EDIT-SECTION:END ==========

package gogit

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/yyle88/sure"
)

type Client88Must struct{ c *Client }

func (c *Client) Must() *Client88Must {
	return &Client88Must{c: c}
}
func (T *Client88Must) Repo() (res *git.Repository) {
	res = T.c.Repo()
	return res
}
func (T *Client88Must) Tree() (res *git.Worktree) {
	res = T.c.Tree()
	return res
}
func (T *Client88Must) AddAll() {
	err := T.c.AddAll()
	sure.Must(err)
}
func (T *Client88Must) Status() (res git.Status) {
	res, err1 := T.c.Status()
	sure.Must(err1)
	return res
}
func (T *Client88Must) CommitAll(info *CommitInfo) (res string) {
	res, err1 := T.c.CommitAll(info)
	sure.Must(err1)
	return res
}
func (T *Client88Must) AmendCommit(cfg *AmendConfig) (res string) {
	res, err1 := T.c.AmendCommit(cfg)
	sure.Must(err1)
	return res
}
func (T *Client88Must) IsLatestCommitPushedToRemote(remoteName string) (res bool) {
	res, err1 := T.c.IsLatestCommitPushedToRemote(remoteName)
	sure.Must(err1)
	return res
}
func (T *Client88Must) IsLatestCommitPushed() (res bool) {
	res, err1 := T.c.IsLatestCommitPushed()
	sure.Must(err1)
	return res
}
func (T *Client88Must) GetCurrentBranch() (res string) {
	res, err1 := T.c.GetCurrentBranch()
	sure.Must(err1)
	return res
}
func (T *Client88Must) GetLatestCommit() (res *object.Commit) {
	res, err1 := T.c.GetLatestCommit()
	sure.Must(err1)
	return res
}
func (T *Client88Must) HasChanges() (res bool) {
	res, err1 := T.c.HasChanges()
	sure.Must(err1)
	return res
}
func (T *Client88Must) GetRemoteURL(remoteName string) (res string) {
	res, err1 := T.c.GetRemoteURL(remoteName)
	sure.Must(err1)
	return res
}
func (T *Client88Must) GetFirstRemoteURL() (res string) {
	res, err1 := T.c.GetFirstRemoteURL()
	sure.Must(err1)
	return res
}
