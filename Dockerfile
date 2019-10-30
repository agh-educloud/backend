FROM golang:latest

LABEL version="1.0"

WORKDIR src/app

#Copy files
COPY src src
COPY .env .
COPY protos protos/
COPY scripts utils/

#Enviroment
USER root
ENV GOPATH=/src/app/src
RUN chmod -R +x utils
RUN scripts/set_up_enviroment.sh

#Build
RUN scripts/generate_go_code.sh
RUN go get -d ./...
RUN go build main.go
