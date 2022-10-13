package main

import (
	_ "embed"
	"flag"
	"io"
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

	var w io.Writer = os.Stdout

	if outpath != "" {
		var file *os.File
		if file, err = os.OpenFile(path.Join(outpath, "index.ts"),
			os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			panic(err)
		}
		defer file.Close()

		w = file
	}
	indexTmpl.Execute(w, apiSpec)
}

func parsePackageTemplate(apiSpec api.Spec) {
	pkgTmpl, err := template.New("package.json").Parse(packageTemplate)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout

	if outpath != "" {
		var file *os.File
		if file, err = os.OpenFile(path.Join(outpath, "package.json"),
			os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			panic(err)
		}
		defer file.Close()

		w = file
	}
	pkgTmpl.Execute(w, apiSpec)
}

func parseTsconfigTemplate(apiSpec api.Spec) {
	tscTmpl, err := template.New("tsconfig.json").Parse(tsconfigTemplate)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout

	if outpath != "" {
		var file *os.File
		if file, err = os.OpenFile(path.Join(outpath, "tsconfig.json"),
			os.O_CREATE|os.O_WRONLY, 0644); err != nil {
			panic(err)
		}
		defer file.Close()

		w = file
	}
	tscTmpl.Execute(w, apiSpec)
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
