package builtin

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestCopyFSRecursivelyCopiesTree(t *testing.T) {
	sourceFS := fstest.MapFS{
		"adapters/codex/root.txt":        {Data: []byte("root")},
		"adapters/codex/nested/file.txt": {Data: []byte("nested")},
	}
	destination := filepath.Join(t.TempDir(), ".codex")

	if err := copyFS(sourceFS, "adapters/codex", destination, false); err != nil {
		t.Fatalf("copyFS returned error: %v", err)
	}

	assertFileContent(t, filepath.Join(destination, "root.txt"), "root")
	assertFileContent(t, filepath.Join(destination, "nested", "file.txt"), "nested")
}

func TestCopyFSRejectsExistingFileWithoutOverwrite(t *testing.T) {
	sourceFS := fstest.MapFS{"adapters/codex/file.txt": {Data: []byte("new")}}
	destination := filepath.Join(t.TempDir(), ".codex")
	if err := os.MkdirAll(destination, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(destination, "file.txt"), []byte("old"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := copyFS(sourceFS, "adapters/codex", destination, false); err == nil {
		t.Fatal("expected error")
	}
}

func TestCopyFSOverwritesExistingFile(t *testing.T) {
	sourceFS := fstest.MapFS{"adapters/codex/file.txt": {Data: []byte("new")}}
	destination := filepath.Join(t.TempDir(), ".codex")
	if err := os.MkdirAll(destination, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(destination, "file.txt"), []byte("old content that is longer"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := copyFS(sourceFS, "adapters/codex", destination, true); err != nil {
		t.Fatalf("copyFS returned error: %v", err)
	}

	assertFileContent(t, filepath.Join(destination, "file.txt"), "new")
}

func assertFileContent(t *testing.T, filename string, expected string) {
	t.Helper()
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != expected {
		t.Fatalf("unexpected content: %q", data)
	}
}

var _ fs.FS = fstest.MapFS{}
