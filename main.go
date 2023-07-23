package main

import (
	_ "embed"
	"flag"
	"os"
	"os/exec"
	"text/template"

	"github.com/cuishu/zero-api/ast"
	"github.com/cuishu/zero-client/api"
)

var (
	//go:embed template/typescript/index.ts.gtpl
	indexTemplate string
	//go:embed template/typescript/package.json.gtpl
	packageTemplate string
	//go:embed template/typescript/tsconfig.json.gtpl
	tsconfigTemplate string
	//go:embed template/typescript/README.md.gtpl
	readmeTemplate string

	//go:embed template/ecmascript/index.d.ts.gtpl
	jsIndexTsTemplate string
	//go:embed template/ecmascript/index.js.gtpl
	jsIndexTemplate string
	//go:embed template/ecmascript/package.json.gtpl
	jsPackageTemplate string
	//go:embed template/ecmascript/README.md.gtpl
	jsReadMeTemplate string
)

var (
	filename   string
	outpath    string
	typescript bool
)

func init() {
	flag.StringVar(&filename, "f", "", "api filename")
	flag.StringVar(&outpath, "o", "", "output path")
	flag.BoolVar(&typescript, "ts", false, "generate typescript")
	flag.Parse()

	if outpath != "" {
		if err := os.MkdirAll(outpath, 0700); err != nil {
			panic(err)
		}
	}
}

func genFileOverwrite(filename, tmpl string, spec any) {
	t, err := template.New(filename).Parse(tmpl)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	if err := t.Execute(file, spec); err != nil {
		panic(err)
	}
}

func compileTypescript() {
	os.Chdir(outpath)
	cmd := exec.Command("npm", "run", "build")
	cmd.Run()
	os.Remove("index.ts")
	os.Remove("tsconfig.json")

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
	cg := api.CodeGenerator{
		TsIndexTemplate:   indexTemplate,
		TsPackageTemplate: packageTemplate,
		TsconfigTemplate:  tsconfigTemplate,
		TsReadmeTemplate:  readmeTemplate,

		JsIndexTsTemplate: jsIndexTsTemplate,
		JsIndexTemplate:   jsIndexTemplate,
		JsPackageTemplate: jsPackageTemplate,
		JsReadMeTemplate:  jsReadMeTemplate,

		Outpath: outpath,
		ApiSpec: apiSpec,
	}

	if typescript {
		cg.GenTsCode()
	} else {
		cg.GenEsCode()
	}
}
