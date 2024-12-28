package gogit_test

import (
	"testing"

	"github.com/go-xlan/gogit"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestNew(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)

	status, err := client.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}

func TestClient_CommitAll(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)

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

func TestClient_IsHashMatchedRemote(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)

	matched, err := client.IsHashMatchedRemote("origin")
	require.NoError(t, err)
	t.Log(matched)
}

func TestClient_IsHashMatchedRemote_NotExist(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)

	matched, err := client.IsHashMatchedRemote("origin2")
	require.NoError(t, err)
	t.Log(matched)
}

func TestClient_IsPushedToAnyRemote(t *testing.T) {
	client, err := gogit.New(runpath.PARENT.Path())
	require.NoError(t, err)

	pushed, err := client.IsPushedToAnyRemote()
	require.NoError(t, err)
	t.Log(pushed)
}
