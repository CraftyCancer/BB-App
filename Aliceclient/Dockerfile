# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /

 COPY go.mod ./
# COPY go.sum ./
 RUN go mod download

COPY *.go ./

RUN go build -o /client-alice

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /client-alice /client-alice

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/client-alice"]