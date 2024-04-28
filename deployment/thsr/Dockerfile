FROM golang:1.21 as builder

# 設置工作目錄
WORKDIR /app

# 將Go模組文件複製到容器中
COPY go.mod .
COPY go.sum .

# 下載所有依賴項
RUN go mod download

# 將源代碼複製到容器中
COPY . .

# 構建應用
RUN go build -tags netgo -ldflags '-s -w' -o app ./cmd/linebot

FROM alpine:latest AS release

RUN --mount=type=secret,id=accessid \
    echo ACCESS_ID=$(cat /run/secrets/accessid) >> ./.env

RUN --mount=type=secret,id=accesskey \
    echo ACCESS_KEY=$(cat /run/secrets/accesskey) >> ./.env

COPY --from=builder /app/app ./server
ENTRYPOINT ["./server"]
