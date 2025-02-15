package gogitassist

import (
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func NewRepo(root string) (*git.Repository, error) {
	repo, err := git.PlainOpen(root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return repo, nil
}

func NewRepoTreeWithIgnore(root string) (repo *git.Repository, tree *git.Worktree, err error) {
	repo, err = NewRepo(root)
	if err != nil {
		return nil, nil, erero.Wro(err)
	}
	tree, err = repo.Worktree()
	if err != nil {
		return nil, nil, erero.Wro(err)
	}

	//最低优先级最不常用
	SetIgnorePatterns(tree, done.VAE(gitignore.LoadSystemPatterns(osfs.New("/"))).Done())
	//中等优先级比较常用
	SetIgnorePatterns(tree, done.VAE(gitignore.LoadGlobalPatterns(osfs.New("/"))).Done())
	//最高优先级非常常用
	SetIgnorePatterns(tree, done.VAE(LoadProjectIgnorePatterns(root)).Done())

	return repo, tree, nil
}

func DebugRepo(repo *git.Repository) {
	topReference, err := repo.Head()
	if err != nil {
		zaplog.LOG.Error("wrong-when-get-head-reference", zap.Error(err))
		return
	}
	zaplog.SUG.Debugln("commit", topReference.String())

	commitObject, err := repo.CommitObject(topReference.Hash())
	if err != nil {
		zaplog.LOG.Error("wrong-when-get-commit-object", zap.Error(err))
		return
	}
	zaplog.SUG.Debugln(commitObject.String())
	zaplog.SUG.Debugln("return", "debug-repo-function-return")
}
