package gogitv5git

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-xlan/gogitv5git/internal/utils"
)

const pkgName = "gogitv5git"

type CommitMessage struct {
	Name    string
	Emails  string
	Message string
}

func (cmm *CommitMessage) name() string {
	return utils.SOrX(cmm.Name, pkgName)
}

func (cmm *CommitMessage) EmailsAtc() string {
	return utils.SOrX(cmm.Emails, pkgName+"@github.com")
}

func (cmm *CommitMessage) CmMessage() string {
	return utils.SOrR(cmm.Message, func() string {
		return fmt.Sprintf(`git commit -m "%s %s"`, pkgName, time.Now().Format("2006-01-02 15:04:05"))
	})
}

func (cmm *CommitMessage) Signature() *object.Signature {
	return &object.Signature{
		Name:  cmm.name(),
		Email: cmm.EmailsAtc(),
		When:  time.Now(),
	}
}
