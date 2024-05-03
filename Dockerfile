# Start with a base Go image
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY . .

# Download dependencies using go mod
RUN go install

# Build the Go application
RUN go build -o app .

# Use a minimal base image to run the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/app .

# Command to run the application
CMD ["./app"]
