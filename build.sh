#!/bin/sh
apk update
apk add pkgconfig libzmq zeromq-dev gcc musl-dev

go mod download
mv ../usr/lib/libzmq.so.5.2.4 .
mv ../usr/lib/libzmq.so.5 .
mv ../usr/lib/libzmq.so .
go build -o profiler *.go
