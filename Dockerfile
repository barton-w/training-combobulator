# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /code

COPY . /code

WORKDIR /code/cmd

RUN go build -o training-combobulator

WORKDIR /code

CMD ["cmd/training-combobulator"]