package cli

import "testing"

func TestParseArgumentsAllowsFirstUnnamed(t *testing.T) {
	arguments, err := ParseArguments([]string{"codex", "dir=/tmp", "overwrite=true"})
	if err != nil {
		t.Fatalf("ParseArguments returned error: %v", err)
	}

	if arguments.Unnamed == nil || *arguments.Unnamed != "codex" {
		t.Fatalf("unexpected unnamed argument: %#v", arguments.Unnamed)
	}

	if arguments.Named["dir"] != "/tmp" {
		t.Fatalf("unexpected dir: %q", arguments.Named["dir"])
	}

	if arguments.Named["overwrite"] != "true" {
		t.Fatalf("unexpected overwrite: %q", arguments.Named["overwrite"])
	}
}

func TestParseArgumentsRejectsUnnamedAfterNamed(t *testing.T) {
	_, err := ParseArguments([]string{"dir=/tmp", "codex"})
	if err == nil {
		t.Fatal("expected error")
	}
}
