FROM golang:1.24-alpine

# 设置 Go 代理
ENV GOPROXY=https://goproxy.cn

# 安装 air
RUN go install github.com/air-verse/air@latest

WORKDIR /app
COPY . .
CMD ["air"]