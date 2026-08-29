package cli

import "testing"

func TestParseArgumentsAllowsCanonicalSyntax(t *testing.T) {
	arguments, err := ParseArguments([]string{"--list", "codex", "dir=/tmp", "overwrite=true"})
	if err != nil {
		t.Fatalf("ParseArguments returned error: %v", err)
	}

	if !arguments.Options["--list"] {
		t.Fatal("--list option not parsed")
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

func TestParseArgumentsAllowsOptionWithoutUnnamed(t *testing.T) {
	arguments, err := ParseArguments([]string{"--list"})
	if err != nil {
		t.Fatalf("ParseArguments returned error: %v", err)
	}

	if !arguments.Options["--list"] {
		t.Fatal("--list option not parsed")
	}
	if arguments.Unnamed != nil {
		t.Fatalf("unexpected unnamed argument: %#v", arguments.Unnamed)
	}
}

func TestParseArgumentsRejectsCommandOptionAfterUnnamed(t *testing.T) {
	_, err := ParseArguments([]string{"codex", "--list"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseArgumentsRejectsCommandOptionAfterNamed(t *testing.T) {
	_, err := ParseArguments([]string{"dir=/tmp", "--list"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseArgumentsRejectsSecondUnnamed(t *testing.T) {
	_, err := ParseArguments([]string{"codex", "claude"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseArgumentsRejectsUnnamedAfterNamed(t *testing.T) {
	_, err := ParseArguments([]string{"dir=/tmp", "codex"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseArgumentsRejectsDuplicateOption(t *testing.T) {
	_, err := ParseArguments([]string{"--list", "--list"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseArgumentsRejectsUnknownGlobalSyntax(t *testing.T) {
	_, err := ParseArguments([]string{"-x", "codex"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseArgumentsRejectsMalformedCommandOption(t *testing.T) {
	for _, value := range []string{"--", "---list", "--list=true"} {
		if _, err := ParseArguments([]string{value}); err == nil {
			t.Fatalf("expected error for %q", value)
		}
	}
}
