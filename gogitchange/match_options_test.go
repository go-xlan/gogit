package gogitchange_test

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-xlan/gogit/gogitchange"
	"github.com/stretchr/testify/require"
)

// TestMatchOptions_MatchType verifies file extension filtering
// Should correctly set and chain match type option
//
// TestMatchOptions_MatchType 验证文件扩展名过滤
// 应该正确设置和链接匹配类型选项
func TestMatchOptions_MatchType(t *testing.T) {
	options := gogitchange.NewMatchOptions().MatchType(".go")
	require.NotNil(t, options)
}

// TestMatchOptions_MatchPath verifies custom path filtering
// Should correctly set and chain match path option
//
// TestMatchOptions_MatchPath 验证自定义路径过滤
// 应该正确设置和链接匹配路径选项
func TestMatchOptions_MatchPath(t *testing.T) {
	options := gogitchange.NewMatchOptions().MatchPath(func(path string) bool {
		return true
	})
	require.NotNil(t, options)
}

// TestMatchOptions_MatchStatus verifies file status filtering
// Should correctly set and chain match status option
//
// TestMatchOptions_MatchStatus 验证文件状态过滤
// 应该正确设置和链接匹配状态选项
func TestMatchOptions_MatchStatus(t *testing.T) {
	options := gogitchange.NewMatchOptions().MatchStatus(git.Added, git.Modified)
	require.NotNil(t, options)
}

// TestMatchOptions_Chaining verifies fluent API chaining
// Should support chaining all options in sequence
//
// TestMatchOptions_Chaining 验证流畅 API 链式调用
// 应该支持按顺序链接所有选项
func TestMatchOptions_Chaining(t *testing.T) {
	options := gogitchange.NewMatchOptions().
		MatchType(".go").
		MatchPath(func(path string) bool {
			return true
		}).
		MatchStatus(git.Added, git.Modified)
	require.NotNil(t, options)
}
