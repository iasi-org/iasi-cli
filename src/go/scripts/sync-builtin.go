package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		fatal(fmt.Errorf("cannot resolve build script location"))
	}

	goRoot := filepath.Dir(filepath.Dir(file))
	cliRoot := filepath.Dir(filepath.Dir(goRoot))
	source := filepath.Join(filepath.Dir(cliRoot), "iasi-core", "builtin")
	destination := filepath.Join(goRoot, "internal", "source", "embedded", "builtin")

	if len(os.Args) == 2 && os.Args[1] == "--clean" {
		if err := os.RemoveAll(destination); err != nil {
			fatal(fmt.Errorf("remove generated builtin %q: %w", destination, err))
		}
		return
	}

	if len(os.Args) != 1 {
		fatal(fmt.Errorf("usage: sync-builtin.go [--clean]"))
	}

	if err := os.RemoveAll(destination); err != nil {
		fatal(fmt.Errorf("remove generated builtin %q: %w", destination, err))
	}

	info, err := os.Stat(source)
	if err != nil {
		if os.IsNotExist(err) {
			fatal(fmt.Errorf("IASI core builtin not found at %q", source))
		}
		fatal(fmt.Errorf("access IASI core builtin %q: %w", source, err))
	}
	if !info.IsDir() {
		fatal(fmt.Errorf("IASI core builtin path %q is not a directory", source))
	}

	if err := copyTree(source, destination); err != nil {
		fatal(err)
	}
}

func copyTree(source string, destination string) error {
	return filepath.WalkDir(source, func(current string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relative, err := filepath.Rel(source, current)
		if err != nil {
			return err
		}

		target := destination
		if relative != "." {
			target = filepath.Join(destination, relative)
		}

		if entry.IsDir() {
			if err := os.MkdirAll(target, 0755); err != nil {
				return fmt.Errorf("create generated builtin directory %q: %w", target, err)
			}
			return nil
		}

		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("builtin contains unsupported symbolic link %q", current)
		}

		return copyFile(current, target)
	})
}

func copyFile(source string, destination string) error {
	input, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("open builtin source file %q: %w", source, err)
	}
	defer input.Close()

	output, err := os.OpenFile(destination, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("create generated builtin file %q: %w", destination, err)
	}

	if _, err := io.Copy(output, input); err != nil {
		output.Close()
		return fmt.Errorf("copy builtin file %q: %w", source, err)
	}

	if err := output.Close(); err != nil {
		return fmt.Errorf("close generated builtin file %q: %w", destination, err)
	}

	return nil
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
