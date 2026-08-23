VERSION=0.0.10
LDFLAGS=-ldflags "-w -s -X main.version=${VERSION}"
all: check_ssl_certificate2

.PHONY: check_ssl_certificate2 linux check lint

check_ssl_certificate2: *.go
	go build $(LDFLAGS) -o check_ssl_certificate2

linux: *.go
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o check_ssl_certificate2

check:
	go test -v ./...

lint:
	golangci-lint run --timeout 5m ./...