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

RUN go build -o /client-bob

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /client-bob /client-bob

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/client-bob"]