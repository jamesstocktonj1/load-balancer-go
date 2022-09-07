# syntax=docker/dockerfile:1
FROM golang:1.18

WORKDIR /app


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /load-database

EXPOSE 3000

CMD [ "/load-database" ]