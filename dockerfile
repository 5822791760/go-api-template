# Step 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Download necessary Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/backend

# Step 2: Create the final lightweight image
FROM alpine:3.18

# Set up a non-root user (security best practice)
RUN adduser -D -g '' appuser

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/app /usr/local/bin/app

# Change ownership to non-root user
RUN chown appuser:appuser /usr/local/bin/app

# Switch to non-root user
USER appuser

# Expose port (replace 8080 with the actual port your app uses)
EXPOSE 3000

# Start the Go application
ENTRYPOINT ["/usr/local/bin/app"]