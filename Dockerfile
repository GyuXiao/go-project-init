FROM golang:1.18-alpine AS builder

# 为镜像设置必要的环境变量
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


# Docker 的最佳实践之一是通过仅保留二进制文件来减小镜像大小，为此，我们将使用一种称为多阶段构建的技术，这意味着我们将通过多个步骤构建镜像。
# 使用这种技术，我们剥离了使用 golang:alpine 作为编译镜像来编译得到二进制可执行文件的过程，
# 并基于 scratch 生成一个简单的、非常小的新镜像。
# 我们将二进制文件从命名为 builder 的第一个镜像中复制到新创建的 scratch 镜像中。
# 创建一个小镜像
FROM scratch

# 从 builder 镜像中把配置文件拷贝到当前目录
COPY ./conf /conf

# 从 builder 镜像中把 /dist/app 拷贝到当前目录
COPY --from=builder /build/gyublog_app /

# 需要运行的命令
CMD ["/gyublog_app", "conf/config.yaml"]