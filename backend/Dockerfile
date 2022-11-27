FROM golang:1.18.3

WORKDIR /backend

ARG GOPATH=/go

# goimportsを実行できるようにする
RUN go install golang.org/x/tools/cmd/goimports@latest

COPY go.mod go.sum /backend/
RUN go mod download
