# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /BBServerclient

COPY BBServerclient/go.mod ./
COPY BBServerclient/BBServer.go ./ 
# COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /bbserver

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build BBServerclient/BBServer.go /bbserver

EXPOSE 8095
EXPOSE 3333

USER nonroot:nonroot

ENTRYPOINT ["/bbserver"]
