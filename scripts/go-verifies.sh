#!/usr/bin/env bash

readonly exec_pos="$(dirname "$0")/../"

cd "$exec_pos" \
&& go mod tidy \
&& go clean ./...\
&& go fix ./...\
&& go fmt ./...\
&& go vet ./... \
&& go test -gcflags=-l ./...
