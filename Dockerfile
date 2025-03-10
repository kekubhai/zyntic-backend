# Use Go 1.23.5-alpine as the base image
FROM golang:1.23.5-alpine

# Set the working directory inside the container
WORKDIR /app

# Install necessary dependencies
RUN apk add --no-cache git make curl gcc g++ musl-dev

# Install Air (Live Reload Tool) from GitHub
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

# Ensure Air is accessible
ENV PATH="/usr/local/bin:${PATH}"

# Set environment variables for Go
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

# Copy only go.mod and go.sum first to leverage Docker caching
COPY go.mod go.sum ./

# Ensure we use Go 1.23 explicitly
RUN go mod edit -go=1.23

# Download dependencies before copying source code
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Build the Go application (optional step, can be removed for live reload)
RUN go build -v -o ./bin/main ./...

# Expose the application port
EXPOSE 3000

# Run the application using Air for hot reloading
CMD ["air", "-c", ".air.toml"]
