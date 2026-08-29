// Command iasi is the IASI command-line entry point.
package main

import (
	"fmt"
	"os"

	"github.com/iasi-org/iasi-cli/internal/cli"
)

var version = "dev"

func main() {
	options, args := cli.ParseOutputOptions(os.Args[1:])
	cli.ConfigureOutput(options)

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	switch args[0] {
	case "version":
		cli.Info("IASI CLI %s", version)
		return

	case "adapter":
		runAdapter(args[1:])
		return

	default:
		cli.Error("unknown command %q", args[0])
		os.Exit(2)
	}
}

func printHelp() {
	fmt.Print(`IASI CLI

Usage:
  iasi <command> <flags> <options>

Commands:
  version    Show CLI version

Use "iasi <command> --help" for more information about a command.
`)
}

func fatal(err error) {
	cli.Error("%v", err)
	os.Exit(2)
}
