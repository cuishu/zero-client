"use strict";
{{.Comment}}
Object.defineProperty(exports, "__esModule", { value: true });
exports.DeleteResp = exports.DeleteReq = exports.AddResp = exports.AddReq = void 0;
{{range .Route}}exports.{{.FuncName}} = {{end}}void 0;

{{range .Types}}
{{.Documents}}
class {{.Name}} {
    /**{{range .Fields}}
     * @param {{.Name}}: {{.Type}} {{.Doc}}{{end}}
     */
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
        this.token = conf.token;
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
                method: '{{.Method}}',{{if eq .Method "GET"}}
                query: query,{{end}}
                data: data,
                headers: { {{if ne .Method "GET"}}
                    'Content-Type': 'application/json',{{end}}{{if .ValidToken}}
                    'X-Token': this.token,{{end}}
                },
            }).then((data)=>{
                if (data.fail) {
                    reject(data.msg);
                } else {
                    reslove(data.data);
                }
            }).catch((err)=>reject(err));
            {{else}}
            let data = req;
            this.upload(url, {
                uri: '{{.Path}}',
                method: '{{.Method}}',{{if eq .Method "GET"}}
                query: query,{{end}}
                data: data,
                headers: {
                    'X-Token': this.token,
                },
            }).then((data)=>{
                if (typeof data === "object") {
                    if (data.fail) {
                        reject(data.msg);
                    } else {
                        reslove(data.data);
                    }
                } else {
                    reslove(data.data);
                }
            }).catch((err)=>reject(err));
            {{end}}
        });
    }
    {{end}}
}
exports.default = {{.ApiName}};
