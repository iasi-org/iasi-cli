package runtime

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWorkflowRuntimeTransitionsAndRejectsInvalidGate(t *testing.T) {
	project := t.TempDir()
	writeFile(t, filepath.Join(project, ".iasi", "manifest.yml"), "schema_version: 1\nversion: 0.1.0\n")
	writeFile(t, filepath.Join(project, ".iasi", "instructions", "rule.md"), "---\nid: rule\nstatus: active\nscope: general\n---\nrule")
	writeFile(t, filepath.Join(project, ".iasi", "commands", "validate.md"), "validate")
	if _, err := Run(project, []string{"workflow", "require", "plan", "INPUTS_VALIDATED"}); err == nil || err.(*Error).Code != 1 {
		t.Fatalf("expected gate rejection, got %v", err)
	}
	if _, err := Run(project, []string{"validate", "passed", "0", "0"}); err != nil {
		t.Fatal(err)
	}
	if output, err := Run(project, []string{"workflow", "status"}); err != nil || output != "INPUTS_VALIDATED\n" {
		t.Fatalf("unexpected status %q (%v)", output, err)
	}
	if _, err := Run(project, []string{"workflow", "transition", "plan", "PLANNED"}); err != nil {
		t.Fatal(err)
	}
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
