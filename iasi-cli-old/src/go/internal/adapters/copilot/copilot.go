package copilot

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
	"iasi-cli/internal/resolver"
)

const schemaVersion = 1
const ownershipMarker = "IASI-GENERATED: copilot"
const manifestRelativePath = ".github/.iasi/copilot-manifest.yml"

type descriptor struct {
	SchemaVersion int                    `yaml:"schema_version"`
	ID            string                 `yaml:"id"`
	Platform      string                 `yaml:"platform"`
	Supports      map[string]bool        `yaml:"supports"`
	Instructions  map[string]instruction `yaml:"instructions"`
	Commands      commandMapping         `yaml:"commands"`
}

type instruction struct {
	Type    string `yaml:"type"`
	Target  string `yaml:"target"`
	ApplyTo string `yaml:"applyTo"`
}

type commandMapping struct {
	Type      string `yaml:"type"`
	Source    string `yaml:"source"`
	TargetDir string `yaml:"target_dir"`
	Suffix    string `yaml:"suffix"`
}

type metadata struct {
	ID     string `yaml:"id"`
	Status string `yaml:"status"`
	Scope  string `yaml:"scope"`
}

type candidate struct {
	ID, Scope, Body, Path string
	Status                string
}

type output struct {
	Path string
	Data []byte
}

type generatedManifest struct {
	SchemaVersion      int      `yaml:"schema_version"`
	Adapter            string   `yaml:"adapter"`
	ContextFingerprint string   `yaml:"context_fingerprint"`
	Generated          []string `yaml:"generated"`
}

func Run(project string) (string, error) {
	context, err := resolver.Resolve(project)
	if err != nil {
		return "", err
	}
	adapter, ok := context.Adapters["copilot"]
	if !ok {
		return "", errors.New("Copilot adapter is not available in this IASI installation")
	}
	desc, err := loadDescriptor(filepath.Join(adapter.Path, "adapter.yml"))
	if err != nil {
		return "", err
	}
	if !desc.Supports["instructions"] || !desc.Supports["commands"] {
		return "", errors.New("Copilot adapter does not support instructions")
	}

	candidates, err := discover(context.Instructions, desc)
	if err != nil {
		return "", err
	}
	previous, manifestExists, err := loadManifest(filepath.Join(project, manifestRelativePath))
	if err != nil {
		return "", err
	}

	outputs, err := buildOutputs(candidates, context.Commands, desc)
	if err != nil {
		return "", err
	}
	fingerprint := contextFingerprint(context, adapter)
	manifestData, err := marshalManifest(fingerprint, outputs)
	if err != nil {
		return "", err
	}
	manifestOutput := output{Path: manifestRelativePath, Data: manifestData}

	stale := []string{}
	if manifestExists {
		stale = stalePaths(previous, outputPaths(outputs))
	}
	if err := preflight(project, outputs, manifestOutput, stale); err != nil {
		return "", err
	}
	if err := commit(project, outputs, manifestOutput, stale); err != nil {
		return "", err
	}

	all := outputPaths(outputs)
	return formatSuccess("effective IASI context", project, all), nil
}

func loadDescriptor(path string) (descriptor, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return descriptor{}, errors.New("Copilot adapter is not available in this IASI installation")
		}
		return descriptor{}, fmt.Errorf("read Copilot adapter: %w", err)
	}
	var result descriptor
	if err := yaml.Unmarshal(data, &result); err != nil {
		return descriptor{}, fmt.Errorf("invalid Copilot adapter descriptor: %w", err)
	}
	if result.SchemaVersion != schemaVersion || result.ID != "copilot" || result.Platform != "github-copilot" {
		return descriptor{}, errors.New("invalid Copilot adapter descriptor")
	}
	if result.Supports["instructions"] && len(result.Instructions) == 0 {
		return descriptor{}, errors.New("Copilot adapter has no instruction mappings")
	}
	for scope, mapping := range result.Instructions {
		if mapping.Type != "repository" && mapping.Type != "path" {
			return descriptor{}, fmt.Errorf("invalid Copilot mapping type for scope: %s", scope)
		}
		if err := validateTarget(mapping.Target); err != nil {
			return descriptor{}, err
		}
		if mapping.Type == "path" && mapping.ApplyTo == "" {
			return descriptor{}, fmt.Errorf("missing applyTo for scope: %s", scope)
		}
	}
	if result.Supports["commands"] {
		if result.Commands.Type != "prompt" || result.Commands.Source != "commands" || result.Commands.TargetDir == "" || result.Commands.Suffix == "" {
			return descriptor{}, errors.New("invalid Copilot command mapping")
		}
		if err := validateTarget(filepath.ToSlash(filepath.Join(result.Commands.TargetDir, "command"+result.Commands.Suffix))); err != nil {
			return descriptor{}, err
		}
	}
	return result, nil
}

func discover(instructions map[string]resolver.Instruction, desc descriptor) ([]candidate, error) {
	var candidates []candidate
	for _, item := range instructions {
		if item.Status != "active" && item.Status != "draft" && item.Status != "deprecated" {
			return nil, fmt.Errorf("invalid instruction status %q: %s", item.Status, item.Path)
		}
		if item.Status == "active" {
			if _, ok := desc.Instructions[item.Scope]; !ok {
				return nil, fmt.Errorf("Unsupported instruction scope for Copilot adapter: %s", item.Scope)
			}
		}
		candidates = append(candidates, candidate{ID: item.ID, Scope: item.Scope, Status: item.Status, Body: item.Body, Path: item.Path})
	}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i].ID < candidates[j].ID })
	return candidates, nil
}

func parseInstruction(data []byte) (metadata, string, error) {
	text := strings.ReplaceAll(strings.TrimPrefix(string(data), "\ufeff"), "\r\n", "\n")
	if !strings.HasPrefix(text, "---\n") {
		return metadata{}, "", errors.New("missing YAML front matter")
	}
	end := strings.Index(text[4:], "\n---")
	if end < 0 {
		return metadata{}, "", errors.New("unterminated YAML front matter")
	}
	end += 4
	var result metadata
	if err := yaml.Unmarshal([]byte(text[4:end]), &result); err != nil {
		return metadata{}, "", fmt.Errorf("invalid YAML front matter: %w", err)
	}
	body := text[end+4:]
	body = strings.TrimPrefix(body, "\n")
	return result, body, nil
}

func buildOutputs(candidates []candidate, commands map[string]resolver.Command, desc descriptor) ([]output, error) {
	groups := map[string][]candidate{}
	for _, candidate := range candidates {
		if candidate.Status == "active" {
			groups[candidate.Scope] = append(groups[candidate.Scope], candidate)
		}
	}
	var outputs []output
	for scope, items := range groups {
		mapping := desc.Instructions[scope]
		sort.Slice(items, func(i, j int) bool { return items[i].ID < items[j].ID })
		var builder strings.Builder
		if mapping.Type == "path" {
			builder.WriteString("---\napplyTo: \"")
			builder.WriteString(mapping.ApplyTo)
			builder.WriteString("\"\n---\n\n")
		}
		builder.WriteString("<!-- IASI-GENERATED: copilot; schema=1 -->\nGenerated from IASI. Do not edit this file manually.\n\n")
		for _, item := range items {
			builder.WriteString("## IASI: ")
			builder.WriteString(item.ID)
			builder.WriteString("\n\n")
			builder.WriteString(item.Body)
			if !strings.HasSuffix(item.Body, "\n") {
				builder.WriteByte('\n')
			}
			builder.WriteByte('\n')
		}
		outputs = append(outputs, output{Path: mapping.Target, Data: []byte(builder.String())})
	}
	commandIDs := make([]string, 0, len(commands))
	for id := range commands {
		commandIDs = append(commandIDs, id)
	}
	sort.Strings(commandIDs)
	for _, id := range commandIDs {
		command := commands[id]
		target := filepath.ToSlash(filepath.Join(desc.Commands.TargetDir, id+desc.Commands.Suffix))
		if err := validateTarget(target); err != nil {
			return nil, err
		}
		content := "<!-- IASI-GENERATED: copilot; schema=1 -->\nGenerated from IASI. Do not edit this file manually.\n\n" + command.Content
		if !strings.HasSuffix(content, "\n") {
			content += "\n"
		}
		outputs = append(outputs, output{Path: target, Data: []byte(content)})
	}
	sort.Slice(outputs, func(i, j int) bool { return outputs[i].Path < outputs[j].Path })
	return outputs, nil
}

func loadManifest(path string) (generatedManifest, bool, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return generatedManifest{}, false, nil
	}
	if err != nil {
		return generatedManifest{}, false, err
	}
	if !strings.Contains(string(data), "# IASI-GENERATED") {
		return generatedManifest{}, true, errors.New("invalid Copilot manifest ownership")
	}
	var result generatedManifest
	if err := yaml.Unmarshal(data, &result); err != nil {
		return generatedManifest{}, true, fmt.Errorf("invalid Copilot manifest: %w", err)
	}
	if result.SchemaVersion != schemaVersion || result.Adapter != "copilot" || result.ContextFingerprint == "" {
		return generatedManifest{}, true, errors.New("invalid Copilot manifest")
	}
	for _, path := range result.Generated {
		if err := validateTarget(path); err != nil {
			return generatedManifest{}, true, fmt.Errorf("invalid Copilot manifest target: %w", err)
		}
	}
	return result, true, nil
}

func marshalManifest(fingerprint string, outputs []output) ([]byte, error) {
	paths := outputPaths(outputs)
	data, err := yaml.Marshal(generatedManifest{SchemaVersion: schemaVersion, Adapter: "copilot", ContextFingerprint: fingerprint, Generated: paths})
	if err != nil {
		return nil, err
	}
	return append([]byte("# IASI-GENERATED: copilot\n"), data...), nil
}

func contextFingerprint(context resolver.Context, adapter resolver.Adapter) string {
	var builder strings.Builder
	builder.Write(adapter.Content)
	instructionIDs := make([]string, 0, len(context.Instructions))
	for id := range context.Instructions {
		instructionIDs = append(instructionIDs, id)
	}
	sort.Strings(instructionIDs)
	for _, id := range instructionIDs {
		item := context.Instructions[id]
		fmt.Fprintf(&builder, "\ninstruction\x00%s\x00%s\x00%s\x00%s", item.ID, item.Status, item.Scope, item.Body)
	}
	commandIDs := make([]string, 0, len(context.Commands))
	for id := range context.Commands {
		commandIDs = append(commandIDs, id)
	}
	sort.Strings(commandIDs)
	for _, id := range commandIDs {
		fmt.Fprintf(&builder, "\ncommand\x00%s\x00%s", id, context.Commands[id].Content)
	}
	sum := sha256.Sum256([]byte(builder.String()))
	return fmt.Sprintf("%x", sum)
}

func preflight(project string, outputs []output, manifest output, stale []string) error {
	seen := map[string]bool{}
	for _, item := range append(outputs, manifest) {
		if seen[item.Path] {
			return fmt.Errorf("duplicate output target: %s", item.Path)
		}
		seen[item.Path] = true
		path := filepath.Join(project, filepath.FromSlash(item.Path))
		if data, err := os.ReadFile(path); err == nil && !strings.Contains(string(data), ownershipMarker) {
			return fmt.Errorf("Cannot generate Copilot instructions because this file already exists and is not managed by IASI: %s", item.Path)
		} else if err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	for _, relative := range stale {
		if seen[relative] {
			continue
		}
		path := filepath.Join(project, filepath.FromSlash(relative))
		data, err := os.ReadFile(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		if !strings.Contains(string(data), ownershipMarker) {
			return fmt.Errorf("stale Copilot output is not owned by IASI: %s", relative)
		}
	}
	return nil
}

func commit(project string, outputs []output, manifest output, stale []string) error {
	temp, err := os.MkdirTemp("", "iasi-copilot-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(temp)
	all := append(outputs, manifest)
	for _, item := range all {
		path := filepath.Join(temp, filepath.FromSlash(item.Path))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(path, item.Data, 0o644); err != nil {
			return err
		}
	}
	affected := map[string]string{}
	backupDir := filepath.Join(temp, "backup")
	if err := os.MkdirAll(backupDir, 0o755); err != nil {
		return err
	}
	for _, item := range all {
		realPath := filepath.Join(project, filepath.FromSlash(item.Path))
		if _, err := os.Stat(realPath); err == nil {
			backup := filepath.Join(backupDir, fmt.Sprintf("%d", len(affected)))
			if err := os.Rename(realPath, backup); err != nil {
				return rollback(affected, nil, err)
			}
			affected[realPath] = backup
		}
	}
	for _, relative := range stale {
		realPath := filepath.Join(project, filepath.FromSlash(relative))
		if _, err := os.Stat(realPath); err == nil {
			backup := filepath.Join(backupDir, fmt.Sprintf("%d", len(affected)))
			if err := os.Rename(realPath, backup); err != nil {
				return rollback(affected, nil, err)
			}
			affected[realPath] = backup
		}
	}
	created := []string{}
	for _, item := range all {
		realPath := filepath.Join(project, filepath.FromSlash(item.Path))
		if err := os.MkdirAll(filepath.Dir(realPath), 0o755); err != nil {
			return rollback(affected, created, err)
		}
		staged := filepath.Join(temp, filepath.FromSlash(item.Path))
		if err := os.Rename(staged, realPath); err != nil {
			return rollback(affected, created, err)
		}
		created = append(created, realPath)
	}
	return nil
}

func rollback(affected map[string]string, created []string, original error) error {
	for _, path := range created {
		_ = os.Remove(path)
	}
	for path, backup := range affected {
		_ = os.MkdirAll(filepath.Dir(path), 0o755)
		_ = os.Rename(backup, path)
	}
	return original
}

func stalePaths(previous generatedManifest, current []string) []string {
	known := map[string]bool{}
	for _, path := range current {
		known[path] = true
	}
	var stale []string
	for _, path := range previous.Generated {
		if !known[path] {
			stale = append(stale, path)
		}
	}
	return stale
}

func outputPaths(outputs []output) []string {
	paths := make([]string, 0, len(outputs))
	for _, item := range outputs {
		paths = append(paths, item.Path)
	}
	sort.Strings(paths)
	return paths
}

func validateTarget(target string) error {
	clean := filepath.ToSlash(filepath.Clean(target))
	if filepath.IsAbs(target) || clean == ".github" || !strings.HasPrefix(clean, ".github/") || strings.Contains(clean, "../") {
		return fmt.Errorf("invalid Copilot target: %s", target)
	}
	return nil
}

func formatSuccess(source, project string, paths []string) string {
	var builder strings.Builder
	builder.WriteString("IASI Copilot adapter\n\nSource : ")
	builder.WriteString(source)
	builder.WriteString("\nTarget : ")
	builder.WriteString(project)
	builder.WriteString("\n\nGenerated:\n")
	for _, path := range paths {
		builder.WriteString("  ")
		builder.WriteString(path)
		builder.WriteByte('\n')
	}
	return builder.String()
}
