# `/execute` command

## Purpose

`/execute` performs the current validated plan.

It is the stage that may change the project implementation.

It MUST execute the plan that exists under:

```text
inputs/obtained/
```

for the exact source-input and instruction context that passed post-plan validation.

## Canonical command

```text
iasi/commands/execute.md
```

## Required gate

```text
PLAN_VALIDATED
```

If the checkpoint is missing, failed or stale, `/execute` MUST stop before changing the project.

## Sources

`/execute` uses:

- effective IASI instructions;
- active source inputs in `inputs/externals/` and `inputs/internals/`;
- the current validated plan in `inputs/obtained/`.

Archived content is historical and ignored.

## Execution rule

Execute only work required by the validated current plan.

`/execute` MUST NOT:

- expand the iteration scope;
- silently add unrelated improvements;
- replace the validated plan with a new plan;
- invent a material unsupported decision.

If execution reveals a missing material decision, `/execute` fails and reports the blocker.

## Success

Successful execution produces:

```text
EXECUTED
```

and permits `/verify`.

## Failure

Failure MUST:

- persist `failed_command = execute`;
- preserve the highest earlier valid checkpoint;
- block `/verify`;
- report the execution blocker or error.

A failed execution may be retried after correction without skipping the gate.
