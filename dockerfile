# Stage 1: Build the application
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
COPY *.go ./
RUN go mod download && \
    go build -o gobackend

# Stage 2: Create a smaller runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/gobackend .
EXPOSE 8080
CMD ["./gobackend"]

