.PHONY: sync-builtin build test clean

sync-builtin:
	go run ./scripts/sync-builtin.go

build: sync-builtin
	go build -ldflags "-X main.version=$${IASI_VERSION:-dev}" -o ../../bin/iasi ./cmd/iasi

test: sync-builtin
	go test ./...

clean:
	go run ./scripts/sync-builtin.go --clean
