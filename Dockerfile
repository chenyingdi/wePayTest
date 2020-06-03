FROM golang:latest

ENV GOBIN $GOROOT/bin
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

WORKDIR /go/src/chenyingdi/wePayTest/

COPY main.go .

RUN go build .

CMD ["./wePayTest"]