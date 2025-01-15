#!/bin/bash

# 确保脚本抛出遇到的错误
set -e

# 获取脚本所在目录的绝对路径
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# 切换到Java目录
cd "${SCRIPT_DIR}"

# 编译打包
echo "Building Java application..."
mvn clean install -U -DskipTests
mvn clean package -U -DskipTests

# 构建镜像并运行
echo "Building and running docker container..."
docker-compose up -d --build 