package manifest

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type document struct {
	SchemaVersion int    `yaml:"schema_version"`
	Version       string `yaml:"version"`
}

func Write(path, version string) error {
	content := fmt.Sprintf("schema_version: 1\nversion: %s\n\ninstalled:\n  instructions: all\n  commands: all\n  skills: all\n  mcp: all\n  adapters: all\n", version)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write manifest: %w", err)
	}
	return nil
}

func ReadVersion(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read manifest: %w", err)
	}
	var result document
	if err := yaml.Unmarshal(data, &result); err != nil {
		return "", fmt.Errorf("parse manifest: %w", err)
	}
	if result.SchemaVersion != 1 {
		return "", fmt.Errorf("unsupported manifest schema version")
	}
	if result.Version == "" {
		return "", fmt.Errorf("manifest has no version")
	}
	return result.Version, nil
}
