# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /src
RUN export GO111MODULE="on"

COPY go.mod ./
COPY go.sum ./
COPY config.json ./
RUN go mod download

COPY . .

# RUN go build *.go

# RUN go build -o /main.go /handler.go /query.go

EXPOSE 8080

# CMD [ "./main" ]
