FROM golang:alpine

RUN apk add -U make
RUN mkdir -p /go/src/git.liquidweb.com/masre/liquidweb-go

WORKDIR /go/src/git.liquidweb.com/masre/liquidweb-go

COPY . .
RUN make build
