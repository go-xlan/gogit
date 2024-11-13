package gogitv5x

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/runpath"
)

func TestGetProjectIgnorePatterns(t *testing.T) {
	root := runpath.PARENT.Path()

	patterns, err := GetProjectIgnorePatterns(root)
	require.NoError(t, err)

	t.Log(len(patterns))
}

func TestGetIgnorePatternsFromPath(t *testing.T) {
	path := runpath.PARENT.Join(".gitignore")

	patterns, err := GetIgnorePatternsFromPath(path)
	require.NoError(t, err)

	t.Log(len(patterns))
}
