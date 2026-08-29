package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"

	"github.com/iasi-org/iasi-cli/internal/builtin"
	"github.com/iasi-org/iasi-cli/internal/cli"
	"github.com/iasi-org/iasi-cli/internal/source"
)

type Adapter struct {
	Agent     string
	Dir       string
	Overwrite bool
}

var adapterSignature = cli.Signature[Adapter]{
	Parameters: []cli.Parameter{
		{Name: "agent", Required: true},
		{Name: "dir"},
		{Name: "overwrite", Default: "false"},
	},
	Build: buildAdapter,
}

func runAdapter(args []string) {
	arguments, err := cli.ParseArguments(args)
	if err != nil {
		fatal(err)
	}

	adapter, err := cli.ProcessArguments(arguments, adapterSignature)
	if err != nil {
		fatal(err)
	}

	resolveAdapterDefaults(&adapter)

	if err := checkAdapterAvailable(adapter); err != nil {
		fatal(err)
	}
	if err := prepareDestination(adapter); err != nil {
		fatal(err)
	}
	if err := builtin.Copy("adapters/"+adapter.Agent, adapter.Dir, adapter.Overwrite); err != nil {
		fatal(err)
	}
}

func buildAdapter(values map[string]string) (Adapter, error) {
	overwrite, err := strconv.ParseBool(values["overwrite"])
	if err != nil {
		return Adapter{}, fmt.Errorf("invalid overwrite value %q", values["overwrite"])
	}
	return Adapter{Agent: values["agent"], Dir: values["dir"], Overwrite: overwrite}, nil
}

func resolveAdapterDefaults(adapter *Adapter) {
	if adapter.Dir == "" {
		adapter.Dir = "." + adapter.Agent
	}
}

func checkAdapterAvailable(adapter Adapter) error {
	info, err := fs.Stat(source.Builtin(), "adapters/"+adapter.Agent)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("adapter %q is not available", adapter.Agent)
		}
		return fmt.Errorf("check adapter %q: %w", adapter.Agent, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("adapter %q is not available", adapter.Agent)
	}
	return nil
}

func prepareDestination(adapter Adapter) error {
	info, err := os.Stat(adapter.Dir)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(adapter.Dir, 0755); err != nil {
			return fmt.Errorf("create destination %q: %w", adapter.Dir, err)
		}
		return nil
	}
	if err != nil {
		return fmt.Errorf("access destination %q: %w", adapter.Dir, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("destination %q is not a directory", adapter.Dir)
	}
	if !adapter.Overwrite {
		return fmt.Errorf("destination %q already exists", adapter.Dir)
	}
	return nil
}
