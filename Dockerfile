FROM golang:1.21-alpine

WORKDIR /src/app

# Install build dependencies
RUN apk add --no-cache git make

# Copy the entire application including vendor directory
COPY . .

# No need to run go mod download since we're using the vendor directory

# Install Air for hot reload (without depending on go mod)
RUN go install github.com/cosmtrek/air@latest

# Build the application using vendored dependencies
RUN go build -mod=vendor -o ./bin/main ./cmd/api/main.go

EXPOSE 3000

# Use Air for hot reloading during development
CMD ["air", "-c", ".air.toml"]