.PHONY: build

build:
	go build -ldflags "-X main.version=$${IASI_VERSION:-dev}" -o ../../bin/iasi ./cmd/iasi
