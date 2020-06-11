FROM golang:1.14 as builder
WORKDIR $GOPATH/src/github.com/currycan/scanner
COPY ./ ./
RUN set -ex;CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make

FROM alpine:3.12
LABEL maintainer="currycan <ansandy@foxmail.com>"
COPY --from=builder $GOPATH/src/github.com/currycan/scanner/scanner /usr/local/bin
RUN apk add -U --no-cache ca-certificates \
    tzdata busybox-extras curl ca-certificates; \
    ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime; \
    echo ${TZ} > /etc/timezone; \
    touch ~/scanner.yaml; \
    rm -rf /var/cache/apk/*

ENV TZ=Asia/Shanghai

WORKDIR /etc/scanner

CMD ["scanner", "--config", "/etc/scanner/scanner.yaml"]
