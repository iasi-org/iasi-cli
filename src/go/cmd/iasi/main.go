// Command iasi is the IASI command-line entry point.
package main

import (
	"fmt"
	"os"
)

var version = "dev"

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	switch args[0] {
	case "version":
		fmt.Printf("IASI CLI %s\n", version)
		return

	case "adapter":
		runAdapter(args[1:])
		return

	default:
		fmt.Fprintf(os.Stderr, "iasi: unknown command %q\n", args[0])
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
	fmt.Fprintf(os.Stderr, "iasi: %v\n", err)
	os.Exit(2)
}
