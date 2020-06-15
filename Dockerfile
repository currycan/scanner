FROM golang:1.14-alpine as builder

WORKDIR $GOPATH/src/github.com/currycan/

#ENV GOPROXY=https://goproxy.cn

RUN set -ex; apk add --no-cache upx ca-certificates tzdata;
COPY ./go.mod ./
COPY ./go.sum ./
RUN set -ex; go mod download
COPY . .
RUN set -ex; \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o scanner; \
  upx --best scanner -o _upx_scanner; \
  mv -f _upx_scanner scanner

FROM alpine:3.12

LABEL maintainer="currycan <ansandy@foxmail.com>"

WORKDIR /etc/scanner

COPY --from=builder /go/src/github.com/currycan/scanner/scanner /usr/bin/scanner
COPY --from=builder /go/src/github.com/currycan/scanner/example/scanner.yaml /etc/scanner/scanner.yaml
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN apk add -U --no-cache \
        busybox-extras curl; \
    rm -rf /var/cache/apk/*

CMD ["scanner", "--config", "/etc/scanner/scanner.yaml"]
