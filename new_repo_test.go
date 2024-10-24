package gogitv5acp

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
)

func TestNewRepo(t *testing.T) {
	repo, err := NewRepo(runpath.PARENT.Path())
	require.NoError(t, err)
	tags, err := repo.Tags()
	require.NoError(t, err)
	t.Log(neatjsons.S(tags))
}
