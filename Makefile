export GO111MODULE=on
LDFLAGS := -s -w

all: env fmt build

build: mojito

env:
	@go version

fmt:
	go fmt ./...

test:
	go test

mojito:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/mojito ./cmd/mojito

clean:
	rm -f ./bin/mojito