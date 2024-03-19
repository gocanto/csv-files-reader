# Stage 1: Build the application and cache dependencies
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Create a minimal production image
FROM golang:1.22 AS production

WORKDIR /app

# Copy only the built executable from the previous stage
COPY --from=builder /app/main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
