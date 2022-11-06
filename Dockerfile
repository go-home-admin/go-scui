FROM feizhaoer/gohome:latest AS builder
#FROM golang:1.18.4-alpine AS builder
#
ADD . /app
WORKDIR /app
#
#RUN apk update && apk upgrade
#RUN apk add protobuf-dev
#RUN go env -w GOPROXY=https://goproxy.cn,direct
#RUN go install github.com/golang/protobuf/protoc-linux-gen-go@v1.5.2
#RUN go install github.com/go-home-admin/toolset@latest

RUN toolset make:protoc-linux \
    toolset make:orm \
    toolset make:route \
    toolset make:bean \
    toolset make:swagger -out_route=@root/web/swagger/swagger.json

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go


FROM alpine:latest AS production
COPY --from=builder /app/main /main

## we can then kick off our newly compiled
## binary exectuable!!
CMD ["./main"]