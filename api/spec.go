package api

import (
	spec "github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

type Spec struct {
	Syntax  spec.ApiSyntax
	ApiName string
	Types   []Type
	Route   []Route
}

func ToSpec(spec *spec.ApiSpec) Spec {
	var ret Spec
	ret.Syntax = spec.Syntax
	ret.ApiName = spec.Service.Name
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
		})
	}
	return ret
}
