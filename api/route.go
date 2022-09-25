package api

import "github.com/zeromicro/go-zero/tools/goctl/api/spec"

type Route struct {
	FuncName string
	Request string
	Response string
	Path string
	Doc spec.Doc
}
