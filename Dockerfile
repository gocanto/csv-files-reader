
# Stage 1: Build the application and cache dependencies
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .

# Stage 2: Create a minimal production image
FROM golang:1.22 AS production
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
