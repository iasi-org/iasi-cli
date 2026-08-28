package resolver

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveComposesLayersWithNearestPrecedence(t *testing.T) {
	workspace := t.TempDir()
	project := filepath.Join(workspace, "project")
	writeLayerFile(t, filepath.Join(workspace, ".iasi", "manifest.yml"), "schema_version: 1\nversion: 0.1.0\n")
	writeLayerFile(t, filepath.Join(workspace, ".iasi", "instructions", "parent.md"), instruction("parent", "general", "Parent"))
	writeLayerFile(t, filepath.Join(workspace, ".iasi", "instructions", "shared.md"), instruction("shared", "general", "Old"))
	writeLayerFile(t, filepath.Join(workspace, ".iasi", "commands", "review.md"), "Parent review")
	writeLayerFile(t, filepath.Join(workspace, ".iasi", "commands", "validate.md"), "Parent validate")
	writeAdapter(t, filepath.Join(workspace, ".iasi", "adapters", "copilot", "adapter.yml"), "copilot")

	writeLayerFile(t, filepath.Join(project, ".iasi", "manifest.yml"), "schema_version: 1\nversion: 0.2.0\n")
	writeLayerFile(t, filepath.Join(project, ".iasi", "instructions", "shared.md"), instruction("shared", "general", "New"))
	writeLayerFile(t, filepath.Join(project, ".iasi", "instructions", "child.md"), instruction("child", "general", "Child"))
	writeLayerFile(t, filepath.Join(project, ".iasi", "commands", "validate.md"), "Child validate")
	writeAdapter(t, filepath.Join(project, ".iasi", "adapters", "copilot", "adapter.yml"), "copilot")

	context, err := Resolve(filepath.Join(project, "src"))
	if err != nil {
		t.Fatal(err)
	}
	if len(context.Layers) != 2 || context.Layers[0].Version != "0.1.0" || context.Layers[1].Version != "0.2.0" {
		t.Fatalf("unexpected layer order: %#v", context.Layers)
	}
	if len(context.Instructions) != 3 || context.Instructions["shared"].Body != "New" {
		t.Fatalf("instructions were not composed: %#v", context.Instructions)
	}
	if len(context.Commands) != 2 || context.Commands["validate"].Content != "Child validate" {
		t.Fatalf("commands were not composed: %#v", context.Commands)
	}
	adapter, ok := context.Adapters["copilot"]
	if !ok || adapter.Path != filepath.Join(project, ".iasi", "adapters", "copilot") {
		t.Fatalf("child adapter did not replace parent atomically: %#v", adapter)
	}
}

func TestResolveRejectsDuplicateInstructionIDInLayer(t *testing.T) {
	workspace := t.TempDir()
	writeLayerFile(t, filepath.Join(workspace, ".iasi", "manifest.yml"), "schema_version: 1\nversion: 0.1.0\n")
	writeLayerFile(t, filepath.Join(workspace, ".iasi", "instructions", "first.md"), instruction("duplicate", "general", "First"))
	writeLayerFile(t, filepath.Join(workspace, ".iasi", "instructions", "second.md"), instruction("duplicate", "general", "Second"))
	if _, err := Resolve(workspace); err == nil {
		t.Fatal("expected duplicate instruction ID failure")
	}
}

func instruction(id, scope, body string) string {
	return "---\nid: " + id + "\nstatus: active\nscope: " + scope + "\n---\n" + body
}

func writeAdapter(t *testing.T, path, id string) {
	writeLayerFile(t, path, "schema_version: 1\nid: "+id+"\nplatform: test\n")
}

func writeLayerFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
