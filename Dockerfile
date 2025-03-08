FROM golang:1.16.3-alpine3.13 

WORKDIR /src/app

RUN go install github.com/comstrek/air@latest

COPY . .
RUN go mod tidy