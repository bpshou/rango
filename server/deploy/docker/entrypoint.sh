#!/bin/bash

if [ "$1" = "nginx" ] || [ "$1" = "golang" ]; then
    # 启动nginx
    nginx

    if [ ! -d "rango" ]; then
        git clone https://gitee.com/decezz/rango.git
    fi

    cd rango

    go mod tidy
    go run main.go
fi

exec "$@"
