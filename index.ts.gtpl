{{.Documents}}
{{range .Types}}
{{.Documents}}
class {{.Name}} {
    {{range .Fields}}{{.Documents}}
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
    {{.Doc}}
    public {{.FuncName}}(req: {{.Request}}) : Promise<{{.Response}}> {
        return new Promise((reslove, reject)=>{
            fetch(`${this.host}{{.Path}}`, {
                method: '{{.Method}}'
            }).then(data=>data.json()).then(data=>reslove(data)).catch(err=>reject(err))
        });
    }
    {{end}}
}