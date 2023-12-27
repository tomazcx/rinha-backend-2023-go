FROM golang:1.21.2 as build

WORKDIR /usr/local/app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./tmp/main ./cmd/api/main.go

FROM alpine

WORKDIR /usr/local/app

COPY .env go.mod go.sum ./

COPY --from=build /usr/local/app/tmp/main ./

ENTRYPOINT [ "./main" ]
