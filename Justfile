#!/usr/bin/env -S just --justfile

default:
  @just --list


##
build: tidy
  @echo "Building editor..."
  go build -o build/editor ./main.go

test:
  go test -v -s ./editor/...

tidy:
    go mod tidy
    go fmt ./editor/...
    go vet ./editor/...