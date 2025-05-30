package gogit_test

import (
	"testing"

	"github.com/go-xlan/gogit"
	"github.com/go-xlan/gogit/gogitassist"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
)

func TestNew(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)
	gogitassist.DebugRepo(client.Repo())

	status, err := client.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}

func TestNewClient(t *testing.T) {
	client := rese.P1(gogit.New(runpath.PARENT.Path()))
	gogitassist.DebugRepo(client.Repo())

	status, err := client.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}

func TestClient_CommitAll(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)
	gogitassist.DebugRepo(client.Repo())

	if false { //not commit in this test case
		err := client.AddAll()
		require.NoError(t, err)

		status, err := client.Status()
		require.NoError(t, err)

		t.Log(neatjsons.S(status))

		commitHash, err := client.CommitAll(gogit.NewCommitInfo("提交代码"))
		require.NoError(t, err)
		t.Log(commitHash)
	}
}

func TestClient_IsLatestCommitPushedToRemote(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)
	gogitassist.DebugRepo(client.Repo())

	matched, err := client.IsLatestCommitPushedToRemote("origin")
	require.NoError(t, err)
	t.Log(matched)
}

func TestClient_IsLatestCommitPushedToRemote_NotExist(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)
	gogitassist.DebugRepo(client.Repo())

	matched, err := client.IsLatestCommitPushedToRemote("origin2")
	require.NoError(t, err)
	t.Log(matched)
	require.False(t, matched)
}

func TestClient_IsLatestCommitPushed(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)
	gogitassist.DebugRepo(client.Repo())

	pushed, err := client.IsLatestCommitPushed()
	require.NoError(t, err)
	t.Log(pushed)
}
