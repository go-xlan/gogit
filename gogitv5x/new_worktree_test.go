package gogitv5x

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestNewWorktree(t *testing.T) {
	worktree, err := NewWorktreeWithIgnore(runpath.PARENT.Up(1))
	require.NoError(t, err)

	status, err := worktree.Status()
	require.NoError(t, err)

	t.Log(neatjsons.S(status))
}
