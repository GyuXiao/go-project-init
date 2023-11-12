# Go-Project-Init

为助力 Go 服务应用的快速开发，本项目相当于一个模板，可以快速初始化一个 Go 项目代码框架。

## 快速开始

```bash
# 克隆到本地
git clone https://github.com/GyuXiao/go-project-init.git

# 进入项目
cd go-project-init

# 运行项目
go run main
```

## 代码工程结构

主要分为以下部分：
- `main.go`：入口文件，用于启动项目
- conf 目录：用于存放配置文件
- routers 目录：用于存放路由相关代码
- handlers 目录：用于存放每个 router 对应的 handler 代码，主要是参数校验、调用业务逻辑代码、返回参数
- service 目录：用于存放业务逻辑代码
- model 目录：用于存放数据库相关代码
- pkg 目录：用于存放公共代码，比如业务错误码、JWT、雪花算法、工具类等
- middleware 目录：用于存放中间件代码，比如日志、JWT 校验、context 信息，错误信息处理等
