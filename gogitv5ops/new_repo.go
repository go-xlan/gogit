package gogitv5ops

import (
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
)

func NewRepo(root string) (*git.Repository, error) {
	repository, err := git.PlainOpen(root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return repository, nil
}

func NewRepoWorktreeWithIgnore(root string) (repository *git.Repository, worktree *git.Worktree, err error) {
	repository, err = NewRepo(root)
	if err != nil {
		return nil, nil, erero.Wro(err)
	}
	worktree, err = repository.Worktree()
	if err != nil {
		return nil, nil, erero.Wro(err)
	}

	//最低优先级最不常用
	SetIgnorePatterns(worktree, done.VAE(gitignore.LoadSystemPatterns(osfs.New("/"))).Done())
	//中等优先级比较常用
	SetIgnorePatterns(worktree, done.VAE(gitignore.LoadGlobalPatterns(osfs.New("/"))).Done())
	//最高优先级非常常用
	SetIgnorePatterns(worktree, done.VAE(LoadProjectIgnorePatterns(root)).Done())

	return repository, worktree, nil
}
