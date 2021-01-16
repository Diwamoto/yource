from golang:1.16-rc-alpine3.12
WORKDIR /go/app
ENV GO111MODULE=off
#GIN_MODE = release 
ENV GIN_MODE=debug
RUN set -eux && \
  apk update && \
  apk add --no-cache git curl bash vim gcc alpine-sdk build-base && \
  go get -u github.com/codegangsta/gin && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/gin github.com/codegangsta/gin
ENV GO111MODULE on