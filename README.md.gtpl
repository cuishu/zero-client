# 使用说明

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