# -------- Stage 1: Build the Go binary --------
    FROM golang:1.24-alpine AS builder

    # Disable CGO for pure static binary
    ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
    
    WORKDIR /logger
    
    # Copy go.mod and go.sum first for better caching
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the rest of the loggerlication
    COPY . .
    
    # Build the binary
    RUN go build -o logger cmd/collect/main.go
    
    # -------- Stage 2: Lightweight runtime --------
    FROM alpine:latest
    
    WORKDIR /logger
    
    # Copy just the binary and .env file
    COPY --from=builder /logger/logger .

    EXPOSE 8085
    EXPOSE 50055