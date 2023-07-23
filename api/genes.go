package api

import (
	_ "embed"
	"io"
	"os"
	"path"
	"strings"
	"text/template"
)

func (c *CodeGenerator) jsParseIndexDTsTemplate(apiSpec Spec) {
	indexTmpl, err := template.New("index.d.ts").Funcs(template.FuncMap{
		"join": strings.Join,
	}).Parse(c.JsIndexTsTemplate)
	if err != nil {
		panic(err)
	}

	var w io.Writer = os.Stdout

	if c.Outpath != "" {
		var file *os.File
		if file, err = os.OpenFile(path.Join(c.Outpath, "index.d.ts"),
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

func (c *CodeGenerator) jsParsePackageTemplate(apiSpec Spec) {
	pkgTmpl, err := template.New("package.json").Parse(c.JsPackageTemplate)
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

func (c *CodeGenerator) jsParseIndexTemplate(apiSpec Spec) {
	tscTmpl, err := template.New("index.js").Parse(c.JsIndexTemplate)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout

	if c.Outpath != "" {
		var file *os.File
		if file, err = os.OpenFile(path.Join(c.Outpath, "index.js"),
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
			panic(err)
		}
		defer file.Close()

		w = file
	}
	tscTmpl.Execute(w, apiSpec)
}

func (c *CodeGenerator) jsParseReadmeTemplate(apiSpec Spec) {
	tscTmpl, err := template.New("README.md").Parse(c.JsIndexTemplate)
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

func (c *CodeGenerator) GenEsCode() {
	c.jsParseIndexDTsTemplate(c.ApiSpec)
	c.jsParsePackageTemplate(c.ApiSpec)
	c.jsParseIndexTemplate(c.ApiSpec)
	c.jsParseReadmeTemplate(c.ApiSpec)
}
