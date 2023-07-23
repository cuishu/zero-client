# 使用说明

{{.Documents}}

## 安装

```
npm install --registry="https://npm.cuishu.site" {{.ApiName}}
```

## 用法

用户需实现普通 http 请求函数和文件上传函数，返回值均为 json 对象

**示例**
```javascript
// reslove 和 reject 必须返回 json 对象
const http_request_function = (url, method, headers, body) => {
    return new Promise((reslove, reject)=>{});
}

const upload_function = (url, method, headers, data) => {
    return new Promise((reslove, reject)=>{});
}

const client = new {{.ApiName}}(host, http_request_function, upload_function);
```

## 接口文档

{{range .Route}}
{{.Doc}}
```javascript
// 请求参数
class {{.ReqType.Name}} {
    {{range .ReqType.Fields}}{{.Documents}}
    {{.Name}}: {{.Type}};
    {{end}}
}
// 返回值
class {{.ResType.Name}} {
    {{range .ResType.Fields}}{{.Documents}}
    {{.Name}}: {{.Type}};
    {{end}}
}
const req = {{.Request}}()
client.{{.FuncName}}(req).then(...).catch(...)
```
{{end}}