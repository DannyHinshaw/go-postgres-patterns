FROM golang:1.20-alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache make && \
    rm -rf /var/cache/apk/* \

ADD https://github.com/pressly/goose/releases/download/v3.7.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /
