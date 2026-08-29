// Command iasi is the IASI command-line entry point.
package main

import (
	"os"
	"sort"

	"github.com/iasi-org/iasi-cli/internal/cli"
)

var version = "dev"

var commands = map[string]func([]string){
	"adapter": runAdapter,
}

func main() {
	options, args := cli.ParseOutputOptions(os.Args[1:])
	cli.ConfigureOutput(options)

	if len(args) == 0 {
		printHelp()
		return
	}

	if args[0] == "--help" || args[0] == "--version" || args[0] == "--list" {
		if len(args) != 1 {
			cli.Error("option %q does not accept arguments", args[0])
			os.Exit(2)
		}
		runOption(args[0])
		return
	}

	runCommand(args[0], args[1:])
}

func runOption(option string) {
	switch option {
	case "--help":
		printHelp()
	case "--version":
		cli.Direct("IASI CLI %s\n", version)
	case "--list":
		listCommands()
	}
}

func runCommand(name string, args []string) {
	command, exists := commands[name]
	if !exists {
		cli.Error("unknown command %q", name)
		os.Exit(2)
	}
	command(args)
}

func printHelp() {
	cli.Direct(`IASI CLI

Usage:
  iasi <command> [flags] [options] [unnamed] [named-parameters] | <options>

Flags: [-s][-v][-d]

Options:
  --list        List commands
  --help        Show help
  --version     Show version
`)
}

func listCommands() {
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		cli.Direct("%s\n", name)
	}
}

func fatal(err error) {
	cli.Error("%v", err)
	os.Exit(2)
}
