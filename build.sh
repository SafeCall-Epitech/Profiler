#!/bin/sh

apk add pkgconfig

go mod download
go build -o profiler *.go