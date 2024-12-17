package gogitv5ops

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestNewRepo(t *testing.T) {
	repo, err := NewRepo(runpath.PARENT.Up(1))
	require.NoError(t, err)
	tags, err := repo.Tags()
	require.NoError(t, err)
	t.Log(neatjsons.S(tags))
}

func TestNewRepoWorktreeWithIgnore(t *testing.T) {
	repository, worktree, err := NewRepoWorktreeWithIgnore(runpath.PARENT.Up(1))
	require.NoError(t, err)

	checkRepoHeadReference(t, repository)

	status, err := worktree.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}

func checkRepoHeadReference(t *testing.T, repository *git.Repository) {
	head, err := repository.Head()
	require.NoError(t, err)
	t.Log(head.String())

	commitObject, err := repository.CommitObject(head.Hash())
	require.NoError(t, err)
	t.Log(commitObject.String())
}
