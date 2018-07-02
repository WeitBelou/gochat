FROM golang:1.10 as builder

RUN go get -u -v golang.org/x/vgo

RUN mkdir /app
WORKDIR /app

ADD go.mod go.mod
ADD go.sum go.sum
ADD cmd cmd
ADD vendor vendor

RUN vgo build -tags=jsoniter -o server ./cmd/server

CMD ["./server"]
