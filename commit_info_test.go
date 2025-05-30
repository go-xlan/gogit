package gogit

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

func TestPackageName(t *testing.T) {
	require.Equal(t, packageName, syntaxgo.CurrentPackageName())
}

func TestPackagePath(t *testing.T) {
	require.Equal(t, packagePath, syntaxgo_reflect.GetPkgPathV4(&CommitInfo{}))
}

func TestCommitInfo_GetCommitMessage(t *testing.T) {
	commitInfo := CommitInfo{Message: "example"}

	require.Equal(t, "example", commitInfo.BuildCommitMessage())
}

func TestCommitInfo_GetSignatureInfo(t *testing.T) {
	commitInfo := CommitInfo{
		Name:    "John Doe",
		Eddress: "johndoe@example.com",
	}
	authorInfo := commitInfo.GetObjectSignature()

	require.Equal(t, "John Doe", authorInfo.Name)
	require.Equal(t, "johndoe@example.com", authorInfo.Email)

	t.Log(neatjsons.S(authorInfo))
}

func TestCommitInfo_CheckFullMessage(t *testing.T) {
	commitInfo := CommitInfo{
		Name:    "Jane Doe",
		Eddress: "janedoe@example.com",
		Message: "example",
	}

	authorInfo := commitInfo.GetObjectSignature()

	require.Equal(t, "Jane Doe", authorInfo.Name)
	require.Equal(t, "janedoe@example.com", authorInfo.Email)
	require.Equal(t, "example", commitInfo.BuildCommitMessage())

	t.Log(neatjsons.S(authorInfo))
}

func TestCommitInfo_CheckNoneMessage(t *testing.T) {
	commitInfo := CommitInfo{}

	t.Log(commitInfo.BuildCommitMessage())

	authorInfo := commitInfo.GetObjectSignature()

	require.Equal(t, "gogit", authorInfo.Name)
	require.Equal(t, "gogit@github.com/go-xlan/gogit", authorInfo.Email)
	require.Contains(t, commitInfo.BuildCommitMessage(), packagePath)

	t.Log(neatjsons.S(authorInfo))
}
