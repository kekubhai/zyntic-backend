FROM golang:1.23.5-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make curl

# Set environment variables for Go
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

# Copy Go module files first for better caching
COPY go.mod go.sum* ./

# Update go.mod to use Go 1.21
RUN go mod edit -go=1.21

# Copy the rest of the application
COPY . .

# Build the application directly (don't install Air)
RUN go build -v -o ./bin/main ./cmd/api/main.go

EXPOSE 3000

# Run the application directly without Air
CMD ["./bin/main"]