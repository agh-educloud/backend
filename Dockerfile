FROM golang:latest

LABEL version="1.0"

RUN mkdir /go/src/app

RUN go get -u github.com/golang/dep/cmd/dep

ADD ./src/main.go /go/src/app

WORKDIR /go/src/app

RUN dep init
RUN go build

CMD ["./app"]

