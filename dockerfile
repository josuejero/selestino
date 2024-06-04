# Dockerfile

# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o selestino ./cmd

# Expose the application port
EXPOSE 8080

# Set the entry point to run the application
CMD ["./selestino"]
