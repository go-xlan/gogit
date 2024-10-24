package gogitv5acp

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestNewWorktree(t *testing.T) {
	worktree, err := NewWorktree(runpath.PARENT.Path())
	require.NoError(t, err)

	status, err := worktree.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}
