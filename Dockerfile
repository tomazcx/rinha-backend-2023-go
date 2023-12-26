FROM golang:1.21.2

WORKDIR /usr/local/app

COPY . .

RUN go mod tidy
