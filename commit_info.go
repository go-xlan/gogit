package gogit

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/yyle88/eroticgo"
	"github.com/yyle88/tern/zerotern"
	"github.com/yyle88/zaplog"
)

const packageName = "gogit"
const packagePath = "github.com/go-xlan/gogit"

type CommitInfo struct {
	Name    string
	Eddress string
	Message string
}

// NewCommitInfo creates a new CommitInfo instance with default values.
func NewCommitInfo(message string) *CommitInfo {
	return &CommitInfo{
		Message: message,
	}
}

// WithName sets the Name field of the CommitInfo and returns the updated instance.
func (c *CommitInfo) WithName(name string) *CommitInfo {
	c.Name = name
	return c
}

// WithEddress sets the Eddress field of the CommitInfo and returns the updated instance.
func (c *CommitInfo) WithEddress(eddress string) *CommitInfo {
	c.Eddress = eddress
	return c
}

// WithMessage sets the Message field of the CommitInfo and returns the updated instance.
func (c *CommitInfo) WithMessage(message string) *CommitInfo {
	c.Message = message
	return c
}

func (c *CommitInfo) BuildCommitMessage() string {
	return zerotern.VF(c.Message, func() string {
		message := fmt.Sprintf(`git commit -m "[%s](%s) %s"`, packageName, packagePath, time.Now().Format("2006-01-02 15:04:05"))
		zaplog.SUG.Debugln(eroticgo.BLUE.Sprint(fmt.Sprintf(`git commit -m "%s"`, message)))
		return message
	})
}

func (c *CommitInfo) GetObjectSignature() *object.Signature {
	return &object.Signature{
		Name:  zerotern.VV(c.Name, packageName),
		Email: zerotern.VV(c.Eddress, fmt.Sprintf("%s@%s", packageName, packagePath)),
		When:  time.Now(),
	}
}
