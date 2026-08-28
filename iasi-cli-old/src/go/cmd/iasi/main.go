package main

import (
	"errors"
	"fmt"
	"os"

	"iasi-cli/internal/adapters/copilot"
	"iasi-cli/internal/install"
	"iasi-cli/internal/runtime"
	"iasi-cli/internal/source"
	"iasi-cli/internal/status"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if runtimeError, ok := err.(*runtime.Error); ok {
			os.Exit(runtimeError.Code)
		}
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("usage: iasi version | iasi install | iasi reinstall | iasi status | iasi adapt copilot")
	}

	switch args[0] {
	case "__runtime":
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		output, err := runtime.Run(cwd, args[1:])
		if err != nil {
			return err
		}
		fmt.Print(output)
		return nil

	case "adapt":
		if len(args) != 2 || args[1] != "copilot" {
			return errors.New("usage: iasi adapt copilot")
		}
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		output, err := copilot.Run(cwd)
		if err != nil {
			return err
		}
		fmt.Print(output)
		return nil

	case "version":
		if len(args) != 1 {
			return errors.New("usage: iasi version")
		}
		version, err := source.Version()
		if err != nil {
			return err
		}
		fmt.Printf("IASI %s\n", version)
		return nil

	case "install", "reinstall":
		command := args[0]
		if len(args) != 1 {
			return errors.New("usage: iasi " + command)
		}

		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		version, err := source.Version()
		if err != nil {
			return err
		}
		if command == "reinstall" {
			_, err = install.Reinstall(cwd, source.Methodology(), version)
		} else {
			_, err = install.Run(cwd, source.Methodology(), version)
		}
		return err

	case "status":
		if len(args) != 1 {
			return errors.New("usage: iasi status")
		}
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}
		result, err := status.Find(cwd)
		if err != nil {
			return err
		}
		version, err := source.Version()
		if err != nil {
			return err
		}
		fmt.Print(status.Format(result, version))
		return nil
	}

	return fmt.Errorf("unknown command %q", args[0])
}
