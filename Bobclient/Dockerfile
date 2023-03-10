# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /Bobclient

COPY Bobclient/go.mod ./
COPY Bobclient/bobClient.go ./
# COPY go.sum ./
 RUN go mod download

COPY *.go ./

RUN go build -o /client-bob

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build Bobclient/bobClient.go /client-bob

EXPOSE 8095

USER nobody

ENTRYPOINT ["/client-bob"]
