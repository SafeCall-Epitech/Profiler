#!/bin/sh
apk update
apk add pkgconfig libzmq zeromq-dev gcc musl-dev

go mod download
go build -o profiler *.go