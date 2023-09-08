#!/bin/bash

if [ "$1" = "nginx" ] || [ "$1" = "golang" ]; then
    # 启动nginx
    nginx

    cd rango/server

    go mod tidy
    go run main.go
fi

exec "$@"
