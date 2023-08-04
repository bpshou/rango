# rango

#### 介绍
gin框架服务

#### 启动
- 根目录执行 `go mod init <服务名>` 模块初始化
- `go mod tidy` 添加&整理模块
- `go mod vendor` 将模块整理到vendor目录
- `go run gin.go` 启动服务

#### 使用类库
1. gin          服务框架
2. logrus       日志库
3. viper        配置文件库
4. xorm         数据库
5. goqu         数据库SQL组装库
6. jwt          jwt认证库
7. otp          otp认证算法库
8. fasthttp     http请求库
9. xorm         数据库

#### 单元测试
- test是单元测试目录，官方规定所有测试文件后缀都必须以`_test.go`结尾
- 使用`官方test工具 + assert包`作为单元测试工具
- 单元测试文件函数名开头只能包含`Test、Benchmark、Example`，分别对应 `测试、基准、示例函数`
- 所有函数必须引入`testing`包

#### 目录结构
├─app               # 应用
│  ├─config         # 配置
│  │  ├─dev
│  │  └─test
│  ├─controller     # 控制器
│  │  ├─api
│  │  └─use      
│  └─service        # 服务
├─log
├─middleware        # 中间件
├─models            # model
│  └─golang      
├─router            # 路由
├─tools             # 工具类
│  ├─cipher
│  ├─logger
│  ├─redis
│  └─viper
└─vendor            # 三方类库

#### api文档
* 文档使用 swag 自动生成 swagger 文档，执行 `swag init` 即可将注释处理成文档
* docker启动 swagger `docker run -d -p 8080:8080 --name swagger swaggerapi/swagger-ui`
* 详细文档参见 https://github.com/swaggo/swag/blob/v1.16.1/README_zh-CN.md
