package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strconv"

	"github.com/iasi-org/iasi-cli/internal/builtin"
	"github.com/iasi-org/iasi-cli/internal/cli"
	"github.com/iasi-org/iasi-cli/internal/source"
)

type Adapter struct {
	Agent     string
	Dir       string
	Overwrite bool
	List      bool
}

var adapterSignature = cli.Signature[Adapter]{
	Parameters: []cli.Parameter{
		{Name: "agent"},
		{Name: "dir"},
		{Name: "overwrite", Default: "false"},
		{Name: "list", Option: "--list"},
	},
	Build: buildAdapter,
}

func runAdapter(args []string) {
	cli.Debug("Adapter raw arguments: %v", args)

	arguments, err := cli.ParseArguments(args)
	if err != nil {
		fatal(err)
	}

	adapter, err := cli.ProcessArguments(arguments, adapterSignature)
	if err != nil {
		fatal(err)
	}

	if err := validateAdapter(adapter); err != nil {
		fatal(err)
	}

	if adapter.List {
		if err := listAdapters(); err != nil {
			fatal(err)
		}
		return
	}

	resolveAdapterDefaults(&adapter)
	cli.Debug("Adapter resolved: agent=%q dir=%q overwrite=%t", adapter.Agent, adapter.Dir, adapter.Overwrite)

	cli.Verbose("Checking adapter %q availability", adapter.Agent)
	if err := checkAdapterAvailable(adapter); err != nil {
		fatal(err)
	}
	cli.Verbose("Preparing destination %q", adapter.Dir)
	if err := prepareDestination(adapter); err != nil {
		fatal(err)
	}
	cli.Verbose("Copying adapter %q to %q", adapter.Agent, adapter.Dir)
	if err := builtin.Copy("adapters/"+adapter.Agent, adapter.Dir, adapter.Overwrite); err != nil {
		fatal(err)
	}

	cli.Success("Adapter %q installed in %q", adapter.Agent, adapter.Dir)
}

func buildAdapter(values map[string]string) (Adapter, error) {
	overwrite, err := strconv.ParseBool(values["overwrite"])
	if err != nil {
		return Adapter{}, fmt.Errorf("invalid overwrite value %q", values["overwrite"])
	}

	list, err := strconv.ParseBool(values["list"])
	if err != nil {
		return Adapter{}, fmt.Errorf("invalid list value %q", values["list"])
	}

	return Adapter{Agent: values["agent"], Dir: values["dir"], Overwrite: overwrite, List: list}, nil
}

func validateAdapter(adapter Adapter) error {
	if adapter.List {
		if adapter.Agent != "" {
			return fmt.Errorf("adapter cannot be specified with --list")
		}
		if adapter.Dir != "" {
			return fmt.Errorf("dir cannot be specified with --list")
		}
		if adapter.Overwrite {
			return fmt.Errorf("overwrite cannot be specified with --list")
		}
		return nil
	}

	if adapter.Agent == "" {
		return fmt.Errorf("argument %q is required", "agent")
	}
	return nil
}

func listAdapters() error {
	entries, err := fs.ReadDir(source.Builtin(), "adapters")
	if err != nil {
		return fmt.Errorf("list builtin adapters: %w", err)
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			names = append(names, entry.Name())
		}
	}
	sort.Strings(names)

	for _, name := range names {
		cli.Info("%s", name)
	}
	return nil
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
