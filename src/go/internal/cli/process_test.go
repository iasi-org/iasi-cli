package cli

import (
	"strconv"
	"testing"
)

type adapterArguments struct {
	Agent     string
	Dir       string
	Overwrite bool
}

func adapterTestSignature() Signature[adapterArguments] {
	return Signature[adapterArguments]{
		Parameters: []Parameter{
			{Name: "agent", Required: true},
			{Name: "dir"},
			{Name: "overwrite", Default: "false"},
		},
		Build: func(values map[string]string) (adapterArguments, error) {
			overwrite, err := strconv.ParseBool(values["overwrite"])
			if err != nil {
				return adapterArguments{}, err
			}
			return adapterArguments{Agent: values["agent"], Dir: values["dir"], Overwrite: overwrite}, nil
		},
	}
}

func TestProcessArgumentsMapsUnnamedToFirstParameter(t *testing.T) {
	arguments, err := ParseArguments([]string{"codex"})
	if err != nil {
		t.Fatalf("ParseArguments returned error: %v", err)
	}

	result, err := ProcessArguments(arguments, adapterTestSignature())
	if err != nil {
		t.Fatalf("ProcessArguments returned error: %v", err)
	}

	if result.Agent != "codex" {
		t.Fatalf("unexpected agent: %q", result.Agent)
	}
	if result.Dir != "" {
		t.Fatalf("dir must remain unresolved: %q", result.Dir)
	}
	if result.Overwrite {
		t.Fatal("overwrite should default to false")
	}
}

func TestProcessArgumentsAcceptsNamedFirstParameter(t *testing.T) {
	arguments, err := ParseArguments([]string{"agent=codex", "dir=/tmp", "overwrite=true"})
	if err != nil {
		t.Fatalf("ParseArguments returned error: %v", err)
	}

	result, err := ProcessArguments(arguments, adapterTestSignature())
	if err != nil {
		t.Fatalf("ProcessArguments returned error: %v", err)
	}

	if result.Agent != "codex" || result.Dir != "/tmp" || !result.Overwrite {
		t.Fatalf("unexpected result: %#v", result)
	}
}

func TestProcessArgumentsRejectsDuplicateFirstParameter(t *testing.T) {
	arguments, err := ParseArguments([]string{"codex", "agent=claude"})
	if err != nil {
		t.Fatalf("ParseArguments returned error: %v", err)
	}

	_, err = ProcessArguments(arguments, adapterTestSignature())
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestProcessArgumentsRejectsUnknownParameter(t *testing.T) {
	arguments, err := ParseArguments([]string{"codex", "unknown=value"})
	if err != nil {
		t.Fatalf("ParseArguments returned error: %v", err)
	}

	_, err = ProcessArguments(arguments, adapterTestSignature())
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestProcessArgumentsRequiresAgent(t *testing.T) {
	arguments, err := ParseArguments(nil)
	if err != nil {
		t.Fatalf("ParseArguments returned error: %v", err)
	}

	_, err = ProcessArguments(arguments, adapterTestSignature())
	if err == nil {
		t.Fatal("expected error")
	}
}
