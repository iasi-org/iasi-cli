package validation

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"iasi-cli/internal/resolver"
)

const stateRelativePath = ".iasi/validation.json"

type State struct {
	SchemaVersion    int    `json:"schema_version"`
	Status           string `json:"status"`
	ValidatedAt      string `json:"validated_at"`
	InstructionsHash string `json:"instructions_hash"`
	InputsHash       string `json:"inputs_hash"`
	CommandHash      string `json:"command_hash"`
	Blockers         int    `json:"blockers"`
	Warnings         int    `json:"warnings"`
}

func NewState(project string, context resolver.Context, status string, blockers, warnings int) (State, error) {
	if status != "passed" && status != "failed" {
		return State{}, fmt.Errorf("invalid validation status %q", status)
	}
	instructionsHash := HashInstructions(context)
	inputsHash, err := HashInputs(project)
	if err != nil {
		return State{}, err
	}
	commandHash, err := HashValidateCommand(context)
	if err != nil {
		return State{}, err
	}
	return State{
		SchemaVersion:    1,
		Status:           status,
		ValidatedAt:      time.Now().UTC().Format(time.RFC3339),
		InstructionsHash: instructionsHash,
		InputsHash:       inputsHash,
		CommandHash:      commandHash,
		Blockers:         blockers,
		Warnings:         warnings,
	}, nil
}

func Write(project string, state State) error {
	if err := validateState(state); err != nil {
		return err
	}
	path := filepath.Join(project, filepath.FromSlash(stateRelativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create validation state directory: %w", err)
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, append(data, '\n'), 0o644); err != nil {
		return fmt.Errorf("write validation state: %w", err)
	}
	return nil
}

func Read(project string) (State, error) {
	data, err := os.ReadFile(filepath.Join(project, filepath.FromSlash(stateRelativePath)))
	if err != nil {
		return State{}, err
	}
	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return State{}, fmt.Errorf("parse validation state: %w", err)
	}
	if err := validateState(state); err != nil {
		return State{}, err
	}
	return state, nil
}

func RequireCurrent(project string, context resolver.Context) error {
	state, err := Read(project)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("project has no validation state")
		}
		return err
	}
	if state.Status != "passed" {
		return errors.New("project validation has failed")
	}
	inputsHash, err := HashInputs(project)
	if err != nil {
		return err
	}
	commandHash, err := HashValidateCommand(context)
	if err != nil {
		return err
	}
	if state.InstructionsHash != HashInstructions(context) || state.InputsHash != inputsHash || state.CommandHash != commandHash {
		return errors.New("project validation is stale")
	}
	return nil
}

func HashInstructions(context resolver.Context) string {
	ids := make([]string, 0, len(context.Instructions))
	for id := range context.Instructions {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	var builder strings.Builder
	for _, id := range ids {
		item := context.Instructions[id]
		fmt.Fprintf(&builder, "%s\x00%s\x00%s\x00%s\n", item.ID, item.Status, item.Scope, item.Body)
	}
	return hash(builder.String())
}

func HashInputs(project string) (string, error) {
	root := filepath.Join(project, "inputs")
	return HashInputsAt(root,
		filepath.Join(root, "externals", "archived"),
		filepath.Join(root, "internals", "archived"),
		filepath.Join(root, "obtained", "archived"),
	)
}

func HashInputsAt(root string, excluded ...string) (string, error) {
	skipped := map[string]bool{}
	for _, path := range excluded {
		skipped[filepath.Clean(path)] = true
	}
	var paths []string
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() && skipped[filepath.Clean(path)] {
			return filepath.SkipDir
		}
		if !entry.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if os.IsNotExist(err) {
		return hash(""), nil
	}
	if err != nil {
		return "", err
	}
	sort.Strings(paths)
	var builder strings.Builder
	for _, path := range paths {
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return "", err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&builder, "%s\x00%s\n", filepath.ToSlash(relative), data)
	}
	return hash(builder.String()), nil
}

func HashValidateCommand(context resolver.Context) (string, error) {
	command, ok := context.Commands["validate"]
	if !ok {
		return "", errors.New("effective validate command is unavailable")
	}
	return hash(command.Content), nil
}

func validateState(state State) error {
	if state.SchemaVersion != 1 || (state.Status != "passed" && state.Status != "failed") || state.InstructionsHash == "" || state.InputsHash == "" || state.CommandHash == "" || state.Blockers < 0 || state.Warnings < 0 {
		return errors.New("invalid validation state")
	}
	return nil
}

func hash(value string) string {
	sum := sha256.Sum256([]byte(value))
	return hex.EncodeToString(sum[:])
}
