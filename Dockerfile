# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
COPY config.json ./
RUN go mod download

COPY *.go ./

RUN go build *.go

# RUN go build -o /main.go /handler.go /query.go

EXPOSE 8080

CMD [ "./main" ]
