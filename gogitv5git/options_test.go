package gogitv5git

import (
	"testing"

	"github.com/yyle88/neatjson/neatjsons"
)

func TestCommitOptions_Signature(t *testing.T) {
	options := CommitOptions{}
	t.Log(neatjsons.S(options.newAuthors()))
}
