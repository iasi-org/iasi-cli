package builtin

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/iasi-org/iasi-cli/internal/source"
)

func Copy(sourcePath string, destination string, overwrite bool) error {
	return copyFS(source.Builtin(), sourcePath, destination, overwrite)
}

func copyFS(sourceFS fs.FS, sourcePath string, destination string, overwrite bool) error {
	return fs.WalkDir(sourceFS, sourcePath, func(current string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relative := strings.TrimPrefix(current, sourcePath)
		relative = strings.TrimPrefix(relative, "/")

		target := destination
		if relative != "" {
			target = filepath.Join(destination, filepath.FromSlash(relative))
		}

		if entry.IsDir() {
			if err := os.MkdirAll(target, 0755); err != nil {
				return fmt.Errorf("create destination directory %q: %w", target, err)
			}
			return nil
		}

		if err := copyFile(sourceFS, current, target, overwrite); err != nil {
			return fmt.Errorf("copy builtin file %q: %w", current, err)
		}
		return nil
	})
}

func copyFile(sourceFS fs.FS, sourcePath string, destination string, overwrite bool) error {
	sourceFile, err := sourceFS.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	flags := os.O_WRONLY | os.O_CREATE
	if overwrite {
		flags |= os.O_TRUNC
	} else {
		flags |= os.O_EXCL
	}

	destinationFile, err := os.OpenFile(destination, flags, 0644)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}
