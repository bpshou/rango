const path = require('path') // node.js的核心，专门用来处理路径的
const HtmlWebpackPlugin = require('html-webpack-plugin') // 导入html插件

module.exports = {
    // 入口
    entry: './src/main.js',
    // 输出
    output: {
        // 文件路径
        path: path.resolve(__dirname, 'dist'),
        // 文件名称
        filename: 'main.js',
        // 自动清空上次打包内容
        clean: true,
    },
    // 加载器
    module: {
        rules: [
            // loader配置
            {
                test: /\.css$/, // 只检测.css文件
                use: [
                    // 执行顺序 从右到左 从上到下
                    'style-loader', // 将js中css通过创建style标签添加html文件中生效
                    'css-loader', // 将css资源编译成common.js的模块到js中
                ],
            },
        ]
    },
    // 插件
    plugins: [
        // 加载html处理插件
        new HtmlWebpackPlugin({
            template: path.resolve(__dirname, 'public/index.html'),
        }),
    ],
    // 开发服务器
    devServer: {
        // 域名
        host: "localhost",
        // 端口
        port: 8081,
        // 是否自动打开浏览器
        open: true,
    },
    // 模式
    mode: 'development',
}
