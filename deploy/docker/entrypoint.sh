#!/bin/bash

if [ "$1" = "nginx" ] || [ "$1" = "golang" ]; then
    # 启动nginx
    nginx

    git clone https://gitee.com/decezz/rango.git
    cd rango

    go mod tidy
    go run main.go
fi

exec "$@"
