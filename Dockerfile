FROM golang:latest

LABEL version="1.0"

WORKDIR src/app

#Copy files
COPY vendor vendor
COPY main.go main.go
COPY .env .
COPY protos protos/
COPY utils utils/

#Enviroment
USER root
RUN chmod -R +x utils
RUN utils/set_up_enviroment.sh

#Build
RUN utils/generate_go_code.sh
RUN go get -d ./...
RUN go build main.go
