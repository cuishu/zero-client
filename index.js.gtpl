{{ range .Syntax.Doc}}
{{.}}
{{end}}

{{ range .Syntax.Comment}}
{{.}}
{{end}}

const syntax = {{.Syntax.Version}}

{{range .Types}}
{{range .Documents}}{{.}}{{end}}
class {{.Name}} {
    {{range .Fields}}{{range .Documents}}{{.}}{{end}}
    {{.Name}}: {{.Type}}
    {{end}}
    constructor({{range .Fields}}{{.Name}}: {{.Type}},{{end}}) {
        {{range .Fields}}this.{{.Name}} = {{.Name}};
        {{end}}
    }
}
{{end}}

class {{.ApiName}} {
    host: string
    constructor(host: string) {
        this.host = host;
    }
    {{range .Route}}
    {{range .Doc}}{{.}}{{end}}
    function {{.FuncName}}(req {{.Request}}) : Promise<{{.Response}}, any> {
        return new Promise((reslove, reject)=>{
            fetch(this.host+'{{.Path}}')
        });
    }
    {{end}}
}