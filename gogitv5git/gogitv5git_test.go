package gogitv5git

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestNew(t *testing.T) {
	root := filepath.Dir(runpath.PARENT.Path())
	client, err := New(root)
	require.NoError(t, err)

	status, err := client.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}

func TestClient_Commit(t *testing.T) {
	root := filepath.Dir(runpath.PARENT.Path())
	client := done.VCE(New(root)).Nice()

	if false { //not commit in this test case
		err := client.AddAll()
		require.NoError(t, err)

		status, err := client.Status()
		require.NoError(t, err)

		t.Log(neatjsons.S(status))

		commitHash, err := client.Commit(CommitOptions{})
		require.NoError(t, err)
		t.Log(commitHash)
	}
}
