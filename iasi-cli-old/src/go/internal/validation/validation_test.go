package validation

import (
	"os"
	"path/filepath"
	"testing"

	"iasi-cli/internal/resolver"
)

func TestStateRoundTripAndStalenessGate(t *testing.T) {
	project := t.TempDir()
	writeFile(t, filepath.Join(project, "inputs", "requirements.md"), "Authentication is required.")
	context := testContext("Use secure authentication.", "Validate requirements.")
	state, err := NewState(project, context, "passed", 0, 1)
	if err != nil {
		t.Fatal(err)
	}
	if err := Write(project, state); err != nil {
		t.Fatal(err)
	}
	if err := RequireCurrent(project, context); err != nil {
		t.Fatalf("expected current state: %v", err)
	}
	writeFile(t, filepath.Join(project, "inputs", "requirements.md"), "Authentication is optional.")
	if err := RequireCurrent(project, context); err == nil {
		t.Fatal("expected stale validation after input change")
	}
}

func TestStateRequiresEffectiveValidateCommand(t *testing.T) {
	if _, err := NewState(t.TempDir(), resolver.Context{}, "passed", 0, 0); err == nil {
		t.Fatal("expected missing validate command failure")
	}
}

func TestHashInputsIgnoresArchivedSubtrees(t *testing.T) {
	project := t.TempDir()
	writeFile(t, filepath.Join(project, "inputs", "externals", "active.md"), "active")
	baseline, err := HashInputs(project)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(project, "inputs", "externals", "archived", "old.md"), "historical")
	writeFile(t, filepath.Join(project, "inputs", "internals", "archived", "old.md"), "historical")
	writeFile(t, filepath.Join(project, "inputs", "obtained", "archived", "old.md"), "historical")
	withArchives, err := HashInputs(project)
	if err != nil {
		t.Fatal(err)
	}
	if baseline != withArchives {
		t.Fatal("archived input content changed the active input hash")
	}
	if err := os.Rename(filepath.Join(project, "inputs", "externals", "active.md"), filepath.Join(project, "inputs", "externals", "archived", "active.md")); err != nil {
		t.Fatal(err)
	}
	afterArchive, err := HashInputs(project)
	if err != nil {
		t.Fatal(err)
	}
	if afterArchive == baseline {
		t.Fatal("archiving an active input did not change the active input hash")
	}
}

func testContext(instruction, command string) resolver.Context {
	return resolver.Context{
		Instructions: map[string]resolver.Instruction{"rule": {ID: "rule", Status: "active", Scope: "general", Body: instruction}},
		Commands:     map[string]resolver.Command{"validate": {ID: "validate", Content: command}},
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
