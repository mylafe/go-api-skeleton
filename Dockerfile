FROM golang:alpine AS builder

ENV GO111MODULE=on \
GOPROXY=https://goproxy.cn,direct \
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o goapiskeleton .

###################
# 接下来创建一个小镜像
###################
FROM debian:stretch-slim

COPY ./env /env

# 从builder镜像中把可执行文件拷贝到当前目录
COPY --from=builder /build/goapiskeleton /

# 需要运行的命令
ENTRYPOINT ["/goapiskeleton", ".env"]
