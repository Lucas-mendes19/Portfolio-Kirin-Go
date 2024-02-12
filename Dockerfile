# Build stage
FROM golang:alpine AS build

WORKDIR /app

# Copy only the necessary files for 'go mod download'
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /http-api ./cmd/server/main.go

# Final stage
FROM alpine:3.16

WORKDIR /

# Copy the binary from the build stage
COPY --from=build /http-api /http-api

CMD ["/http-api"]