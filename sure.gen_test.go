package gogit

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
	param := sure_cls_gen.NewClassGenOptions(runpath.PARENT.Path())
	param.WithNewClassNameParts("88")
	param.WithNamingPatternType(sure_cls_gen.STYLE_SUFFIX_CAMELCASE_TYPE)
	param.MoreErrorHandlingModes(sure.MUST)

	cfg := &sure_cls_gen.ClassGenConfig{
		ClassGenOptions: param,
		PackageName:     syntaxgo.CurrentPackageName(),
		ImportOptions:   syntaxgo_ast.NewPackageImportOptions().SetInferredObject(git.Status{}),
		OutputPath:      runtestpath.SrcPath(t),
	}
	sure_cls_gen.GenerateClasses(cfg, Client{})
}
