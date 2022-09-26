package api

import (
	"strings"

	spec "github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

type Info struct {
	Title   string
	Desc    string
	Author  string
	Email   string
	Version string
}

func toInfo(props map[string]string) Info {
	var info Info
	info.Title = strings.Trim(props["title"], `"`)
	info.Author = strings.Trim(props["author"], `"`)
	info.Desc = strings.Trim(props["desc"], `"`)
	info.Email = strings.Trim(props["email"], `"`)
	info.Version = strings.Trim(props["version"], `"`)
	if info.Version == "" {
		panic("Info 中缺少版本号")
	}
	return info
}

type Spec struct {
	Syntax  spec.ApiSyntax
	Info    Info
	ApiName string
	Types   []Type
	Route   []Route
}

func ToSpec(spec *spec.ApiSpec) Spec {
	var ret Spec
	ret.Syntax = spec.Syntax
	ret.ApiName = spec.Service.Name
	ret.Info = toInfo(spec.Info.Properties)
	for _, item := range spec.Types {
		ret.Types = append(ret.Types, convertSpecType(item))
	}
	for _, item := range spec.Service.Routes() {
		ret.Route = append(ret.Route, Route{
			FuncName: item.Handler,
			Request:  item.RequestTypeName(),
			Response: item.ResponseTypeName(),
			Path:     item.Path,
			Doc:      item.HandlerDoc,
			Method:   strings.ToUpper(item.Method),
		})
	}
	return ret
}
