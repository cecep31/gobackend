FROM golang:1.20-alpine AS dependencies
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Build the application
FROM dependencies AS build
COPY . .
RUN go build -o gobackend

# Create a smaller runtime image
FROM alpine
WORKDIR /app
COPY --from=build /app/gobackend .
EXPOSE 8080
CMD ["./gobackend"]



# for fly.io
# FROM golang:1.20-alpine
# WORKDIR /app
# COPY . .
# RUN go mod download
# RUN go build -o gobackend
# EXPOSE 8080
# CMD ["./gobackend"]
