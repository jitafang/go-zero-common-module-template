#### orm数据库 ent. 使用
官网：https://entgo.io/

将数据库操作单抽出来，根据数据库来分模块；
entgo只作为orm进行数据库操作，生成的数据库结构比较死板

```
# 在模块下初始化一个新的 Schema
go run entgo.io/ent/cmd/ent new User
# 修改generate.go
//go:generate go run entgo.io/ent/cmd/ent generate  --target=./model ./schema
# 生成资源文件
go generate ./ent
```