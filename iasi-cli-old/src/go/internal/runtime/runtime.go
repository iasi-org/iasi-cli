package runtime

import (
	"fmt"
	"strconv"

	"iasi-cli/internal/resolver"
	"iasi-cli/internal/validation"
	"iasi-cli/internal/workflow"
)

type Error struct {
	Code int
	Err  error
}

func (err *Error) Error() string { return err.Err.Error() }

func Run(project string, args []string) (string, error) {
	if len(args) == 0 {
		return "", reject("missing runtime operation")
	}
	context, err := resolver.Resolve(project)
	if err != nil {
		return "", internal(err)
	}
	switch args[0] {
	case "workflow":
		return runWorkflow(project, context, args[1:])
	case "validate":
		return runValidation(project, context, args[1:])
	default:
		return "", reject("unknown runtime operation")
	}
}

func runWorkflow(project string, context resolver.Context, args []string) (string, error) {
	if len(args) == 0 {
		return "", reject("missing workflow operation")
	}
	switch args[0] {
	case "status":
		if len(args) != 1 {
			return "", reject("usage: iasi __runtime workflow status")
		}
		state, err := workflow.Current(project, context)
		if err != nil {
			return "", internal(err)
		}
		return fmt.Sprintf("%s\n", state.Checkpoint), nil
	case "require":
		if len(args) != 3 {
			return "", reject("usage: iasi __runtime workflow require <command> <checkpoint>")
		}
		if _, err := workflow.Require(project, context, args[1], args[2]); err != nil {
			return "", reject(err.Error())
		}
		return "", nil
	case "transition":
		if len(args) != 3 {
			return "", reject("usage: iasi __runtime workflow transition <command> <checkpoint>")
		}
		if err := workflow.Succeed(project, context, args[1], args[2]); err != nil {
			return "", reject(err.Error())
		}
		return "", nil
	case "fail":
		if len(args) != 2 {
			return "", reject("usage: iasi __runtime workflow fail <command>")
		}
		if err := workflow.Fail(project, context, args[1]); err != nil {
			return "", internal(err)
		}
		return "", nil
	default:
		return "", reject("unknown workflow operation")
	}
}

func runValidation(project string, context resolver.Context, args []string) (string, error) {
	if len(args) != 3 {
		return "", reject("usage: iasi __runtime validate <passed|failed> <blockers> <warnings>")
	}
	blockers, err := strconv.Atoi(args[1])
	if err != nil || blockers < 0 {
		return "", reject("invalid blocker count")
	}
	warnings, err := strconv.Atoi(args[2])
	if err != nil || warnings < 0 {
		return "", reject("invalid warning count")
	}
	if args[0] == "passed" && blockers != 0 {
		return "", reject("passed validation cannot have blockers")
	}
	state, err := validation.NewState(project, context, args[0], blockers, warnings)
	if err != nil {
		return "", internal(err)
	}
	if err := validation.Write(project, state); err != nil {
		return "", internal(err)
	}
	if args[0] == "failed" {
		if err := workflow.Fail(project, context, "validate"); err != nil {
			return "", internal(err)
		}
		return "", nil
	}
	workflowState, err := workflow.Current(project, context)
	if err != nil {
		return "", internal(err)
	}
	target := workflow.InputsValidated
	if workflowState.Checkpoint == workflow.Planned {
		target = workflow.PlanValidated
	}
	if err := workflow.Succeed(project, context, "validate", target); err != nil {
		return "", reject(err.Error())
	}
	return "", nil
}

func reject(message string) error { return &Error{Code: 1, Err: fmt.Errorf("%s", message)} }
func internal(err error) error    { return &Error{Code: 2, Err: err} }
