# Stage 1: Build the application
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o dist/gobackend

# Stage 2: Create a smaller runtime image
FROM gcr.io/distroless/static:latest AS runtime
WORKDIR /app
COPY --from=builder /app/dist/gobackend .
EXPOSE 8080
CMD ["./gobackend"]

