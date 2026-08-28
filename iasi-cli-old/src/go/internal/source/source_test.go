package source

import (
	"io/fs"
	"testing"
)

func TestMethodologyIsEmbedded(t *testing.T) {
	if _, err := fs.ReadFile(Methodology(), "iasi/instructions/code/style.md"); err != nil {
		t.Fatalf("expected methodology to be embedded: %v", err)
	}
	if _, err := fs.ReadFile(Methodology(), "iasi/adapters/copilot/adapter.yml"); err != nil {
		t.Fatalf("expected adapter to be embedded: %v", err)
	}
}

func TestVersionIsEmbedded(t *testing.T) {
	version, err := Version()
	if err != nil {
		t.Fatal(err)
	}
	if version != "0.1.0" {
		t.Fatalf("expected embedded version 0.1.0, got %s", version)
	}
}
