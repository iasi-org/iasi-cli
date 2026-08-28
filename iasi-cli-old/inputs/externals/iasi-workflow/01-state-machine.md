# IASI workflow state machine

## Purpose

IASI commands form a controlled workflow.

They are not independent slash commands that may execute in arbitrary order.

A later forward stage MUST NOT execute unless every required earlier stage has succeeded for the current exact context.

## Core sequence

```text
source inputs
    ↓
/validate              pre-plan validation
    ↓ PASS
/plan
    ↓ PASS
/validate              post-plan validation
    ↓ PASS
/execute
    ↓ PASS
/verify
    ↓ PASS
VERIFIED
```

`/archive` is a lifecycle operation specified separately.

## Source inputs and plan inputs

IASI distinguishes:

```text
source inputs = active inputs/externals + active inputs/internals
plan inputs   = active inputs/obtained
```

Archived subtrees are historical and are excluded.

This distinction prevents a previous iteration's obtained plan from blocking pre-plan validation of a new iteration.

## Checkpoints

```text
INPUTS
INPUTS_VALIDATED
PLANNED
PLAN_VALIDATED
EXECUTED
VERIFIED
```

### INPUTS

Current source inputs have not yet passed pre-plan validation.

### INPUTS_VALIDATED

Current source inputs have passed pre-plan validation.

`/plan` may execute.

### PLANNED

`/plan` has archived the previous obtained plan, when present, and generated the new current plan in `inputs/obtained/`.

The new plan has not yet passed post-plan validation.

### PLAN_VALIDATED

Current source inputs plus the current obtained plan have passed post-plan validation.

`/execute` may execute.

### EXECUTED

`/execute` has successfully performed the current validated plan.

`/verify` may execute.

### VERIFIED

`/verify` has successfully verified the current execution against the current validated plan and source inputs.

## Failure rule

If a stage fails, no later forward stage may execute.

Examples:

```text
pre-plan /validate → FAILED
/plan              → BLOCKED
```

```text
/plan     → FAILED
post-plan /validate → BLOCKED until plan succeeds
```

```text
post-plan /validate → FAILED
/execute            → BLOCKED
```

```text
/execute → FAILED
/verify  → BLOCKED
```

## Recovery

The failed stage may be retried after correcting its cause.

A failed `/verify` may return control to `/execute` for corrective execution against the same current validated plan. After corrective `/execute` succeeds, `/verify` must run again.

Any source-input or instruction change invalidates downstream checkpoints and requires pre-plan `/validate` again.

Any new or changed current obtained plan invalidates execution and verification and requires post-plan `/validate` again.

## No implicit skipping

No adapter or agent may skip a checkpoint because the task appears trivial.

## Context identity

A checkpoint means:

```text
this stage succeeded for these exact source inputs,
these exact applicable instructions,
and, where applicable, this exact current plan
```

A historical success for different content never authorizes current work.
