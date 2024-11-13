package gogitv5x

import (
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
)

func NewRepoWorktreeWithIgnore(root string) (repo *git.Repository, worktree *git.Worktree, err error) {
	repo, err = NewRepo(root)
	if err != nil {
		return nil, nil, erero.Wro(err)
	}
	worktree, err = repo.Worktree()
	if err != nil {
		return nil, nil, erero.Wro(err)
	}

	//最低优先级最不常用
	SetIgnorePatterns(worktree, done.VAE(gitignore.LoadSystemPatterns(osfs.New("/"))).Done())
	//中等优先级比较常用
	SetIgnorePatterns(worktree, done.VAE(gitignore.LoadGlobalPatterns(osfs.New("/"))).Done())
	//最高优先级非常常用
	SetIgnorePatterns(worktree, done.VAE(GetProjectIgnorePatterns(root)).Done())

	return repo, worktree, nil
}

func NewWorktreeWithIgnore(root string) (worktree *git.Worktree, err error) {
	_, worktree, err = NewRepoWorktreeWithIgnore(root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return worktree, nil
}
