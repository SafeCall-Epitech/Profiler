#!/bin/sh


go mod download
go build -o profiler *.go
