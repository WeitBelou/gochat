FROM golang:1.10-alpine3.7 as builder

RUN apk add --no-cache git
RUN go get -u -v golang.org/x/vgo

RUN mkdir /app
WORKDIR /app

ADD cmd cmd
ADD go.mod go.mod
ADD go.sum go.sum

RUN vgo build -tags=jsoniter ./cmd/...

FROM alpine:3.7
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/server .
CMD ["./server"]
