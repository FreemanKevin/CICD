#!/bin/bash

# 确保脚本抛出遇到的错误
set -e

# 获取脚本所在目录的绝对路径
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# 切换到Go目录
cd "${SCRIPT_DIR}"

# 清理旧的构建文件
echo "Cleaning..."
rm -rf bin/*

# 下载依赖并构建
echo "Building Go application..."
go mod download
go mod tidy
CGO_ENABLED=0 GOOS=linux go build -o bin/app ./cmd/main.go

# 构建镜像并运行
echo "Building and running docker container..."
docker-compose up -d --build 