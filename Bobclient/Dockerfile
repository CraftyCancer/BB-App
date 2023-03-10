# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /Bobclient

COPY Bobclient/go.mod ./
COPY Bobclient/bobClient.go ./
RUN go mod download

COPY *.go ./

RUN go build -o /client-bob

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /client-bob /client-bob

RUN chmod +x /client-bob

EXPOSE 8095

USER nobody

ENTRYPOINT ["/client-bob"]
