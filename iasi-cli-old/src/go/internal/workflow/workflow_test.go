package workflow

import (
	"os"
	"path/filepath"
	"testing"

	"iasi-cli/internal/resolver"
)

func TestPlanGateAndInputInvalidation(t *testing.T) {
	project := t.TempDir()
	writeFile(t, filepath.Join(project, "inputs", "externals", "scope.md"), "scope")
	context := testContext()
	if _, err := Require(project, context, "plan", InputsValidated); err == nil {
		t.Fatal("plan ran before input validation")
	}
	if err := Succeed(project, context, "validate", InputsValidated); err != nil {
		t.Fatal(err)
	}
	if _, err := Require(project, context, "plan", InputsValidated); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(project, "inputs", "externals", "scope.md"), "changed scope")
	if _, err := Require(project, context, "plan", InputsValidated); err == nil {
		t.Fatal("stale validation authorized planning")
	}
}

func TestFailureBlocksLaterStages(t *testing.T) {
	project := t.TempDir()
	context := testContext()
	if err := Succeed(project, context, "validate", InputsValidated); err != nil {
		t.Fatal(err)
	}
	if err := Fail(project, context, "plan"); err != nil {
		t.Fatal(err)
	}
	if _, err := Require(project, context, "execute", PlanValidated); err == nil {
		t.Fatal("failed plan did not block execution")
	}
}

func TestFullWorkflowTransitions(t *testing.T) {
	project := t.TempDir()
	context := testContext()
	if err := Succeed(project, context, "validate", InputsValidated); err != nil {
		t.Fatal(err)
	}
	if _, err := Require(project, context, "plan", InputsValidated); err != nil {
		t.Fatal(err)
	}
	if err := Succeed(project, context, "plan", Planned); err != nil {
		t.Fatal(err)
	}
	if err := Succeed(project, context, "validate", PlanValidated); err != nil {
		t.Fatal(err)
	}
	if _, err := Require(project, context, "execute", PlanValidated); err != nil {
		t.Fatal(err)
	}
	if err := Succeed(project, context, "execute", Executed); err != nil {
		t.Fatal(err)
	}
	if _, err := Require(project, context, "verify", Executed); err != nil {
		t.Fatal(err)
	}
	if err := Succeed(project, context, "verify", Verified); err != nil {
		t.Fatal(err)
	}
	state, err := Current(project, context)
	if err != nil || state.Checkpoint != Verified || state.FailedCommand != nil {
		t.Fatalf("unexpected final state: %#v (%v)", state, err)
	}
}

func TestCorrectiveExecutionAfterVerificationFailure(t *testing.T) {
	project := t.TempDir()
	context := testContext()
	if err := Succeed(project, context, "validate", InputsValidated); err != nil {
		t.Fatal(err)
	}
	if err := Succeed(project, context, "plan", Planned); err != nil {
		t.Fatal(err)
	}
	if err := Succeed(project, context, "validate", PlanValidated); err != nil {
		t.Fatal(err)
	}
	if err := Succeed(project, context, "execute", Executed); err != nil {
		t.Fatal(err)
	}
	if err := Fail(project, context, "verify"); err != nil {
		t.Fatal(err)
	}
	if _, err := Require(project, context, "execute", PlanValidated); err != nil {
		t.Fatal(err)
	}
	if err := Succeed(project, context, "execute", Executed); err != nil {
		t.Fatal(err)
	}
}

func testContext() resolver.Context {
	return resolver.Context{Instructions: map[string]resolver.Instruction{"rule": {ID: "rule", Status: "active", Scope: "general", Body: "rule"}}, Commands: map[string]resolver.Command{"validate": {ID: "validate", Content: "validate"}}}
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
