{{.Documents}}
{{range .Types}}
{{.Documents}}
class {{.Name}} {
    {{range .Fields}}{{.Documents}}
    {{.Name}}: {{.Type}};
    {{end}}
    constructor({{range .Fields}}{{.Name}}: {{.Type}},{{end}}) {
        {{range .Fields}}this.{{.Name}} = {{.Name}};
        {{end}}
    }
}
{{end}}

{{.ServiceDoc}}
class {{.ApiName}} {
    host: string;
    http_request: any;
    upload: any;
    constructor(host: string, http_request: any, upload: any) {
        this.host = host;
        this.http_request = http_request;
        this.upload = upload;
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
            this.http_request(`${this.host}{{.Path}}`, {
                method: '{{.Method}}',
                data: data,
            }).then(data=>data.json()).then(data=>reslove(data)).catch(err=>reject(err));
        });
    }
    {{end}}
}