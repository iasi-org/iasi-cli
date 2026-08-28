# Workflow command contract

## Shared command behaviour

Every forward IASI workflow command MUST follow this sequence:

```text
1. resolve effective IASI context
2. load workflow state
3. verify fingerprints/currentness
4. check required checkpoint
5. reject unresolved predecessor failure
6. perform command work
7. persist command result
8. transition workflow state on success
9. persist failure state on failure
```

## Success

A command may transition the workflow only after its own work has completed successfully.

It MUST NOT pre-authorize the next phase.

Example:

```text
/plan started
```

does not mean:

```text
PLANNED
```

Only successful plan generation means `PLANNED`.

## Failure

A failed command MUST:

- leave no false success checkpoint;
- persist which stage failed;
- block all later forward stages;
- allow the failed stage to be retried;
- provide a concise explanation of the blocking reason.

## Recovery

Correcting inputs or instructions may require returning to `/validate`.

Generating a replacement plan requires post-plan `/validate` again.

The workflow always moves forward from the most recent valid checkpoint for the current exact context.

## Platform independence

Copilot, Codex and future platforms must all observe the same workflow state and gates.

A platform-native prompt MUST NOT bypass IASI state because the platform itself believes a command can run.

The dependency direction remains:

```text
platform adapter
      ↓
IASI command
      ↓
IASI workflow gate
      ↓
command work
```

## Acceptance criteria

The gate system is complete when:

1. `/plan` cannot execute without successful pre-plan validation;
2. `/execute` cannot execute until the generated plan has passed post-plan validation;
3. `/verify` cannot execute until `/execute` succeeds;
4. failure of any stage blocks every later forward stage;
5. retrying the failed stage is possible;
6. upstream content changes invalidate dependent downstream checkpoints;
7. gate state persists across agent/platform calls;
8. stale fingerprints cannot authorize work;
9. gate logic is shared rather than reimplemented independently per command;
10. adapters cannot weaken or bypass the IASI workflow.
