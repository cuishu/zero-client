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
    public {{.FuncName}}(req: {{.Request}}) : Promise<{{.Response}}> {
        return new Promise((reslove, reject)=>{
            fetch(`${this.host}{{.Path}}`, {
                method: '{{.Method}}'
            }).then(data=>data.json()).then(data=>reslove(data)).catch(err=>reject(err))
        });
    }
    {{end}}
}