FROM golang:latest

LABEL version="1.0"

#Copy files
COPY main /go/src/app/main
COPY .env /go/src/app
COPY protos /go/src/app/protos
COPY utils /go/src/app/utils

WORKDIR /go/src/app

#Enviroment
USER root
RUN chmod -R +x utils
RUN utils/set_up_enviroment.sh

#Build
RUN utils/generate_go_code.sh
RUN dep init

WORKDIR /go/src/app/main
RUN go build

CMD ["./app"]

