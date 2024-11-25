# Use a lightweight Go image
FROM golang:1.22.0-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN go build -o ngrok_clone

# Use a minimal image for running the app
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/ngrok_clone .
COPY --from=builder /app/.env .
EXPOSE 8777
CMD ["./authentication-server"]
