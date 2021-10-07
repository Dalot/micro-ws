#!/bin/sh

export GO111MODULE=on
export CGO_ENABLED=1
go test -mod vendor -v -race -cover ./..