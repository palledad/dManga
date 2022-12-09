FROM golang:1.18.3

WORKDIR /backend

ARG GOPATH=/go

# goimportsを実行できるようにする
RUN go install golang.org/x/tools/cmd/goimports@latest

RUN mkdir /scripts \
    && wget 'https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh' -P /scripts  \
    && chmod 755 /scripts/wait-for-it.sh

COPY go.mod go.sum /backend/
RUN go mod download
