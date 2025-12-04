package gogitassist

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"github.com/yyle88/erero"
	"github.com/yyle88/osexistpath/osomitexist"
)

// SetIgnorePatterns adds ignore patterns to worktree excludes
// Appends provided patterns to existing worktree exclusion rules
// Enables flexible ignore pattern management at runtime
//
// SetIgnorePatterns 向工作树排除规则中添加忽略模式
// 将提供的模式附加到现有工作树排除规则
// 在运行时实现灵活的忽略模式管理
func SetIgnorePatterns(worktree *git.Worktree, patterns []gitignore.Pattern) {
	worktree.Excludes = append(worktree.Excludes, patterns...)
}

// LoadProjectIgnorePatterns loads gitignore patterns from project .gitignore file
// Checks if root path exists and loads patterns from .gitignore file
// Returns blank patterns if no .gitignore file found
//
// LoadProjectIgnorePatterns 从项目 .gitignore 文件加载 gitignore 模式
// 检查根路径是否存在并从 .gitignore 文件加载模式
// 如果未找到 .gitignore 文件则返回空模式
func LoadProjectIgnorePatterns(root string) ([]gitignore.Pattern, error) {
	if osomitexist.IsRoot(root) {
		patterns, err := LoadIgnorePatternsFromPath(filepath.Join(root, ".gitignore"))
		if err != nil {
			return nil, erero.Wro(err)
		}
		return patterns, nil
	}
	return []gitignore.Pattern{}, nil
}

// LoadIgnorePatternsFromPath loads ignore patterns from specified file path
// Reads file content and parses gitignore patterns
// Returns blank patterns if file does not exist
//
// LoadIgnorePatternsFromPath 从指定文件路径加载忽略模式
// 读取文件内容并解析 gitignore 模式
// 如果文件不存在则返回空模式
func LoadIgnorePatternsFromPath(path string) ([]gitignore.Pattern, error) {
	if osomitexist.IsFile(path) {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, erero.Wro(err)
		}
		patterns, err := LoadIgnorePatternsFromText(string(data))
		if err != nil {
			return nil, erero.Wro(err)
		}
		return patterns, nil
	}
	return []gitignore.Pattern{}, nil
}

// LoadIgnorePatternsFromText parses gitignore patterns from text content
// Processes each line, ignoring comments and blank lines
// Returns parsed gitignore patterns that are usable
//
// LoadIgnorePatternsFromText 从文本内容解析 gitignore 模式
// 处理每一行，忽略注释和空行
// 返回可用的已解析 gitignore 模式
func LoadIgnorePatternsFromText(text string) ([]gitignore.Pattern, error) {
	var patterns = make([]gitignore.Pattern, 0)
	for _, stx := range strings.Split(text, "\n") {
		if stx = strings.TrimSpace(stx); stx != "" && !strings.HasPrefix(stx, "#") {
			patterns = append(patterns, gitignore.ParsePattern(stx, []string{}))
		}
	}
	return patterns, nil
}
