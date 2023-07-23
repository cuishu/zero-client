"use strict";
{{.Comment}}
Object.defineProperty(exports, "__esModule", { value: true });
exports.DeleteResp = exports.DeleteReq = exports.AddResp = exports.AddReq = void 0;
{{range .Route}}exports.{{.FuncName}} = {{end}}void 0;

{{range .Types}}
{{.Documents}}
class {{.Name}} {
    constructor({{range .Fields}}{{.Name}},{{end}}) {
        {{range .Fields}}this.{{.Name}} = {{.Name}};
        {{end}}
    }
}
exports.{{.Name}} = {{.Name}};
{{end}}

{{.ServiceDoc}}
class {{.ApiName}} {
    constructor(conf) {
        this.host = conf.host;
        this.http_request = conf.http_request;
        this.upload = conf.upload;
    }
    {{range .Route}}
    {{.Comment}}
    {{.FuncName}}(req) {
        return new Promise((reslove, reject)=>{
            let url = `${this.host}{{.Path}}`;
            {{if eq .ReqType.IsJSON true}}
            let data = req;
            {{if eq .Method "GET"}}
            const query = Object.keys(data).map((x)=>(`${x}=${data[x]}`)).join('&');
            {{end}}
            this.http_request(url, {
                uri: '{{.Path}}',
                method: '{{.Method}}',
                {{if eq .Method "GET"}}query: query,{{end}}
                data: data,
            }).then((data)=>{
                if (data.fail) {
                    reject(data.msg);
                } else {
                    reslove(data.data);
                }
            }).catch((err)=>reject(err));
            {{else}}
            let data = new FormData();
            {{range .ReqType.Fields}}data.append('{{.Name}}', req.{{.Name}});
            {{end}}
            {{end}}
        });
    }
    {{end}}
}
exports.default = Example;
