#! /bin/sh

set -e

if ! [ -x "$(command -v go)" ]; then
    echo "go is not installed"
    exit 1
fi
if ! [ -x "$(command -v git)" ]; then
    echo "git is not installed"
    exit 1
fi
if [ -z "${GOPATH}" ]; then
    echo "set GOPATH"
    exit 1
fi

export GO111MODULE=on

go clean --modcache
#go get github.com/wiqram/IG-Trading-Microservices/protos@main
go mod vendor
go mod verify
go install -v ./...