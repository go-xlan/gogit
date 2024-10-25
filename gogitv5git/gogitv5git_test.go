package gogitv5git

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
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
