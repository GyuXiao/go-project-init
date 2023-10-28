FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .


# 将我们的代码编译成二进制可执行文件app
RUN go build -o gyublog_app .

#
## 移动到用于存放生成的二进制文件的 /dist 目录
#WORKDIR /dist
#
## 将二进制文件从 /build 目录复制到这里
#RUN cp /build/app .
#
## 声明服务端口
#EXPOSE 8081
#
## 待续
#
## 启动容器时运行的命令
#CMD ["/dist/app"]

# 创建一个小镜像
FROM scratch

# 从 builder 镜像中把配置文件拷贝到当前目录
COPY ./conf /conf

# 从 builder 镜像中把 /dist/app 拷贝到当前目录
COPY --from=builder /build/gyublog_app /

#RUN set -eux; \
#	apt-get update; \
#	apt-get install -y \
#		--no-install-recommends \
#		netcat; \
#        chmod 755 wait-for.sh

# 需要运行的命令
CMD ["/gyublog_app", "conf/config.yaml"]