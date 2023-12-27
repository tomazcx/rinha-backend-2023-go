FROM golang:1.21.2

WORKDIR /usr/local/app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./tmp/main ./cmd/api/main.go

ENTRYPOINT [ "./tmp/main" ]
