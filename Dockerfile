FROM golang:latest

WORKDIR /go/src/chenyingdi/wePayTest/

COPY main.go .

RUN go build .

CMD ["./wePayTest"]