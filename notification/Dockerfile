# -------- Stage 1: Build the Go binary --------
    FROM golang:1.24-alpine AS builder

    # Disable CGO for pure static binary
    ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
    
    WORKDIR /notification
    
    # Copy go.mod and go.sum first for better caching
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the rest of the notificationlication
    COPY . .
    
    # Build the binary
    RUN go build -o notification cmd/collect/main.go
    
    # -------- Stage 2: Lightweight runtime --------
    FROM alpine:latest
    
    WORKDIR /notification
    
    # Copy just the binary and .env file
    COPY --from=builder /notification/notification .

    EXPOSE 8082
    EXPOSE 50052