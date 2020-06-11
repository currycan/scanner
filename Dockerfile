FROM golang:1.14 as builder
WORKDIR /go/src/github.com/currycan/scanner
RUN set -ex; apk update; apk add --no-cache gcc
COPY ./ ./
RUN go build -o scanner

FROM alpine:latest
LABEL maintainer="currycan <ansandy@foxmail.com>"
COPY --from=builder /go/src/github.com/currycan/scanner/scanner /usr/local/bin
RUN apk add -U --no-cache ca-certificates \
    tzdata busybox-extras curl ca-certificates; \
    ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime; \
    echo ${TZ} > /etc/timezone; \
    touch ~/scanner.yaml; \
    rm -rf /var/cache/apk/*

ENV TZ=Asia/Shanghai

WORKDIR /etc/scanner

CMD ["scanner", "--config", "/etc/scanner/scanner.yaml"]
