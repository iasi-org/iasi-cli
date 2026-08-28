package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunReinstallReplacesLocalInstallation(t *testing.T) {
	workspace := t.TempDir()
	if err := os.MkdirAll(filepath.Join(workspace, ".iasi"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(workspace, ".iasi", "manifest.yml"), []byte("schema_version: 1\nversion: 0.1.0\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	originalDirectory, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(workspace); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(originalDirectory) })

	if err := run([]string{"reinstall"}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(workspace, ".iasi", "manifest.yml")); err != nil {
		t.Fatalf("reinstall did not install IASI: %v", err)
	}
}

func TestRunInstallCommandsRejectArguments(t *testing.T) {
	for _, command := range []string{"install", "reinstall"} {
		err := run([]string{command, "--workspace"})
		if err == nil || err.Error() != "usage: iasi "+command {
			t.Fatalf("expected %s to reject arguments, got %v", command, err)
		}
	}
}
