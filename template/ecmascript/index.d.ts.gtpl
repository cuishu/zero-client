{{.Comment}}
{{range .Types}}
{{.Documents}}
export declare class {{.Name}} {
    [key: string]: any;
    {{range .Fields}}{{.Documents}}
    {{.Name}}: {{.Type}};
    {{end}}
    /**{{range .Fields}}
     * @param {{.Name}} {{.Type}} {{.Doc}}{{end}}
     */
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
    /**
     * {{.Doc}}
     *
     * @param req {{.Request}} {{.ReqType.Doc}}
     * @return {{.Response}}
     */
    public {{.FuncName}}(req: {{.Request}}) : Promise<{{.Response}}>;
    {{end}}
}
