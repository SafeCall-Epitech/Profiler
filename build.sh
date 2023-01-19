#!/bin/sh
apk update
apk add pkgconfig libzmq zeromq-dev gcc musl-dev

go mod download
# ls ../usr/lib
# cp ../usr/lib/libzmq.so.5.2.4 .
# cp ../usr/lib/libzmq.so.5 .
# cp ../usr/lib/libzmq.so .
go build -o profiler *.go
