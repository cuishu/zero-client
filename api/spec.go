package api

import (
	"strings"

	"github.com/cuishu/zero-api/ast"
)

type Info struct {
	Title   string
	Desc    string
	Author  string
	Email   string
	Version string
}

func toInfo(astInfo ast.Info) Info {
	var info Info
	info.Author = astInfo.Author
	info.Email = astInfo.Email
	info.Version = astInfo.Version
	if info.Version == "" {
		panic("Info 中缺少版本号")
	}
	return info
}

type Spec struct {
	Documents string
	Info      Info
	ApiName   string
	Types     []Type
	Route     []Route
}

func ToSpec(spec *ast.Spec) Spec {
	var ret Spec
	ret.Documents = spec.Comment
	ret.ApiName = spec.Service.Name
	ret.Info = toInfo(spec.Info)
	for _, item := range spec.Types {
		ret.Types = append(ret.Types, convertSpecType(item))
	}
	for _, item := range spec.Service.Apis {
		ret.Route = append(ret.Route, Route{
			FuncName: item.Handler,
			Request:  item.Input,
			Response: item.Output,
			Path:     item.URI,
			Doc:      item.Comment,
			Method:   strings.ToUpper(item.Method),
		})
	}
	return ret
}
