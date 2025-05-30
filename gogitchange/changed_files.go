package gogitchange

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/yyle88/erero"
	"github.com/yyle88/must"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/osexistpath/ossoftexist"
)

type ChangedFileManager struct {
	projectPath string
	tree        *git.Worktree
}

func NewChangedFileManager(projectPath string, worktree *git.Worktree) *ChangedFileManager {
	return &ChangedFileManager{
		projectPath: osmustexist.ROOT(must.Nice(projectPath)),
		tree:        worktree,
	}
}

// Foreach 找到变化的文件（除了删除的）把变动的文件格式化再提交
func (m *ChangedFileManager) Foreach(matchOptions *MatchOptions, process func(path string) error) error {
	statusMap, err := m.tree.Status()
	if err != nil {
		return erero.Wro(err)
	}

	for relativePath, status := range statusMap {
		// 需要过滤掉已经删除的
		if status.Staging == git.Deleted {
			continue
		}

		// 过滤掉扩展名不匹配的
		if matchOptions.matchType != "" && filepath.Ext(relativePath) != matchOptions.matchType {
			continue
		}

		// 当需要过滤路径时，就可以通过这个函数过滤，把不需要处理的路径排除掉
		if matchOptions.matchPath != nil && !matchOptions.matchPath(filepath.Join(m.projectPath, relativePath)) {
			continue
		}

		// 拼接出文件的路径
		path := filepath.Join(m.projectPath, relativePath)

		// 当需要对这个文件执行特殊操作时，把操作传进来起，操作可以是修改这个文件的内容，这时就得要求这个文件是存在的，而不是被删除的，或者不存在的
		if ossoftexist.IsFile(path) {
			//这里的操作可以是打印文件内容，也可以是修改文件内容-比如替换文件内容-比如格式化go代码内容
			if err := process(path); err != nil {
				return erero.Wro(err)
			}
		}
	}
	return nil
}

func (m *ChangedFileManager) ListChangedFilePaths(matchOptions *MatchOptions) ([]string, error) {
	var paths = make([]string, 0)
	if err := m.Foreach(matchOptions, func(path string) error {
		paths = append(paths, path)
		return nil
	}); err != nil {
		return nil, erero.Wro(err)
	}
	return paths, nil
}

// FormatChangedGoFiles 格式化所有变化的 Go 代码文件。(该功能已经移到别的模块里)
// func (m *ChangedFileManager) FormatChangedGoFiles(matchOptions *MatchOptions) error {
// code move to -> https://github.com/go-mate/go-commit
// }

// ForeachChangedGoFile 遍历所有变化的 Go 代码文件，再对每个文件执行指定的处理函数
func (m *ChangedFileManager) ForeachChangedGoFile(matchOptions *MatchOptions, process func(path string) error) error {
	if err := m.Foreach(matchOptions, func(path string) error {
		if filepath.Ext(path) != ".go" {
			return nil
		}
		if err := process(path); err != nil {
			return erero.Wro(err)
		}
		return nil
	}); err != nil {
		return erero.Wro(err)
	}
	return nil
}
