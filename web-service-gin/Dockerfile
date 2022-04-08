FROM golang:1.18.0-alpine3.15 AS build

WORKDIR /web-service-gin
COPY .env ./
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY ./json-data ./

RUN go build -o /go-web-service

CMD ["/go-web-service"]