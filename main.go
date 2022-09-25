package main

import (
	_ "embed"
	"flag"
	"os"
	"strings"
	"text/template"

	"github.com/cuishu/zero-client/api"
	parser "github.com/zeromicro/go-zero/tools/goctl/api/parser"
)

//go:embed index.js.gtpl
var apiTemplate string

var filename string

func init() {
	flag.StringVar(&filename, "f", "", "api filename")
	flag.Parse()
}

func main() {
	spec, err := parser.Parse(filename)
	if err != nil {
		panic(err)
	}
	if err := spec.Validate(); err != nil {
		panic(err)
	}
	tmpl, err := template.New("index.js").Funcs(template.FuncMap{
		"join": strings.Join,
	}).Parse(apiTemplate)
	if err != nil {
		panic(err)
	}
	jsSpec := api.ToSpec(spec)
	tmpl.Execute(os.Stdout, jsSpec)
}
