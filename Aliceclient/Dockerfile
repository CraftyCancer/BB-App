# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /Aliceclient

COPY Aliceclient/go.mod ./
COPY Aliceclient/aliceClient.go ./

RUN go mod download

COPY *.go ./

RUN go build -o /Aliceclient/aliceClient.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /Aliceclient/aliceClient.go /client-alice

EXPOSE 8095

USER nonroot:nonroot

CMD ["./client-alice"]

# ENTRYPOINT ["/client-alice"]

