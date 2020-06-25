FROM golang:1.13

WORKDIR /app

COPY . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon
