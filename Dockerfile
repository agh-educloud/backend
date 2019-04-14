FROM golang:latest

ADD ./src /go/src/app
WORKDIR /go/src/app

CMD ["go", "run", "main.go"]
