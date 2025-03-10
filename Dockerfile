FROM golang:1.23.5-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make curl

# Download pre-built Air binary
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

# Ensure /usr/local/bin is in PATH
ENV PATH="/usr/local/bin:${PATH}"

# Set environment variables for Go
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

# Copy Go module files first for better caching
COPY go.mod go.sum* ./

# Update go.mod to use Go 1.21
RUN go mod edit -go=1.21

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -v -o ./bin/main ./...

EXPOSE 3000

# Use Air for hot reloading
CMD ["air", "-c", ".air.toml"]