# Instruction schema

This document defines the minimum structure of an iasi instruction.

## Metadata

Every instruction MUST start with YAML front matter containing:

```yaml
---
id: <scope>.<name>
version: 0.1.0
status: draft | active | deprecated
scope: general | documentation | code | diagrams | <other-domain>
applies_to:
  - <task-or-artifact-type>
---
```

Optional fields:

```yaml
extends:
  - <instruction-id>
tags:
  - <tag>
```

### `id`

A stable, lowercase identifier using dot notation.

Examples:

```text
general.behavior
documentation.style
code.testing
```

### `version`

The version of the instruction itself. Start at `0.1.0` while the schema is experimental.

### `status`

The lifecycle state of the instruction.

### `scope`

The domain in which the instruction belongs.

### `applies_to`

The kinds of tasks, artifacts or agents for which the instruction is relevant.

### `extends`

Optional instruction IDs whose rules are explicitly refined by this instruction.

## Required sections

Every instruction MUST contain these sections.

### Purpose

Why the instruction exists and what behavior it is intended to control.

### Rules

Positive rules the agent MUST or SHOULD follow.

Use:

- `MUST` for mandatory behavior.
- `SHOULD` for the expected default when no justified exception exists.
- `MAY` for permitted behavior.

Each rule SHOULD describe one observable behavior.

### Constraints

Behavior the agent MUST NOT perform.

Constraints exist to make boundaries explicit rather than relying on implication.

### Validation

Checks that can be used to determine whether the instruction was respected.

Validation SHOULD focus on observable output or behavior rather than hidden reasoning.

## Optional sections

An instruction MAY also contain:

- `Examples`
- `Rationale`
- `Exceptions`
- `Related instructions`

Add optional sections only when they improve interpretation or validation.

## Design rules

An instruction MUST NOT:

- depend on a specific AI vendor unless its scope explicitly targets an adapter;
- mix unrelated concerns only to reduce the number of files;
- duplicate another instruction without adding a meaningful refinement;
- require access to hidden reasoning;
- contain project-specific facts in a shared instruction;
- claim that a task succeeded without an observable validation path.

## Conflict handling

Instructions are additive by default.

When two instructions conflict:

1. explicit task or project instructions take precedence over domain defaults;
2. domain-specific instructions take precedence over general instructions;
3. the more specific rule overrides the broader rule only for the conflicting behavior;
4. non-conflicting rules remain active;
5. unresolved conflicts MUST be surfaced rather than silently ignored.
