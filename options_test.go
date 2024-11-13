package gogitv5git

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/syntaxgo"
)

func TestCommitMessage_Signature(t *testing.T) {
	options := CommitMessage{}
	t.Log(options.CmMessage())

	t.Log(neatjsons.S(options.Signature()))
}

func TestMessage(t *testing.T) {
	require.Equal(t, pkgName, syntaxgo.CurrentPackageName())
}
