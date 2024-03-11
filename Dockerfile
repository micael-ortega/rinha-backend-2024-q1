FROM golang:1.22 AS base

WORKDIR /app

COPY ./go.mod ./go.sum ./

COPY ./cmd ./cmd

COPY ./internal ./internal

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o crebito ./cmd

FROM debian AS build

WORKDIR /go/api

COPY --from=base /go/api/crebito ./

COPY .env .

# RUN apt update && apt upgrade

RUN chmod +x ./crebito

CMD [ "/go/api/crebito" ]



