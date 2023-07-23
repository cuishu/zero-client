{{.Comment}}
{{range .Types}}
{{.Documents}}
export declare class {{.Name}} {
    [key: string]: any;
    {{range .Fields}}{{.Documents}}
    {{.Name}}: {{.Type}};
    {{end}}
    constructor({{range .Fields}}{{.Name}}: {{.Type}},{{end}});
}
{{end}}

{{.ServiceDoc}}
export default class {{.ApiName}} {
    host: string;
    http_request: any;
    upload: any;
    constructor(conf: any);
    {{range .Route}}
    {{.Comment}}
    public {{.FuncName}}(req: {{.Request}}) : Promise<{{.Response}}>;
    {{end}}
}
