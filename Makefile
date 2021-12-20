.PHONY: test

QAE_APPROOT=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
export QAE_APPROOT

test:
	go test -v ./...
