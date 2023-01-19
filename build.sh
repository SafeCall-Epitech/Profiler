#!/bin/sh

apt-get install libzmq3-dev
apk add pkgconfig

go mod download
go build -o profiler *.go
