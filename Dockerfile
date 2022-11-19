FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go get

RUN go build

EXPOSE 8080

CMD ./gobackend
