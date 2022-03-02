##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app
ADD . /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /handle-big-post-request

##
## Deploy
##
FROM gcr.io/distroless/base-debian10


COPY --from=build /handle-big-post-request /handle-big-post-request

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/handle-big-post-request"]