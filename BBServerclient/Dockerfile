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

RUN go build -o /bbserver

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /bbserver /bbserver

EXPOSE 8080
EXPOSE 3333

USER nonroot:nonroot

ENTRYPOINT ["/bbserver"]