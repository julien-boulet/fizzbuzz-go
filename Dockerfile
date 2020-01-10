# Start from the latest golang base image
FROM golang:alpine as builder

# Installing librdkafka
RUN apk add --update --no-cache alpine-sdk bash python ca-certificates libressl tar git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils librdkafka-dev pkgconf

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=true GOOS=linux go build -a -installsuffix cgo -o main .


######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy db migrations file
COPY --from=builder /app/db/migrations ./db/migrations

# Re-Installing librdkafka for execution
RUN apk add --update --no-cache librdkafka-dev pkgconf

WORKDIR /root/

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]