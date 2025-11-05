# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o geocoder ./cmd/server/server.go

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root

# Copy binary from builder
COPY --from=builder /build/geocoder .

CMD ["./geocoder"]

EXPOSE 8080 9092
