FROM golang:1.17 as build-env

ENV NAME "tests"
WORKDIR /opt/${NAME}

COPY go.mod .
COPY go.sum .

COPY . .
