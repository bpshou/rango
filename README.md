# rango

#### 介绍
gin框架服务

#### 启动
- 根目录执行 `go mod init <服务名>` 模块初始化
- `go mod tidy` 添加&整理模块
- `go mod vendor` 将模块整理到vendor目录
- `go run gin.go` 启动服务

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

#### 使用类库
1. gin          服务框架
2. logrus       日志库
3. viper        配置文件库
3. xorm         数据库
