package gogitv5git

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/sure_cls_gen"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

func TestGen(t *testing.T) {
	param := sure_cls_gen.NewGenParam(runpath.PARENT.Path())
	param.SetSubClassNamePartWords("88")
	param.SetSubClassNameStyleEnum(sure_cls_gen.STYLE_SUFFIX_CAMELCASE_TYPE)
	param.SetSureEnum(sure.MUST)

	cfg := &sure_cls_gen.Config{
		GenParam:      param,
		PkgName:       syntaxgo.CurrentPackageName(),
		ImportOptions: syntaxgo_ast.NewPackageImportOptions().SetInferredObject(git.Status{}),
		SrcPath:       runtestpath.SrcPath(t),
	}
	sure_cls_gen.Gen(cfg, Client{})
}
