package install

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"

	"iasi-cli/internal/manifest"
)

var categories = []string{"instructions", "commands", "skills", "mcp", "adapters"}

func Run(workspace string, methodology fs.FS, version string) (string, error) {
	return run(workspace, methodology, version, false)
}

func Reinstall(workspace string, methodology fs.FS, version string) (string, error) {
	return run(workspace, methodology, version, true)
}

func run(workspace string, methodology fs.FS, version string, replace bool) (string, error) {
	target := filepath.Join(workspace, ".iasi")
	hasManifest, validationState, err := inspectTarget(target)
	if err != nil {
		return "", err
	}
	if replace && !hasManifest {
		return "", fmt.Errorf("IASI is not installed in this directory: %s", target)
	}
	if !replace && hasManifest {
		return "", fmt.Errorf("IASI is already installed in this directory: %s", target)
	}

	temporary, err := os.MkdirTemp(workspace, ".iasi.tmp-")
	if err != nil {
		return "", fmt.Errorf("create installation directory: %w", err)
	}
	defer os.RemoveAll(temporary)

	for _, category := range categories {
		destination := filepath.Join(temporary, category)
		if err := os.MkdirAll(destination, 0o755); err != nil {
			return "", fmt.Errorf("create %s directory: %w", category, err)
		}
		sourceDirectory := filepath.ToSlash(filepath.Join("iasi", category))
		if err := copyDirectory(methodology, sourceDirectory, destination); err != nil {
			return "", fmt.Errorf("copy %s: %w", category, err)
		}
	}
	if err := manifest.Write(filepath.Join(temporary, "manifest.yml"), version); err != nil {
		return "", err
	}
	if validationState != nil {
		if err := os.WriteFile(filepath.Join(temporary, "validation.json"), validationState, 0o644); err != nil {
			return "", fmt.Errorf("preserve validation state: %w", err)
		}
	}
	if err := replaceTarget(target, temporary); err != nil {
		return "", err
	}
	return target, nil
}

func inspectTarget(target string) (bool, []byte, error) {
	entries, err := os.ReadDir(target)
	if os.IsNotExist(err) {
		return false, nil, nil
	}
	if err != nil {
		return false, nil, fmt.Errorf("inspect installation path: %w", err)
	}
	_, err = os.Stat(filepath.Join(target, "manifest.yml"))
	hasManifest := err == nil
	if err != nil && !os.IsNotExist(err) {
		return false, nil, fmt.Errorf("inspect installation manifest: %w", err)
	}
	if hasManifest {
		if _, err := manifest.ReadVersion(filepath.Join(target, "manifest.yml")); err != nil {
			return false, nil, fmt.Errorf("invalid existing installation: %w", err)
		}
	}
	allowed := map[string]bool{"manifest.yml": hasManifest, "validation.json": true}
	if hasManifest {
		for _, category := range categories {
			allowed[category] = true
		}
	}
	for _, entry := range entries {
		if !allowed[entry.Name()] {
			return false, nil, fmt.Errorf("ambiguous existing IASI content: %s", filepath.Join(target, entry.Name()))
		}
	}
	var validationState []byte
	if data, err := os.ReadFile(filepath.Join(target, "validation.json")); err == nil {
		validationState = data
	} else if !os.IsNotExist(err) {
		return false, nil, fmt.Errorf("read validation state: %w", err)
	}
	return hasManifest, validationState, nil
}

func replaceTarget(target, temporary string) error {
	if _, err := os.Stat(target); os.IsNotExist(err) {
		if err := os.Rename(temporary, target); err != nil {
			return fmt.Errorf("complete installation: %w", err)
		}
		return nil
	} else if err != nil {
		return fmt.Errorf("inspect installation target: %w", err)
	}
	backup := target + ".backup"
	if err := os.RemoveAll(backup); err != nil {
		return fmt.Errorf("clear installation backup: %w", err)
	}
	if err := os.Rename(target, backup); err != nil {
		return fmt.Errorf("stage existing installation: %w", err)
	}
	if err := os.Rename(temporary, target); err != nil {
		_ = os.Rename(backup, target)
		return fmt.Errorf("complete installation: %w", err)
	}
	if err := os.RemoveAll(backup); err != nil {
		return fmt.Errorf("remove previous installation backup: %w", err)
	}
	return nil
}

func copyDirectory(source fs.FS, directory, destination string) error {
	if _, err := fs.Stat(source, directory); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	var paths []string
	if err := fs.WalkDir(source, directory, func(path string, info fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		paths = append(paths, path)
		return nil
	}); err != nil {
		return err
	}
	sort.Strings(paths)
	for _, path := range paths {
		info, err := fs.Stat(source, path)
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(directory, path)
		if err != nil {
			return err
		}
		if relative == "." {
			continue
		}
		target := filepath.Join(destination, relative)
		if info.IsDir() {
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
			continue
		}
		data, err := fs.ReadFile(source, path)
		if err != nil {
			return err
		}
		if err := os.WriteFile(target, data, 0o644); err != nil {
			return err
		}
	}
	return nil
}
