FROM golang:1.22.4-bookworm
LABEL authors="zen"
# 更换国内源
COPY debian.sources /etc/apt/sources.list.d/
# 更新软件
RUN apt update
RUN apt install -y ffmpeg

# 复制go程序
RUN mkdir /app
WORKDIR /app
COPY . .
# 配置env
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/go/bin
RUN go mod vendor
# 启动程序
ENTRYPOINT ["go", "run","/app/main.go"]