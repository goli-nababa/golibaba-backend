FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install required packages
RUN apk add --no-cache tzdata

# Copy go mod files
COPY go.mod go.sum ./

# Copy vendor directory
COPY vendor/ vendor/

# Copy the rest of the source code
COPY . .

# Build the application with vendor
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o main ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy timezone data
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy config files
COPY config.json .

# Create logs directory
RUN mkdir -p /app/logs

EXPOSE 8081

CMD ["./main"]