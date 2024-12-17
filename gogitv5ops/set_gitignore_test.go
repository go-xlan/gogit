package gogitv5ops

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/runpath"
)

func TestLoadProjectIgnorePatterns(t *testing.T) {
	root := runpath.PARENT.Up(1)

	patterns, err := LoadProjectIgnorePatterns(root)
	require.NoError(t, err)
	require.NotEmpty(t, patterns)

	t.Log(len(patterns))
}

func TestLoadIgnorePatternsFromPath(t *testing.T) {
	path := runpath.PARENT.UpTo(1, ".gitignore")

	patterns, err := LoadIgnorePatternsFromPath(path)
	require.NoError(t, err)
	require.NotEmpty(t, patterns)

	t.Log(len(patterns))
}
