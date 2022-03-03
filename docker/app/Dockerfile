FROM golang:alpine

WORKDIR /handle-big-post-request

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./handle-big-post-request