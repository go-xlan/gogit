package gogit

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

func TestCommitInfo_GetCommitMessage(t *testing.T) {
	commitInfo := CommitInfo{Message: "example"}

	require.Equal(t, "example", commitInfo.GetCommitMessage())
}

func TestCommitInfo_GetSignatureInfo(t *testing.T) {
	commitInfo := CommitInfo{
		Name:    "John Doe",
		Eddress: "johndoe@example.com",
	}
	objectSignature := commitInfo.GetObjectSignature()

	require.Equal(t, "John Doe", objectSignature.Name)
	require.Equal(t, "johndoe@example.com", objectSignature.Email)

	t.Log(neatjsons.S(objectSignature))
}

func TestPackageName(t *testing.T) {
	require.Equal(t, packageName, syntaxgo.CurrentPackageName())
}

func TestPackagePath(t *testing.T) {
	require.Equal(t, packagePath, syntaxgo_reflect.GetPkgPathV4(&CommitInfo{}))
}

func TestCommitInfo_CheckFullMessage(t *testing.T) {
	commitInfo := CommitInfo{
		Name:    "Jane Doe",
		Eddress: "janedoe@example.com",
		Message: "example",
	}

	objectSignature := commitInfo.GetObjectSignature()

	require.Equal(t, "Jane Doe", objectSignature.Name)
	require.Equal(t, "janedoe@example.com", objectSignature.Email)
	require.Equal(t, "example", commitInfo.GetCommitMessage())

	t.Log(neatjsons.S(objectSignature))
}

func TestCommitInfo_CheckNoneMessage(t *testing.T) {
	commitInfo := CommitInfo{}

	t.Log(commitInfo.GetCommitMessage())

	objectSignature := commitInfo.GetObjectSignature()

	require.Equal(t, "gogit", objectSignature.Name)
	require.Equal(t, "gogit@github.com/go-xlan/gogit", objectSignature.Email)
	require.Contains(t, commitInfo.GetCommitMessage(), packagePath)

	t.Log(neatjsons.S(objectSignature))
}
