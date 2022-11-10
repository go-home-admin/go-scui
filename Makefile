GOBIN := $(shell go env GOBIN)
ATDIR := $(shell pwd)

# 安装代码工具(开发机器需要)
# export GOPATH=$HOME/go PATH=$PATH:$GOPATH/bin
install:
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/go-home-admin/toolset@latest

# Orm自动维护
make-orm:
	toolset make:orm

# 只维护 protoc-linux
protoc:
	toolset make:protoc

make-route:
	toolset make:route

make-swagger:
	toolset make:swagger -out_route=@root/web/swagger/swagger.json

make-js:
	toolset make:js -in=@root/web/swagger/swagger.json -tag=admin

make-bean:
	toolset make:bean

# 生成全部
gen:bin/protoc-linux make-orm make-route make-bean make-swagger make-js

# 调试启动
dev:gen
	go run main.go

