#!/bin/sh

apk add --no-cache libzmq-dev
apk add pkgconfig

go mod download
go build -o profiler *.go
