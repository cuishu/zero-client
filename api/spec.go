package api

import (
	"fmt"
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
	Documents  string
	Comment    string
	Info       Info
	ApiName    string
	Types      []Type
	ServiceDoc string
	Route      []Route
}

var typeMap map[string]Type

func init() {
	typeMap = make(map[string]Type)
}

func toDocument(comment string, n int) string {
	space := strings.Repeat(" ", n)
	return fmt.Sprintf("/**\n%s* %s\n%s*/",
		space,
		strings.TrimSpace(strings.Trim(strings.Trim(comment, "/"), "*")),
		space)
}

func ToSpec(spec *ast.Spec) Spec {
	var ret Spec
	ret.Comment = spec.Comment
	ret.Documents = strings.TrimSpace(strings.Trim(strings.Trim(spec.Service.Comment, "/"), "*"))
	ret.ApiName = spec.Service.Name
	ret.Info = toInfo(spec.Info)
	for _, item := range spec.Types {
		t := convertSpecType(item)
		ret.Types = append(ret.Types, t)
		typeMap[t.Name] = t
	}
	ret.ServiceDoc = toDocument(ret.Comment, 1)
	for _, item := range spec.Service.Apis {
		ret.Route = append(ret.Route, Route{
			FuncName:   item.Handler,
			Request:    item.Input,
			ReqType:    typeMap[item.Input],
			Response:   item.Output,
			ResType:    typeMap[item.Output],
			Path:       item.URI,
			Comment:    toDocument(item.Comment, 4),
			Doc:        strings.TrimSpace(strings.Trim(strings.Trim(item.Comment, "/"), "*")),
			Method:     strings.ToUpper(item.Method),
			ValidToken: item.ValidToken,
		})
	}
	return ret
}
