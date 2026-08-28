package source

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"
)

//go:generate powershell -NoProfile -ExecutionPolicy Bypass -File ../../scripts/sync-iasi.ps1

// The build process synchronizes the repository's canonical IASI tree here.
//
//go:embed embedded/VERSION embedded/iasi/**
var methodology embed.FS

func Methodology() fs.FS {
	root, err := fs.Sub(methodology, "embedded")
	if err != nil {
		panic(err)
	}
	return root
}

func Version() (string, error) {
	data, err := fs.ReadFile(methodology, "embedded/VERSION")
	if err != nil {
		return "", fmt.Errorf("read embedded version: %w", err)
	}
	version := strings.TrimSpace(string(data))
	if version == "" {
		return "", fmt.Errorf("embedded version is empty")
	}
	return version, nil
}
