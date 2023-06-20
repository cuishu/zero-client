package main

import (
	_ "embed"
	"flag"
	"io"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/cuishu/zero-api/ast"
	"github.com/cuishu/zero-client/api"
)

var (
	//go:embed index.ts.gtpl
	indexTemplate string
	//go:embed package.json.gtpl
	packageTemplate string
	//go:embed tsconfig.json.gtpl
	tsconfigTemplate string
	//go:embed README.md.gtpl
	readmeTemplate string
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
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
			panic(err)
		}
		defer file.Close()

		w = file
	}
	if err := indexTmpl.Execute(w, apiSpec); err != nil {
		panic(err)
	}
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
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
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
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
			panic(err)
		}
		defer file.Close()

		w = file
	}
	tscTmpl.Execute(w, apiSpec)
}

func parseReadmeTemplate(apiSpec api.Spec) {
	tscTmpl, err := template.New("README.md").Parse(readmeTemplate)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout

	if outpath != "" {
		var file *os.File
		if file, err = os.OpenFile(path.Join(outpath, "README.md"),
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
			panic(err)
		}
		defer file.Close()

		w = file
	}
	tscTmpl.Execute(w, apiSpec)
}

func main() {
	spec := ast.Parse(filename)
	if spec == nil {
		return
	}
	if err := spec.Validate(); err != nil {
		panic(err)
	}
	apiSpec := api.ToSpec(spec)
	parseIndexTemplate(apiSpec)
	parsePackageTemplate(apiSpec)
	parseTsconfigTemplate(apiSpec)
	parseReadmeTemplate(apiSpec)
}
