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

{{.ServiceDoc}}
class {{.ApiName}} {
    host: string
    constructor(host: string) {
        this.host = host;
    }
    {{range .Route}}
    {{.Doc}}
    public {{.FuncName}}(req: {{.Request}}) : Promise<{{.Response}}> {
        return new Promise((reslove, reject)=>{
            {{if eq .ReqType.IsJSON true}}let data = req;{{else}}
            let data = new FormData();
            {{range .ReqType.Fields}}data.append('{{.Name}}', req.{{.Name}});
            {{end}}
            {{end}}
            fetch(`${this.host}{{.Path}}`, {
                method: '{{.Method}}',
                data: data,
            }).then(data=>data.json()).then(data=>reslove(data)).catch(err=>reject(err));
        });
    }
    {{end}}
}