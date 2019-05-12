FROM golang:latest

LABEL version="1.0"

RUN mkdir /go/src/app

#Downloading dep and protoc
RUN go get -u github.com/golang/dep/cmd/dep
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
RUN apt-get update
RUN apt-get install unzip
RUN unzip protoc-3.6.1-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/local/bin/
RUN mv protoc3/include/* /usr/local/include/
RUN rm -rf protoc-3.6.1-linux-x86_64.zip -d protoc3
RUN go get -u github.com/golang/protobuf/protoc-gen-go

#Copy files
COPY main /go/src/app
COPY protos /go/src/app/protos
COPY generate_go_code.sh /go/src/app

WORKDIR /go/src/app

#Build
RUN dep init
RUN ./generate_go_code.sh
RUN go build

CMD ["./app"]

