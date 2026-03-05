# ---------- STAGE 1: Build ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (needed for some modules)
RUN apk add --no-cache git

# Copy go mod files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of project
COPY . .

# Build optimized binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o app ./cmd

# ---------- STAGE 2: Runtime ----------
FROM alpine:latest

WORKDIR /root/

# Install certificates (needed for HTTPS calls)
RUN apk add --no-cache ca-certificates

# Copy only binary from builder
COPY --from=builder /app/app .

# Expose port
EXPOSE 8080

# Run app
CMD ["./app"]