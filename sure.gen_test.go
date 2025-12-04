package gogit

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/sure_cls_gen"
	"github.com/yyle88/syntaxgo"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

// TestGen generates Must support classes for Client using sure package
// Creates error-free API wrappers with automatic error handling patterns
// Generated code provides panic-based alternatives to error-returning methods
//
// TestGen 使用 sure 库为 Client 生成 Must 包装器类
// 创建带自动错误处理模式的无错误 API 包装器
// 生成的代码提供基于 panic 的替代方法来替代返回错误的方法
func TestGen(t *testing.T) {
	param := sure_cls_gen.NewClassGenOptions(runpath.PARENT.Path())
	param.WithNewClassNameParts("88")
	param.WithNamingPatternType(sure_cls_gen.STYLE_SUFFIX_CAMELCASE_TYPE)
	param.MoreErrorHandlingModes(sure.MUST)

	cfg := &sure_cls_gen.ClassGenConfig{
		ClassGenOptions: param,
		PackageName:     syntaxgo.CurrentPackageName(),
		ImportOptions:   syntaxgo_ast.NewPackageImportOptions().SetInferredObject(git.Status{}).SetInferredObject(object.Commit{}),
		OutputPath:      runtestpath.SrcPath(t),
	}
	sure_cls_gen.GenerateClasses(cfg, Client{})
}
