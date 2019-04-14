FROM golang:latest

LABEL version="1.0"

RUN mkdir /go/src/app

RUN go get -u github.com/golang/dep/cmd/dep

ADD ./src/main.go /go/src/app
#COPY ./Gopkg.toml /go/src/app # TODO <- this should be instead of dep init

WORKDIR /go/src/app

RUN dep init
RUN dep ensure
RUN go test -v
RUN go build

CMD ["./app"]

