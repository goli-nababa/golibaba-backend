FROM golang:1.22.2-alpine AS builder

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server ./cmd/main.go

FROM debian:buster-slim 
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
COPY --from=builder /app/config.json /app/config.json

ENV CONFIG_FILE /app/config.json
ENV APP_ENV docker

CMD [ "/app/server" ]
 