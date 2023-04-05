# zero-client

## INSTALL
```bash
go install github.com/cuishu/zero-client@latest
```

## 使用方法

生成 typescript 客户端代码
```bash
# 将 typescript 代码打印到标准输出
zero-client -f example.api
```

```bash
# 输出到 client 目录
zero-client -f example.api -o client
```