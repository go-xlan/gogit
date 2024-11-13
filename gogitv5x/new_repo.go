package gogitv5x

import (
	"github.com/go-git/go-git/v5"
	"github.com/yyle88/erero"
)

func NewRepo(root string) (*git.Repository, error) {
	repo, err := git.PlainOpen(root)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return repo, nil
}
