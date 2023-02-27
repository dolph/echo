# syntax=docker/dockerfile:1
FROM golang:1.19-alpine
WORKDIR /app

# Fetch dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Build from source
COPY *.go ./
RUN go build -o "/echo"

EXPOSE 8080

CMD [ "/echo" ]
