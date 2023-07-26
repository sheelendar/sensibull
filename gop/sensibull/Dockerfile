# Dockerfile

# Use the official Go image as the base image
FROM golang:1.20.5-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /sensibuill

# Expose the port on which the Go application listens
EXPOSE 19093

# Command to run the Go application
CMD ["/sensibuill"]