# Validation state

## State location

```text
<cwd>/.iasi/validation.json
```

Validation state is local and never inherited.

## Minimum structure

```json
{
  "schema_version": 1,
  "phase": "post_plan",
  "status": "passed",
  "validated_at": "2026-08-17T00:50:00+02:00",
  "instructions_hash": "...",
  "source_inputs_hash": "...",
  "plan_hash": "...",
  "command_hash": "...",
  "blockers": 0,
  "warnings": 1
}
```

Allowed `phase` values:

```text
pre_plan
post_plan
```

Allowed `status` values:

```text
passed
failed
```

## Pre-plan state

For `phase = pre_plan`:

- `source_inputs_hash` is required;
- `plan_hash` is `null`;
- a successful result authorizes `/plan`, not `/execute`.

## Post-plan state

For `phase = post_plan`:

- `source_inputs_hash` is required;
- `plan_hash` is required;
- a successful result authorizes `/execute`.

## Fingerprints

### `instructions_hash`

Effective composed instructions.

### `source_inputs_hash`

Active files under:

```text
inputs/externals/
inputs/internals/
```

excluding only their exact root `archived/` subtrees.

### `plan_hash`

Active current plan under:

```text
inputs/obtained/
```

excluding only:

```text
inputs/obtained/archived/
```

### `command_hash`

Effective canonical `/validate` command content.

## Coordination with workflow state

`/validate` MUST update `.iasi/validation.json` and `.iasi/workflow.json` consistently.

A validation result is not a completed workflow gate unless both states reflect the same successful phase and fingerprints.

## Staleness

Pre-plan validation becomes stale when:

- source inputs change;
- effective instructions change;
- `/validate` semantics change.

Post-plan validation becomes stale when any of the above changes or when the current obtained plan changes.

`iasi reinstall` may preserve local state files; fingerprints determine whether preserved state remains current.
