FROM golang:1.27 as build-env

ENV NAME "tests"
WORKDIR /opt/${NAME}

COPY go.mod .
COPY go.sum .

COPY . .
