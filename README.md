# openapi-play/k8s

这个项目是 一个用于尝试 OpenAPI + golang + k8s 开发的 玩具仓库，

main.go 是入口文件会开启一个http server 提供一个 k8s 的 资源读接口

api/openapi.yaml 是 openapi 的定义文件，使用 oapi-codegen 来生成代码

api/js-client 是 生成的 js 客户端代码，用于测试

api/test.js 是 测试文件，用于测试生成的 js 客户端代码

api/schema 来自于 k8s 的 openapi-spec 的 json 文件，用于生成 k8s 的资源类型 的 前端代码
