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

if [ "$1" = "online" ]; then
    # 重建
    docker-compose build --no-cache --force-rm
    docker-compose up --force-recreate -d
else
    docker-compose -f ../../server/deploy/docker-compose/docker-compose.yml up --force-recreate -d
fi

docker network connect mysql rango
docker network connect redis rango

# 清理
docker system prune -f
