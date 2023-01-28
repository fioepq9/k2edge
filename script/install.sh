#!/bin/sh

# 设置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,direct

# goctl 安装
# https://go-zero.dev/cn/docs/goctl/goctl#%E5%AE%89%E8%A3%85-goctl
go install github.com/zeromicro/go-zero/tools/goctl@latest

# protoc & protoc-gen-go安装
# https://go-zero.dev/cn/docs/prepare/protoc-install#%E6%96%B9%E5%BC%8F%E4%B8%80goctl%E4%B8%80%E9%94%AE%E5%AE%89%E8%A3%85
goctl env check -i -f --verbose

# gum 安裝
# https://github.com/charmbracelet/gum#installation
#go install github.com/charmbracelet/gum@latest

# typer 安裝
# https://typer.tiangolo.com/#installation
pip install "typer[all]"

# gitmoji-cli 安裝
# https://github.com/carloscuesta/gitmoji#using-gitmoji-cli
npm i -g gitmoji-cli