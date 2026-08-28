# Persistent workflow state

## Purpose

Workflow gates must survive individual agent calls and platform sessions.

IASI maintains local machine-readable workflow state in:

```text
.iasi/workflow.json
```

This state is local to the current project and is never inherited from parent installations.

## Minimum structure

```json
{
  "version": 1,
  "checkpoint": "PLAN_VALIDATED",
  "last_command": "validate",
  "last_result": "passed",
  "failed_command": null,
  "source_inputs_hash": "...",
  "instructions_hash": "...",
  "plan_hash": "...",
  "updated_at": "2026-08-17T00:50:00+02:00"
}
```

## Fingerprints

### `source_inputs_hash`

Fingerprint only the active source-input set:

```text
inputs/externals/
inputs/internals/
```

excluding exactly:

```text
inputs/externals/archived/
inputs/internals/archived/
```

### `plan_hash`

Fingerprint only the active current plan:

```text
inputs/obtained/
```

excluding exactly:

```text
inputs/obtained/archived/
```

It may be `null` before a current plan exists.

### `instructions_hash`

Fingerprint the effective composed instruction set.

## Failure persistence

When a forward stage fails, persist:

- the highest still-valid checkpoint;
- `last_result = failed`;
- `failed_command = <command>`.

A stale earlier success MUST NOT be allowed to authorize later work.

## Currentness

Every gated command MUST recompute the fingerprints relevant to its checkpoint before trusting `workflow.json`.

If source inputs or instructions change, the effective checkpoint falls back to `INPUTS`.

If the current plan changes while source inputs and instructions remain current, the effective checkpoint cannot be later than `PLANNED`.

## Atomic writes

Workflow-state writes MUST be atomic where practical.

A command MUST NOT report workflow success before the corresponding state transition has been persisted.
