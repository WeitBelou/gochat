FROM golang:1.10

RUN go get github.com/constabulary/gb/...

WORKDIR /app

ADD vendor vendor
ADD src src

RUN gb build -tags=jsoniter
RUN cp bin/server-jsoniter server

CMD ["./server"]
