package install

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunCopiesSourceAndCreatesEmptyCategories(t *testing.T) {
	sourceRoot := t.TempDir()
	workspace := t.TempDir()
	writeTestFile(t, filepath.Join(sourceRoot, "iasi", "instructions", "general", "behavior.md"), "be clear")
	writeTestFile(t, filepath.Join(sourceRoot, "iasi", "instructions", "documentation", "guide.md"), "guide")
	writeTestFile(t, filepath.Join(sourceRoot, "iasi", "adapters", "copilot", "adapter.yml"), "id: copilot")

	path, err := Run(workspace, os.DirFS(sourceRoot), "0.2.0")
	if err != nil {
		t.Fatal(err)
	}
	for _, category := range categories {
		if info, err := os.Stat(filepath.Join(path, category)); err != nil || !info.IsDir() {
			t.Fatalf("category %s was not created", category)
		}
	}
	manifest, err := os.ReadFile(filepath.Join(path, "manifest.yml"))
	if err != nil || !strings.Contains(string(manifest), "schema_version: 1") || !strings.Contains(string(manifest), "version: 0.2.0") || strings.Contains(string(manifest), "profile:") {
		t.Fatalf("manifest was not created correctly: %v", err)
	}
	content, err := os.ReadFile(filepath.Join(path, "instructions", "general", "behavior.md"))
	if err != nil || string(content) != "be clear" {
		t.Fatalf("instruction was not copied: %v", err)
	}
	adapter, err := os.ReadFile(filepath.Join(path, "adapters", "copilot", "adapter.yml"))
	if err != nil || string(adapter) != "id: copilot" {
		t.Fatalf("adapter was not copied: %v", err)
	}
}

func TestRunRejectsExistingInstallation(t *testing.T) {
	workspace := t.TempDir()
	writeTestFile(t, filepath.Join(workspace, ".iasi", "manifest.yml"), "schema_version: 1\nversion: 0.2.0\n")
	if _, err := Run(workspace, os.DirFS(t.TempDir()), "0.2.0"); err == nil || !strings.Contains(err.Error(), "already installed") {
		t.Fatalf("expected existing installation error, got %v", err)
	}
}

func TestReinstallReplacesExistingInstallation(t *testing.T) {
	sourceRoot := t.TempDir()
	workspace := t.TempDir()
	writeTestFile(t, filepath.Join(sourceRoot, "iasi", "instructions", "general", "behavior.md"), "new rule")
	writeTestFile(t, filepath.Join(workspace, ".iasi", "instructions", "general", "behavior.md"), "old rule")
	writeTestFile(t, filepath.Join(workspace, ".iasi", "manifest.yml"), "schema_version: 1\nversion: 0.2.0\n")
	writeTestFile(t, filepath.Join(workspace, ".iasi", "validation.json"), "{\"status\":\"passed\"}")

	path, err := Reinstall(workspace, os.DirFS(sourceRoot), "0.3.0")
	if err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(filepath.Join(path, "instructions", "general", "behavior.md"))
	if err != nil || string(content) != "new rule" {
		t.Fatalf("installation was not reinstalled: %v", err)
	}
	state, err := os.ReadFile(filepath.Join(path, "validation.json"))
	if err != nil || string(state) != "{\"status\":\"passed\"}" {
		t.Fatalf("validation state was not preserved: %v", err)
	}
}

func TestRunPreservesValidationStateWithoutInstallation(t *testing.T) {
	sourceRoot := t.TempDir()
	workspace := t.TempDir()
	writeTestFile(t, filepath.Join(sourceRoot, "iasi", "instructions", "general", "behavior.md"), "rule")
	writeTestFile(t, filepath.Join(workspace, ".iasi", "validation.json"), "{\"status\":\"failed\"}")

	path, err := Run(workspace, os.DirFS(sourceRoot), "0.3.0")
	if err != nil {
		t.Fatal(err)
	}
	state, err := os.ReadFile(filepath.Join(path, "validation.json"))
	if err != nil || string(state) != "{\"status\":\"failed\"}" {
		t.Fatalf("validation state was not preserved: %v", err)
	}
}

func TestRunRejectsAmbiguousStateOnlyDirectory(t *testing.T) {
	workspace := t.TempDir()
	writeTestFile(t, filepath.Join(workspace, ".iasi", "instructions", "rule.md"), "unmanaged")
	if _, err := Run(workspace, os.DirFS(t.TempDir()), "0.3.0"); err == nil || !strings.Contains(err.Error(), "ambiguous") {
		t.Fatalf("expected ambiguous state-only directory failure, got %v", err)
	}
}

func writeTestFile(t *testing.T, path, content string) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
