# Stage 1: Build the application
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
# CGO_ENABLED=0 prevents linking against C libraries
# -ldflags="-w -s" strips debug information, reducing binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main .

# Stage 2: Create the final, minimal image
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /app/main

# Optionally copy static assets or templates if needed later
# COPY templates ./templates
# COPY static ./static

# Expose the port the application will run on (matches docker-compose)
EXPOSE 8080

# Command to run the application
CMD ["/app/main"]
