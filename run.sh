#!/bin/sh

export PATH=$(go env GOPATH)/bin:$PATH

swag init

go build -o main

./main
