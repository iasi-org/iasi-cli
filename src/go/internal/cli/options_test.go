package cli

import (
	"reflect"
	"testing"
)

func TestParseOutputOptions(t *testing.T) {
	options, args := ParseOutputOptions([]string{"adapter", "codex", "-v", "dir=tools", "-d"})

	if options.Silent {
		t.Fatal("Silent = true, want false")
	}
	if !options.Verbose {
		t.Fatal("Verbose = false, want true")
	}
	if !options.Debug {
		t.Fatal("Debug = false, want true")
	}

	want := []string{"adapter", "codex", "dir=tools"}
	if !reflect.DeepEqual(args, want) {
		t.Fatalf("args = %#v, want %#v", args, want)
	}
}

func TestParseOutputOptionsSilent(t *testing.T) {
	options, args := ParseOutputOptions([]string{"-s", "adapter", "codex"})

	if !options.Silent {
		t.Fatal("Silent = false, want true")
	}
	if !reflect.DeepEqual(args, []string{"adapter", "codex"}) {
		t.Fatalf("args = %#v", args)
	}
}
