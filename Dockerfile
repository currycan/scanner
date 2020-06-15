FROM golang:1.14-alpine as builder
WORKDIR $GOPATH/src/github.com/currycan/
#ENV GOPROXY=https://goproxy.cn
RUN set -ex; apk add --no-cache upx ca-certificates tzdata;
COPY ./go.mod ./
COPY ./go.sum ./
RUN set -ex; go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o server &&\
  upx --best server -o _upx_server && \
  mv -f _upx_server server

FROM alpine:3.12
LABEL maintainer="currycan <ansandy@foxmail.com>"
COPY --from=builder /go/src/github.com/currycan/scanner/server /usr/bin/scanner
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN apk add -U --no-cache \
        busybox-extras curl;  \
    touch ~/scanner.yaml; \
    rm -rf /var/cache/apk/*

WORKDIR /etc/scanner

CMD ["scanner", "--config", "/etc/scanner/scanner.yaml"]
