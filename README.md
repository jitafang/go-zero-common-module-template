# go-zero公共模块模板

#### 介绍
go-zero的公共模块和第三方库

#### go-zero goctl 使用
```
# 生成本地模板
goctl template init --home=./template
```

```
# 快速生成 Go HTTP 服务
goctl api new demo --style=go_zero
```

```
# 切到对应服务中使用相关api操作
# 递归格式化目录下的 api 文件 再根据 api 文件生成 Go HTTP 代码
goctl api format --dir=./api
# 校验 api 文件是否符合规范。
goctl api validate --api=./api/main.api
# 根据 api 文件生成 Go HTTP 代码
goctl api go --api=./api/main.api --home=../template --style=go_zero  --dir=.
```


#### orm数据库 ent. 使用
官网：https://entgo.io/
```
# 在对应服务中使用
# 初始化一个新的 Schema
go run entgo.io/ent/cmd/ent new User
# 修改generate.go
//go:generate go run entgo.io/ent/cmd/ent generate  --target=./model ./schema
# 生成资源文件
go generate ./ent
```