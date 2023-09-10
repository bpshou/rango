## 建立项目
1. npm init -y
2. npm install webpack webpack-cli -D
3. npx webpack --mode development

## 配置webpack
### 建立webpack配置
1. 创建webpack.config.js
2. 写入代码
```
const path = require('path') // node.js的核心，专门用来处理路径的

module.exports = {
    // 入口
    entry: './src/main.js',
    // 输出
    output: {
        // 文件路径
        path: path.resolve(__dirname, 'dist'),
        // 文件名称
        filename: 'main.js'
    },
    // 加载器
    module: {
        rules: [
        ]
    },
    // 插件
    plugins: [
    ],
    // 模式
    mode: 'development',
}

```
3. 执行 `npx webpack`

### 加载css资源
1. 安装 `npm i style-loader css-loader -D`

### 加载html资源
1. 安装 `npm i html-webpack-plugin -D`

### 开发服务器自动化
1. 安装 `npm i webpack-dev-server -D`


