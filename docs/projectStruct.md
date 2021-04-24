```
ffly-plus
├── config ## 配置
├── controller ## API实现,用来读取输入、调用业务处理、返回结果
│   └── api
│       └── v1
├── docs ## swag 文档
├── internal ##内部逻辑，业务目录
│   ├── cache　## 缓存
│   ├── code　## 错误码设计
│   ├── config　## 配置
│   ├── proto　## grpc proto
│   ├── sentinelm ## sentinel 限流
│   └── version　## 版本
├── models　## 数据库交互
├── pkg　## 一些封装好的 package
│   ├── token
│   └── utils
├── router ## 路由及中间件目录
│   ├── api
│   └── middleware
├── rpc ## 业务逻辑层
├── service ## 业务逻辑层
└── tool ## 小工具
    main.go # 项目入口文件
```
参考`https://github.com/colinrs/ffly-plus`