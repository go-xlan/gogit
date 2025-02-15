package gogitassist_test

import (
	"testing"

	"github.com/go-xlan/gogit/gogitassist"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestNewRepo(t *testing.T) {
	repo, err := gogitassist.NewRepo(runpath.PARENT.Up(1))
	require.NoError(t, err)
	tags, err := repo.Tags()
	require.NoError(t, err)
	t.Log(neatjsons.S(tags))
}

func TestNewRepoTreeWithIgnore(t *testing.T) {
	repo, tree, err := gogitassist.NewRepoTreeWithIgnore(runpath.PARENT.Up(1))
	require.NoError(t, err)
	gogitassist.DebugRepo(repo)

	status, err := tree.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}
