package status

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFindFromCurrentDirectory(t *testing.T) {
	workspace := t.TempDir()
	writeStatusFile(t, filepath.Join(workspace, ".iasi", "manifest.yml"), "schema_version: 1\nversion: 0.1.0\n")
	writeStatusFile(t, filepath.Join(workspace, ".iasi", "instructions", "one.md"), "---\nid: one\nstatus: active\nscope: general\n---\none")
	writeStatusFile(t, filepath.Join(workspace, ".iasi", "instructions", "nested", "two.md"), "---\nid: two\nstatus: active\nscope: general\n---\ntwo")
	writeStatusFile(t, filepath.Join(workspace, ".iasi", "commands", "run.md"), "run")

	result, err := Find(workspace)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Layers) != 1 || result.Layers[0].Path != filepath.Join(workspace, ".iasi") || result.Layers[0].Version != "0.1.0" {
		t.Fatalf("unexpected installation: %+v", result)
	}
	if result.Counts["instructions"] != 2 || result.Counts["commands"] != 1 || result.Counts["skills"] != 0 || result.Counts["mcp"] != 0 {
		t.Fatalf("unexpected counts: %+v", result.Counts)
	}
}

func TestFormatShowsInstalledAndBinaryVersions(t *testing.T) {
	result := Result{Layers: []Layer{{Path: `C:\workspace\.iasi`, Version: "0.1.0"}}, Counts: map[string]int{}}
	output := Format(result, "0.2.0")
	if !strings.Contains(output, "C:\\workspace\\.iasi  0.1.0") || !strings.Contains(output, "Binary : 0.2.0") {
		t.Fatalf("expected both versions in output: %s", output)
	}
}

func TestFindAscendsToParent(t *testing.T) {
	workspace := t.TempDir()
	child := filepath.Join(workspace, "project", "src")
	if err := os.MkdirAll(child, 0o755); err != nil {
		t.Fatal(err)
	}
	writeStatusFile(t, filepath.Join(workspace, ".iasi", "manifest.yml"), "schema_version: 1\nversion: 0.1.0\n")

	result, err := Find(child)
	if err != nil || len(result.Layers) != 1 || result.Layers[0].Path != filepath.Join(workspace, ".iasi") {
		t.Fatalf("expected parent installation, got %+v, %v", result, err)
	}
}

func TestFindReportsMissingInstallation(t *testing.T) {
	_, err := Find(t.TempDir())
	if !errors.Is(err, ErrNotInstalled) {
		t.Fatalf("expected ErrNotInstalled, got %v", err)
	}
}

func writeStatusFile(t *testing.T, path, content string) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
