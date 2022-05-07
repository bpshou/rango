# rango

#### 介绍
gin框架服务

#### 启动
- 根目录执行 `go mod init <服务名>` 模块初始化
- `go mod tidy` 添加&整理模块
- `go run gin.go` 启动服务

#### 目录结构
├─route     # 路由
├─service   # 服务
└─utils     # 工具类

#### 使用类库
1. gin          服务框架
2. logrus       日志库
3. viper        配置文件库
