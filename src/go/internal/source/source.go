package source

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"
)

//go:generate powershell -NoProfile -ExecutionPolicy Bypass -File ../../scripts/sync-iasi.ps1

//go:embed embedded/VERSION embedded/iasi/**
var embedded embed.FS

func Builtin() fs.FS {
	root, err := fs.Sub(embedded, "embedded/iasi")
	if err != nil {
		panic(err)
	}
	return root
}

func Version() (string, error) {
	data, err := fs.ReadFile(embedded, "embedded/VERSION")
	if err != nil {
		return "", fmt.Errorf("read embedded version: %w", err)
	}
	version := strings.TrimSpace(string(data))
	if version == "" {
		return "", fmt.Errorf("embedded version is empty")
	}
	return version, nil
}
