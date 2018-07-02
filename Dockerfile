FROM golang:1.10 as builder

RUN go get -u -v golang.org/x/vgo

RUN mkdir /app
WORKDIR /app

ADD go.mod go.mod
ADD go.sum go.sum
ADD vendor vendor
ADD cmd cmd
ADD lib lib

RUN vgo build -tags=jsoniter -o server gochat/cmd/server

CMD ["./server"]
