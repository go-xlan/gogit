package gogitv5git

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-xlan/gogitv5git/internal/utils"
)

type CommitOptions struct {
	Name    string
	Emails  string
	Message string
}

func (options *CommitOptions) name() string {
	return utils.SOrX(options.Name, "gogitv5git")
}

func (options *CommitOptions) emails() string {
	return utils.SOrX(options.Emails, "gogitv5git@github.com")
}

func (options *CommitOptions) CmMessage() string {
	return utils.SOrR(options.Message, func() string {
		return fmt.Sprintf(`git commit -m "%s %s"`, "gogitv5git", time.Now().Format("2006-01-02 15:04:05"))
	})
}

func (options *CommitOptions) Signature() *object.Signature {
	return &object.Signature{
		Name:  options.name(),
		Email: options.emails(),
		When:  time.Now(),
	}
}
