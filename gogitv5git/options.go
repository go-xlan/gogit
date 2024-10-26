package gogitv5git

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-xlan/gogitv5acp/internal/utils"
)

type CommitOptions struct {
	Name    string
	Emails  string
	Message string
}

func (options *CommitOptions) newName() string {
	return utils.SOrX(options.Name, "gogitv5acp")
}

func (options *CommitOptions) newEmails() string {
	return utils.SOrX(options.Emails, "gogitv5acp@github.com")
}

func (options *CommitOptions) newMessage() string {
	return utils.SOrR(options.Message, func() string {
		return fmt.Sprintf(`git commit -m "%s %s"`, "gogitv5acp", time.Now().Format("2006-01-02 15:04:05"))
	})
}

func (options *CommitOptions) newAuthors() *object.Signature {
	return &object.Signature{
		Name:  options.newName(),
		Email: options.newEmails(),
		When:  time.Now(),
	}
}
