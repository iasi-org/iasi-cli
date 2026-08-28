package workflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"iasi-cli/internal/resolver"
	"iasi-cli/internal/validation"
)

const statePath = ".iasi/workflow.json"

const (
	Inputs          = "INPUTS"
	InputsValidated = "INPUTS_VALIDATED"
	Planned         = "PLANNED"
	PlanValidated   = "PLAN_VALIDATED"
	Executed        = "EXECUTED"
	Verified        = "VERIFIED"
)

type State struct {
	Version          int     `json:"version"`
	Checkpoint       string  `json:"checkpoint"`
	LastCommand      string  `json:"last_command"`
	LastResult       string  `json:"last_result"`
	FailedCommand    *string `json:"failed_command"`
	InputsHash       string  `json:"inputs_hash"`
	InstructionsHash string  `json:"instructions_hash"`
	PlanHash         *string `json:"plan_hash"`
	UpdatedAt        string  `json:"updated_at"`
}

func Current(project string, context resolver.Context) (State, error) {
	state, err := Load(project)
	if os.IsNotExist(err) {
		return initial(project, context)
	}
	if err != nil {
		return State{}, err
	}
	return refresh(project, context, state)
}

func Load(project string) (State, error) {
	data, err := os.ReadFile(filepath.Join(project, filepath.FromSlash(statePath)))
	if err != nil {
		return State{}, err
	}
	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return State{}, fmt.Errorf("parse workflow state: %w", err)
	}
	if state.Version != 1 || !validCheckpoint(state.Checkpoint) || (state.LastResult != "" && state.LastResult != "passed" && state.LastResult != "failed") {
		return State{}, errors.New("invalid workflow state")
	}
	return state, nil
}

func Require(project string, context resolver.Context, command, required string) (State, error) {
	state, err := Current(project, context)
	if err != nil {
		return State{}, err
	}
	if state.FailedCommand != nil {
		if command == "execute" && *state.FailedCommand == "verify" && state.Checkpoint == Executed {
			return state, nil
		}
		return State{}, fmt.Errorf("IASI workflow blocked: %s failed", *state.FailedCommand)
	}
	if state.Checkpoint != required {
		return State{}, fmt.Errorf("IASI workflow blocked: %s requires %s; current checkpoint is %s", command, required, state.Checkpoint)
	}
	return state, nil
}

func Succeed(project string, context resolver.Context, command, checkpoint string) error {
	state, err := Current(project, context)
	if err != nil {
		return err
	}
	if !allowedTransition(state, command, checkpoint) {
		return fmt.Errorf("invalid workflow transition: %s from %s to %s", command, state.Checkpoint, checkpoint)
	}
	state.Checkpoint = checkpoint
	state.LastCommand = command
	state.LastResult = "passed"
	state.FailedCommand = nil
	return Write(project, state)
}

func Fail(project string, context resolver.Context, command string) error {
	state, err := Current(project, context)
	if err != nil {
		return err
	}
	state.LastCommand = command
	state.LastResult = "failed"
	state.FailedCommand = &command
	return Write(project, state)
}

func Write(project string, state State) error {
	state.Version = 1
	state.UpdatedAt = time.Now().Format(time.RFC3339)
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	path := filepath.Join(project, filepath.FromSlash(statePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	temporary, err := os.CreateTemp(filepath.Dir(path), "workflow-*.json")
	if err != nil {
		return err
	}
	temporaryPath := temporary.Name()
	defer os.Remove(temporaryPath)
	if _, err := temporary.Write(append(data, '\n')); err != nil {
		temporary.Close()
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	return os.Rename(temporaryPath, path)
}

func initial(project string, context resolver.Context) (State, error) {
	inputs, err := validation.HashInputs(project)
	if err != nil {
		return State{}, err
	}
	plan, err := HashPlan(project)
	if err != nil {
		return State{}, err
	}
	return State{Version: 1, Checkpoint: Inputs, InputsHash: inputs, InstructionsHash: validation.HashInstructions(context), PlanHash: plan}, nil
}

func refresh(project string, context resolver.Context, state State) (State, error) {
	fresh, err := initial(project, context)
	if err != nil {
		return State{}, err
	}
	if state.InputsHash != fresh.InputsHash || state.InstructionsHash != fresh.InstructionsHash || !samePlan(state.PlanHash, fresh.PlanHash) {
		return fresh, nil
	}
	return state, nil
}

func HashPlan(project string) (*string, error) {
	root := filepath.Join(project, "inputs", "obtained")
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	hash, err := validation.HashInputsAt(root, filepath.Join(root, "archived"))
	if err != nil {
		return nil, err
	}
	return &hash, nil
}

func samePlan(left, right *string) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	return *left == *right
}

func validCheckpoint(value string) bool {
	return value == Inputs || value == InputsValidated || value == Planned || value == PlanValidated || value == Executed || value == Verified
}

func allowedTransition(state State, command, checkpoint string) bool {
	if command == "validate" {
		return (state.Checkpoint == Inputs && checkpoint == InputsValidated) || (state.Checkpoint == Planned && checkpoint == PlanValidated)
	}
	if command == "plan" {
		return state.Checkpoint == InputsValidated && checkpoint == Planned
	}
	if command == "execute" {
		return (state.Checkpoint == PlanValidated || (state.Checkpoint == Executed && state.FailedCommand != nil && *state.FailedCommand == "verify")) && checkpoint == Executed
	}
	return command == "verify" && state.Checkpoint == Executed && checkpoint == Verified
}
