FROM golang:1.20-alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache make && \
    rm -rf /var/cache/apk/* \

RUN go install golang.org/x/tools/cmd/goimports@v0.8.0

RUN go install mvdan.cc/gofumpt@v0.5.0

ADD https://github.com/pressly/goose/releases/download/v3.7.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /
