{{.Comment}}
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
    {{.Comment}}
    public {{.FuncName}}(req: {{.Request}}) : Promise<{{.Response}}> {
        return new Promise((reslove, reject)=>{
            let url = `${this.host}{{.Path}}`;
            {{if eq .ReqType.IsJSON true}}
            let data = req;
            {{if eq .Method "GET"}}
            const query = Object.entries(data).map(x=>x.reduce((a,b)=>(`${a}=${b}`))).reduce((a,b)=>(`${a}&${b}`),'');
            if (query != '') {
                url += '?' + query;
            }
            {{end}}
            this.http_request(url, {
                method: '{{.Method}}',
                data: data,
            }).then(data=>{
                if (data.fail) {
                    reject(data.msg);
                } else {
                    reslove(data.data);
                }
            }).catch(err=>reject(err));
            {{else}}
            let data = new FormData();
            {{range .ReqType.Fields}}data.append('{{.Name}}', req.{{.Name}});
            {{end}}
            {{end}}
        });
    }
    {{end}}
}