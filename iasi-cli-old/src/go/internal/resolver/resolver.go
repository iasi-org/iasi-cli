package resolver

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
	"iasi-cli/internal/manifest"
)

var ErrNotInstalled = errors.New("IASI is not installed for this location")

type Layer struct {
	Path    string
	Version string
}

type Instruction struct {
	ID, Status, Scope, Path, Body string
}

type Command struct {
	ID, Path, Content string
}

type Adapter struct {
	ID, Path string
	Content  []byte
}

type Context struct {
	Layers       []Layer
	Instructions map[string]Instruction
	Commands     map[string]Command
	Adapters     map[string]Adapter
	Skills       map[string]string
	MCP          map[string]string
}

type instructionMetadata struct {
	ID     string `yaml:"id"`
	Status string `yaml:"status"`
	Scope  string `yaml:"scope"`
}

type adapterMetadata struct {
	SchemaVersion int    `yaml:"schema_version"`
	ID            string `yaml:"id"`
	Platform      string `yaml:"platform"`
}

func Resolve(start string) (Context, error) {
	layers, err := FindLayers(start)
	if err != nil {
		return Context{}, err
	}
	if len(layers) == 0 {
		return Context{}, ErrNotInstalled
	}
	context := Context{
		Layers:       layers,
		Instructions: map[string]Instruction{},
		Commands:     map[string]Command{},
		Adapters:     map[string]Adapter{},
		Skills:       map[string]string{},
		MCP:          map[string]string{},
	}
	for _, layer := range layers {
		if err := resolveInstructions(layer, context.Instructions); err != nil {
			return Context{}, err
		}
		if err := resolveCommands(layer, context.Commands); err != nil {
			return Context{}, err
		}
		if err := resolveAdapters(layer, context.Adapters); err != nil {
			return Context{}, err
		}
		if err := resolveSubtrees(filepath.Join(layer.Path, "skills"), context.Skills); err != nil {
			return Context{}, err
		}
		if err := resolveSubtrees(filepath.Join(layer.Path, "mcp"), context.MCP); err != nil {
			return Context{}, err
		}
	}
	return context, nil
}

func FindLayers(start string) ([]Layer, error) {
	var reverse []Layer
	for current := filepath.Clean(start); ; current = filepath.Dir(current) {
		manifestPath := filepath.Join(current, ".iasi", "manifest.yml")
		if _, err := os.Stat(manifestPath); err == nil {
			version, err := manifest.ReadVersion(manifestPath)
			if err != nil {
				return nil, fmt.Errorf("invalid IASI layer %s: %w", filepath.Join(current, ".iasi"), err)
			}
			reverse = append(reverse, Layer{Path: filepath.Join(current, ".iasi"), Version: version})
		} else if !os.IsNotExist(err) {
			return nil, fmt.Errorf("inspect IASI layer: %w", err)
		}
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
	}
	layers := make([]Layer, len(reverse))
	for index, layer := range reverse {
		layers[len(reverse)-1-index] = layer
	}
	return layers, nil
}

func resolveInstructions(layer Layer, effective map[string]Instruction) error {
	root := filepath.Join(layer.Path, "instructions")
	paths, err := markdownFiles(root)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	seen := map[string]bool{}
	for _, path := range paths {
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		if isSupportPath(relative) {
			continue
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		metadata, body, err := parseInstruction(string(content))
		if err != nil {
			return fmt.Errorf("invalid instruction %s: %w", path, err)
		}
		if seen[metadata.ID] {
			return fmt.Errorf("duplicate instruction ID %q in %s", metadata.ID, layer.Path)
		}
		seen[metadata.ID] = true
		effective[metadata.ID] = Instruction{ID: metadata.ID, Status: metadata.Status, Scope: metadata.Scope, Path: filepath.ToSlash(relative), Body: body}
	}
	return nil
}

func resolveCommands(layer Layer, effective map[string]Command) error {
	root := filepath.Join(layer.Path, "commands")
	paths, err := markdownFiles(root)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	for _, path := range paths {
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		if isSupportPath(relative) {
			continue
		}
		identity := strings.TrimSuffix(filepath.ToSlash(relative), ".md")
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		effective[identity] = Command{ID: identity, Path: filepath.ToSlash(relative), Content: string(content)}
	}
	return nil
}

func resolveAdapters(layer Layer, effective map[string]Adapter) error {
	root := filepath.Join(layer.Path, "adapters")
	entries, err := os.ReadDir(root)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	sort.Slice(entries, func(left, right int) bool { return entries[left].Name() < entries[right].Name() })
	for _, entry := range entries {
		if !entry.IsDir() || entry.Name() == "schema" {
			continue
		}
		adapterPath := filepath.Join(root, entry.Name())
		descriptorPath := filepath.Join(adapterPath, "adapter.yml")
		content, err := os.ReadFile(descriptorPath)
		if err != nil {
			return fmt.Errorf("read adapter %s: %w", adapterPath, err)
		}
		var metadata adapterMetadata
		if err := yaml.Unmarshal(content, &metadata); err != nil {
			return fmt.Errorf("invalid adapter %s: %w", adapterPath, err)
		}
		if metadata.SchemaVersion != 1 || metadata.ID == "" || metadata.Platform == "" || metadata.ID != entry.Name() {
			return fmt.Errorf("invalid adapter %s", adapterPath)
		}
		effective[metadata.ID] = Adapter{ID: metadata.ID, Path: adapterPath, Content: content}
	}
	return nil
}

func resolveSubtrees(root string, effective map[string]string) error {
	entries, err := os.ReadDir(root)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	sort.Slice(entries, func(left, right int) bool { return entries[left].Name() < entries[right].Name() })
	for _, entry := range entries {
		if entry.IsDir() {
			effective[entry.Name()] = filepath.Join(root, entry.Name())
		}
	}
	return nil
}

func markdownFiles(root string) ([]string, error) {
	var paths []string
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !entry.IsDir() && strings.EqualFold(filepath.Ext(entry.Name()), ".md") {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(paths)
	return paths, nil
}

func isSupportPath(relative string) bool {
	parts := strings.Split(filepath.ToSlash(relative), "/")
	if len(parts) > 1 && parts[0] == "schema" {
		return true
	}
	name := parts[len(parts)-1]
	return name == "README.md" || strings.HasPrefix(name, "README_")
}

func parseInstruction(content string) (instructionMetadata, string, error) {
	content = strings.ReplaceAll(strings.TrimPrefix(content, "\ufeff"), "\r\n", "\n")
	if !strings.HasPrefix(content, "---\n") {
		return instructionMetadata{}, "", fmt.Errorf("missing YAML front matter")
	}
	remainder := strings.TrimPrefix(content, "---\n")
	boundary := strings.Index(remainder, "\n---\n")
	if boundary < 0 {
		return instructionMetadata{}, "", fmt.Errorf("unterminated YAML front matter")
	}
	var metadata instructionMetadata
	if err := yaml.Unmarshal([]byte(remainder[:boundary]), &metadata); err != nil {
		return instructionMetadata{}, "", err
	}
	if metadata.ID == "" || metadata.Status == "" || metadata.Scope == "" {
		return instructionMetadata{}, "", fmt.Errorf("missing id, status, or scope")
	}
	return metadata, remainder[boundary+len("\n---\n"):], nil
}
