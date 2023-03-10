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

RUN go build -o /bbdatabase

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /bbdatabase .

# Add a shell to the container explicitly
RUN apt-get update && apt-get install -y bash

EXPOSE 3333

USER nobody

ENTRYPOINT ["/bbdatabase"]
