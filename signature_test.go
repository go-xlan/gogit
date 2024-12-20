package gogit

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

func TestCommitInfo_GetCommitMessage(t *testing.T) {
	commitInfo := CommitInfo{Message: "Initial commit"}

	require.Equal(t, "Initial commit", commitInfo.GetCommitMessage())
}

func TestCommitInfo_GetSignatureInfo(t *testing.T) {
	commitInfo := CommitInfo{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
	signatureInfo := commitInfo.GetSignatureInfo()

	require.Equal(t, "John Doe", signatureInfo.Name)
	require.Equal(t, "johndoe@example.com", signatureInfo.Email)

	t.Log(neatjsons.S(signatureInfo))
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
		Email:   "janedoe@example.com",
		Message: "example",
	}

	signatureInfo := commitInfo.GetSignatureInfo()

	require.Equal(t, "Jane Doe", signatureInfo.Name)
	require.Equal(t, "janedoe@example.com", signatureInfo.Email)
	require.Equal(t, "example", commitInfo.GetCommitMessage())

	t.Log(neatjsons.S(signatureInfo))
}

func TestCommitInfo_CheckNoneMessage(t *testing.T) {
	commitInfo := CommitInfo{}

	t.Log(commitInfo.GetCommitMessage())
	t.Log(neatjsons.S(commitInfo.GetSignatureInfo()))
}
