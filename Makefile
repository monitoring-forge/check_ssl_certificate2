VERSION=0.0.9
GITCOMMIT?=$(shell git describe --dirty --always 2>/dev/null)
LDFLAGS=-ldflags "-w -s -X main.version=${VERSION} -X main.commit=${GITCOMMIT}"

all: check_ssl_certificate2

.PHONY: check_ssl_certificate2

check_ssl_certificate2: *.go
	go build $(LDFLAGS) -o check_ssl_certificate2

linux: *.go
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o check_ssl_certificate2

check:
	go test -v ./...

lint:
	golangci-lint run --timeout 5m ./...