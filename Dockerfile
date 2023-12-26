FROM golang:1.21.2

WORKDIR /usr/local/app

COPY . .

EXPOSE 8000

RUN go mod tidy
