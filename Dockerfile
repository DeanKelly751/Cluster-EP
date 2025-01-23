# Start with a Go base image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /cluster-ep-poc

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum are not changed
RUN go mod download

# Copy the source code into the container
COPY *.go ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /cluster-ep-poc

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /cluster-ep-poc .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./cluster-ep-poc"]
