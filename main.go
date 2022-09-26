package main

import (
	"bytes"
	_ "embed"
	"flag"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/cuishu/zero-client/api"
	parser "github.com/zeromicro/go-zero/tools/goctl/api/parser"
)

var (
	//go:embed index.ts.gtpl
	indexTemplate string
	//go:embed package.json.gtpl
	packageTemplate string
	//go:embed tsconfig.json.gtpl
	tsconfigTemplate string
)

var (
	filename string
	outpath  string
)

func init() {
	flag.StringVar(&filename, "f", "", "api filename")
	flag.StringVar(&outpath, "o", "", "output path")
	flag.Parse()

	if outpath != "" {
		if err := os.MkdirAll(outpath, 0700); err != nil {
			panic(err)
		}
	}
}

func parseIndexTemplate(apiSpec api.Spec) {
	indexTmpl, err := template.New("index.ts").Funcs(template.FuncMap{
		"join": strings.Join,
	}).Parse(indexTemplate)
	if err != nil {
		panic(err)
	}
	var buffer bytes.Buffer
	indexTmpl.Execute(&buffer, apiSpec)
	err = ioutil.WriteFile(path.Join(outpath, "index.ts"), buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func parsePackageTemplate(apiSpec api.Spec) {
	pkgTmpl, err := template.New("package.json").Parse(packageTemplate)
	if err != nil {
		panic(err)
	}
	var buffer bytes.Buffer
	pkgTmpl.Execute(&buffer, apiSpec)
	ioutil.WriteFile(path.Join(outpath, "package.json"), buffer.Bytes(), 0644)
}

func parseTsconfigTemplate(apiSpec api.Spec) {
	tscTmpl, err := template.New("tsconfig.json").Parse(tsconfigTemplate)
	if err != nil {
		panic(err)
	}
	var buffer bytes.Buffer
	tscTmpl.Execute(&buffer, apiSpec)
	ioutil.WriteFile(path.Join(outpath, "tsconfig.json"), buffer.Bytes(), 0644)
}

func main() {
	spec, err := parser.Parse(filename)
	if err != nil {
		panic(err)
	}
	if err := spec.Validate(); err != nil {
		panic(err)
	}
	apiSpec := api.ToSpec(spec)
	parseIndexTemplate(apiSpec)
	parsePackageTemplate(apiSpec)
	parseTsconfigTemplate(apiSpec)
}
