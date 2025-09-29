package gogitassist_test

import (
	"testing"

	"github.com/go-xlan/gogit/gogitassist"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/runpath"
)

// TestLoadProjectIgnorePatterns tests loading ignore patterns from project .gitignore
// Verifies function can load patterns from project root .gitignore file
//
// TestLoadProjectIgnorePatterns 测试从项目 .gitignore 加载忽略模式
// 验证函数能够从项目根 .gitignore 文件加载模式
func TestLoadProjectIgnorePatterns(t *testing.T) {
	root := runpath.PARENT.Up(1)

	patterns, err := gogitassist.LoadProjectIgnorePatterns(root)
	require.NoError(t, err)
	require.NotEmpty(t, patterns)

	t.Log(len(patterns))
}

// TestLoadIgnorePatternsFromPath tests loading ignore patterns from specific path
// Verifies function can parse .gitignore file at specified location
//
// TestLoadIgnorePatternsFromPath 测试从特定路径加载忽略模式
// 验证函数能够解析指定位置的 .gitignore 文件
func TestLoadIgnorePatternsFromPath(t *testing.T) {
	path := runpath.PARENT.UpTo(1, ".gitignore")

	patterns, err := gogitassist.LoadIgnorePatternsFromPath(path)
	require.NoError(t, err)
	require.NotEmpty(t, patterns)

	t.Log(len(patterns))
}
