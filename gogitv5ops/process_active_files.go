package gogitv5ops

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/yyle88/erero"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/must"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/osexistpath/ossoftexist"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type ProcessingOptions struct {
	projectRoot string // project location path. 目前还不知道怎么从他们给的结构里拿到，因此只能再次传进来，比较无奈呢

	matchWithExtension string // ".go" / ".txt". match the file name extension
	matchNoneExtension bool   // match path without extension.

	matchPath func(string) bool
}

func NewProcessingOptions(root string) *ProcessingOptions {
	return &ProcessingOptions{
		projectRoot: osmustexist.ROOT(must.Nice(root)),
	}
}

func (options *ProcessingOptions) WithFileExtension(matchWithExtension string) *ProcessingOptions {
	options.matchWithExtension = matchWithExtension
	return options
}

func (options *ProcessingOptions) WithNoneExtension(matchNoneExtension bool) *ProcessingOptions {
	options.matchNoneExtension = matchNoneExtension
	return options
}

func (options *ProcessingOptions) MatchPath(matchPath func(path string) bool) *ProcessingOptions {
	options.matchPath = matchPath
	return options
}

// ProcessActiveFiles 找到变化的文件（除了删除的）把变动的文件格式化再提交
func (options *ProcessingOptions) ProcessActiveFiles(worktree *git.Worktree, processingFunc func(path string) error) error {
	statusMap, err := worktree.Status()
	if err != nil {
		return erero.Wro(err)
	}

	for subPath, sts := range statusMap {
		// 需要过滤掉已经删除的
		if sts.Staging == git.Deleted {
			continue
		}

		// 过滤掉扩展名不匹配的
		if options.matchWithExtension != "" {
			if extension := filepath.Ext(subPath); extension != options.matchWithExtension {
				continue
			}
		}

		//当需要匹配不含扩展名的时，就用这个逻辑
		if options.matchNoneExtension {
			if extension := filepath.Ext(subPath); extension != "" {
				continue
			}
		}

		absPath := filepath.Join(options.projectRoot, subPath)

		//当需要过滤路径时，就可以通过这个函数过滤，把不需要处理的路径排除掉
		if options.matchPath != nil && !options.matchPath(absPath) {
			continue
		}

		//当需要对这个文件执行特殊操作时，把操作传进来起，操作可以是修改这个文件的内容，这时就得要求这个文件是存在的，而不是被删除的，或者不存在的
		if sts.Staging != git.Deleted && ossoftexist.IsFile(absPath) {
			//这里的操作可以是打印文件内容，也可以是修改文件内容-比如替换文件内容-比如格式化go代码内容
			if err := processingFunc(absPath); err != nil {
				return erero.Wro(err)
			}
		}
	}
	return nil
}

func (options *ProcessingOptions) CollectChangedFiles(worktree *git.Worktree) ([]string, error) {
	var activeFiles = make([]string, 0)
	if err := options.ProcessActiveFiles(worktree, func(path string) error {
		activeFiles = append(activeFiles, path)
		return nil
	}); err != nil {
		return nil, erero.Wro(err)
	}
	return activeFiles, nil
}

func (options *ProcessingOptions) FormatModifiedGoFiles(worktree *git.Worktree) error {
	if err := options.ProcessActiveFiles(worktree, func(path string) error {
		if extension := filepath.Ext(path); extension != ".go" {
			return nil
		}

		zaplog.ZAPS.Skip1.LOG.Info("golang-format-source", zap.String("path", path))

		if err := formatgo.FormatFile(path); err != nil {
			return erero.Wro(err)
		}
		return nil
	}); err != nil {
		return erero.Wro(err)
	}
	return nil
}
