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

RUN go build -o /bbdatabase

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /bbdatabase /bbdatabase

EXPOSE 3333

USER nonroot:nonroot

ENTRYPOINT ["/bbdatabase"]