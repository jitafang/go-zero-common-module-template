

```
# 递归格式化目录下的 api 文件 再根据 api 文件生成 Go HTTP 代码
goctl api format --dir=./api

goctl api validate --api=./api/main.api
goctl api go --api=./api/main.api --home=../template --style=go_zero  --dir=.
```