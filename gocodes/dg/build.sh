#!/usr/bin/env bash

echo "清空bin目录"
rm -rf bin/*

# 默认项目名为当前路径最后一段的文件夹名
PROJECT_PATH=`pwd`
PROJECT_NAME=${PROJECT_NAME:-`basename ${PROJECT_PATH}`}
PROJECT_BRANCH=${PROJECT_BRANCH:-`git symbolic-ref --short -q HEAD`}
PROJECT_VERSION=${PROJECT_VERSION:-`cat VERSION`}
PROJECT_REPOSITORY="dockerhub.deepglint.com/deepface/${PROJECT_NAME}-${PROJECT_BRANCH}"
echo "项目名称: ${PROJECT_NAME}"
echo "项目分支: ${PROJECT_BRANCH}"
echo "项目版本: ${PROJECT_VERSION}"
echo "仓库名称: ${PROJECT_REPOSITORY}"

# 编译
echo "编译中"
# 编译
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -v -o bin/${PROJECT_NAME} || (echo "编译错误" && exit 1) || exit 1
echo "编译完成"
# 打包
echo "打包中"
#cp releasenote bin/ || (echo "拷贝releasenote失败" && exit 1) || exit 1
cp sv_start.sh bin/ || (echo "拷贝sv_start.sh失败" && exit 1) || exit 1
echo "打包完成"
# 构筑镜像
echo "构筑开始"
docker build --build-arg path=bin --build-arg name=${PROJECT_NAME} --tag ${PROJECT_REPOSITORY}:${PROJECT_VERSION} . || (echo "构筑错误" && exit 1) || exit 1
echo "构筑完成"
# 推送镜像
echo "推送开始"
docker push ${PROJECT_REPOSITORY}:${PROJECT_VERSION} || (echo "推送错误" && exit 1) || exit 1
echo "推送完成"
