# --- Builder Stage: Compile Go binary ---
FROM golang:1.24-bullseye AS builder

# Set working directory
WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application statically
RUN CGO_ENABLED=0 GOOS=linux go build -o telemetry .

# --- Final Stage: Minimal runtime image ---
FROM debian:bullseye-slim

# Set working directory
WORKDIR /app

# Install SQLite for DB access
RUN apt-get update && \
    apt-get install -y sqlite3 && \
    rm -rf /var/lib/apt/lists/*

# Copy the compiled binary from builder stage
COPY --from=builder /app/telemetry /app/telemetry

# Copy the persistent data folder (DB)
COPY data ./data

# Expose application port
EXPOSE 8080

# Default command to run the telemetry server
CMD ["./telemetry"]
