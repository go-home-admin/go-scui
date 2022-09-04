GOBIN := $(shell go env GOBIN)
ATDIR := $(shell pwd)

# 安装代码工具(开发机器需要)
# export GOPATH=$HOME/go PATH=$PATH:$GOPATH/bin
mac-install:
	protoc --version || brew install protobuf						# mac下自动安装, win环境手动安装
	go get -u google.golang.org/protobuf/proto
	go install github.com/golang/protobuf/protoc-gen-go@latest		# proto 工具链, 生成go代码插件
	go install github.com/go-home-admin/toolset@latest

# Orm自动维护
make-orm:
	toolset make:orm

# 只维护 protoc
protoc:
	toolset  make:protoc

make-route:
	toolset  make:route

make-swagger:
	toolset  make:swagger -out_route=@root/web/swagger/swagger.json

make-bean:
	toolset  make:bean

# 生成全部
gen:protoc make-route make-bean make-swagger

# 调试启动
dev:gen
	go run main.go

