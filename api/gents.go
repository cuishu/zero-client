package api

import (
	_ "embed"
	"io"
	"os"
	"path"
	"strings"
	"text/template"
)

type CodeGenerator struct {
	TsIndexTemplate   string
	TsPackageTemplate string
	TsconfigTemplate  string
	TsReadmeTemplate  string

	JsIndexTsTemplate string
	JsIndexTemplate   string
	JsPackageTemplate string
	JsReadMeTemplate  string

	Outpath string
	ApiSpec Spec
}

func (c *CodeGenerator) parseIndexTemplate(apiSpec Spec) {
	indexTmpl, err := template.New("index.ts").Funcs(template.FuncMap{
		"join": strings.Join,
	}).Parse(c.TsIndexTemplate)
	if err != nil {
		panic(err)
	}

	var w io.Writer = os.Stdout

	if c.Outpath != "" {
		var file *os.File
		if file, err = os.OpenFile(path.Join(c.Outpath, "index.ts"),
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

func (c *CodeGenerator) parsePackageTemplate(apiSpec Spec) {
	pkgTmpl, err := template.New("package.json").Parse(c.TsPackageTemplate)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout

	if c.Outpath != "" {
		var file *os.File
		if file, err = os.OpenFile(path.Join(c.Outpath, "package.json"),
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
			panic(err)
		}
		defer file.Close()

		w = file
	}
	pkgTmpl.Execute(w, apiSpec)
}

func (c *CodeGenerator) parseTsconfigTemplate(apiSpec Spec) {
	tscTmpl, err := template.New("tsconfig.json").Parse(c.TsconfigTemplate)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout

	if c.Outpath != "" {
		var file *os.File
		if file, err = os.OpenFile(path.Join(c.Outpath, "tsconfig.json"),
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
			panic(err)
		}
		defer file.Close()

		w = file
	}
	tscTmpl.Execute(w, apiSpec)
}

func (c *CodeGenerator) parseReadmeTemplate(apiSpec Spec) {
	tscTmpl, err := template.New("README.md").Parse(c.TsReadmeTemplate)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout

	if c.Outpath != "" {
		var file *os.File
		if file, err = os.OpenFile(path.Join(c.Outpath, "README.md"),
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
			panic(err)
		}
		defer file.Close()

		w = file
	}
	tscTmpl.Execute(w, apiSpec)
}

func (c *CodeGenerator) GenTsCode() {
	c.parseIndexTemplate(c.ApiSpec)
	c.parsePackageTemplate(c.ApiSpec)
	c.parseTsconfigTemplate(c.ApiSpec)
	c.parseReadmeTemplate(c.ApiSpec)
}
