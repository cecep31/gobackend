FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go get

RUN go build

EXPOSE 80

CMD ./gobackend
