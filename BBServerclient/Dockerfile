# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /BBServerClient

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

COPY --from=build BBServer/BBServer.go /bbserver

EXPOSE 8080
EXPOSE 3333

USER nonroot:nonroot

ENTRYPOINT ["/bbserver"]
