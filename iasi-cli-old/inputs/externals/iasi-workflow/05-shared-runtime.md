# Shared IASI runtime services

## Purpose

Canonical Markdown commands describe behaviour.

They are not, by themselves, an implementation of workflow state, filesystem mutation or safety.

IASI MUST provide shared runtime services used by every canonical command.

The exact Go package names are implementation details, but the behaviour MUST be implemented once and reused.

## Required shared services

### Workflow state

Provide shared operations for:

- reading `.iasi/workflow.json`;
- validating its schema;
- recomputing current fingerprints;
- evaluating required checkpoints;
- recording success and failure;
- invalidating downstream checkpoints;
- atomic state writes.

### Validation state

Provide shared operations for:

- reading and writing `.iasi/validation.json`;
- distinguishing pre-plan and post-plan validation;
- checking currentness against the appropriate source/plan fingerprints.

### Active-input discovery

Provide one canonical discovery implementation.

Recognized logical branches are:

```text
inputs/externals/
inputs/internals/
inputs/obtained/
```

Only these exact archive roots are historical:

```text
inputs/externals/archived/
inputs/internals/archived/
inputs/obtained/archived/
```

A nested directory with the same basename is not automatically historical.

For example:

```text
inputs/externals/design/archived/
```

is an active input directory unless another specification explicitly says otherwise.

### Deterministic fingerprints

Hash normalized relative paths and file contents in deterministic sorted order.

Do not depend on filesystem enumeration order or absolute paths.

### Safe filesystem operations

Provide shared safe operations for command-owned mutations such as:

- moving one input into its branch `archived/`;
- snapshotting the complete current obtained plan;
- replacing the current obtained plan;
- creating required directories.

Before mutation:

- resolve paths;
- enforce branch boundaries;
- reject traversal escapes;
- reject symlink escapes;
- reject accidental overwrite.

### Atomicity and rollback safety

Plan replacement and archive operations MUST NOT destroy the only copy of active content when a later write fails.

Use temporary paths, verification and atomic rename/move where the filesystem permits.

## Adapter boundary

Adapters project canonical IASI commands to platform-native files.

Adapters MUST NOT own:

- workflow gates;
- state transitions;
- active-input discovery;
- path-safety rules;
- archive semantics;
- plan-replacement semantics.

Those belong to the shared IASI runtime.

## Implementation completeness

A canonical command is not considered implemented merely because:

```text
iasi/commands/<name>.md
```

and:

```text
.github/prompts/<name>.prompt.md
```

exist.

The shared runtime behaviour required by that command and its tests must also exist.
