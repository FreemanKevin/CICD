#!/bin/bash

# 确保脚本抛出遇到的错误
set -e

# 获取脚本所在目录的绝对路径
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# 切换到JS目录
cd "${SCRIPT_DIR}"

# 清理环境
rm -rf dist
rm -rf node_modules

# 安装依赖 (使用淘宝镜像)
echo "Installing dependencies..."
npm install --registry=https://registry.npmmirror.com

# 构建前端资源
echo "Building frontend assets..."
npm run build

# 构建镜像并运行
echo "Building and running docker container..."
docker-compose up -d --build