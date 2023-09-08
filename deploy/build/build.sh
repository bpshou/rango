#!/bin/bash

# 启动依赖
cd $(dirname $0)
cd ../../../

if [ ! -d "compose" ]; then
    git clone https://gitee.com/decezz/compose.git
fi

docker-compose -f compose/mysql/docker-compose.yml up -d mysql
docker-compose -f compose/redis/docker-compose.yml up -d


# 启动服务
cd -
cd ../docker-compose

# 重建
docker-compose build --no-cache --force-rm
docker-compose up --force-recreate -d

# 清理
docker system prune -f
