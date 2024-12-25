FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app cmd/main.go

FROM alpine:latest

WORKDIR /usr/local/bin
COPY --from=builder /go/bin/app /usr/local/bin/app

EXPOSE 8080

CMD ["/usr/local/bin/app", "--config", "/usr/local/bin/config.json"]