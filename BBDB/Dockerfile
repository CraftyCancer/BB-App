# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /BBDB

COPY BBDB/go.mod ./
COPY BBDB/BBDB.go ./
RUN go mod download

COPY *.go ./

RUN go build -o /bbdatabase && chmod +x /bbdatabase

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build BBDB/BBDB.go /bbdatabase

EXPOSE 3333

USER nobody

ENTRYPOINT ["/bbdatabase"]
