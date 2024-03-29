# 使用说明

{{.Documents}}

## 安装

```
npm install --registry="https://npm.wash-painting.cn" {{.ApiName}}
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

const config = {
  host: 'https://ftms.qingyuantop.top',
  http_request: fetchsomething,
  upload: fetchsomething,
}

const client = new {{.ApiName}}(config);
```

## 接口文档

{{range .Route}}
{{.Doc}}

**请求参数***
|名称|类型|校验规则|说明|
|:-:|:-:|:-:|:-:|
{{range .ReqType.Fields}}|{{.Name}}|{{.Type}}|{{.Validate}}|{{.Documents}}|
{{end}}
class {{.ReqType.Name}} {
    {{range .ReqType.Fields}}{{.Documents}}
    {{.Name}}: {{.Type}};
    {{end}}
}
**返回值**
|名称|类型|说明|
|:-:|:-:|:-:|
{{range .RespType.Fields}}|{{.Name}}|{{.Type}}|{{.Documents}}|
{{end}}

{{end}}
class {{.ResType.Name}} {
    {{range .ResType.Fields}}{{.Documents}}
    {{.Name}}: {{.Type}};
    {{end}}
}

```javascript
const req = {{.Request}}()
client.{{.FuncName}}(req).then(...).catch(...)
```
{{end}}