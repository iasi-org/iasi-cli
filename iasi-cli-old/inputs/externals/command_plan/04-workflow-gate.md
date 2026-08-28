# `/plan` workflow integration

## Required gate

`/plan` requires:

```text
INPUTS_VALIDATED
```

for the current pre-plan active input context.

If that checkpoint is missing, failed or stale, `/plan` MUST stop before archiving the previous plan or generating a new one.

## Successful transition

After the previous plan has been safely archived and the new plan has been generated successfully:

```text
checkpoint = PLANNED
```

The new plan is not yet authorized for execution.

`/execute` remains blocked.

The required next command is:

```text
/validate
```

## Failure

If `/plan` fails:

- `PLANNED` MUST NOT be recorded;
- downstream stages remain blocked;
- `failed_command` is set to `plan`;
- the operation must preserve the previous plan according to the plan atomicity rules.

A successful retry clears the unresolved plan failure and transitions to `PLANNED`.
