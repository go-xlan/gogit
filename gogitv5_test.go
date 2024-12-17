package gogitv5git_test

import (
	"testing"

	"github.com/go-xlan/gogitv5git"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestNew(t *testing.T) {
	root := runpath.PARENT.Path()
	client, err := gogitv5git.New(root)
	require.NoError(t, err)

	status, err := client.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}

func TestClient_CommitAll(t *testing.T) {
	root := runpath.PARENT.Path()
	client := done.VCE(gogitv5git.New(root)).Nice()

	if false { //not commit in this test case
		err := client.AddAll()
		require.NoError(t, err)

		status, err := client.Status()
		require.NoError(t, err)

		t.Log(neatjsons.S(status))

		commitHash, err := client.CommitAll(gogitv5git.NewCommitInfo("提交代码"))
		require.NoError(t, err)
		t.Log(commitHash)
	}
}
