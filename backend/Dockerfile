FROM golang:1.18.3

WORKDIR /backend

ARG GOPATH=/go

COPY go.mod go.sum /backend/
RUN go mod download
