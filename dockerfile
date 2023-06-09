FROM golang:1.19 AS builder
 
WORKDIR /build

COPY . .

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn\ 
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -o main .
ENTRYPOINT ["/workdir/main"]