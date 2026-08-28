# /execute

Require the current shared IASI workflow checkpoint to be `PLAN_VALIDATED` for
the exact active source-input, instruction, and plan context. If the gate is
missing, failed, or stale, stop before changing the project.

Use the shared runtime gate and transition:

```text
iasi __runtime workflow require execute PLAN_VALIDATED
...
iasi __runtime workflow transition execute EXECUTED
```

Execute only work required by the current validated plan under
`inputs/obtained/`. Use effective IASI instructions and active source inputs in
`inputs/externals/` and `inputs/internals/`; ignore archived content.

Do not expand scope, add unrelated improvements, replace the validated plan, or
invent material unsupported decisions. If execution exposes such a decision,
report the blocker and fail.

On success, persist shared workflow state with checkpoint `EXECUTED`, clear any
resolved failure, and allow `/verify`. On failure, preserve the highest earlier
valid checkpoint, persist `failed_command = execute`, and block `/verify`. A
corrective retry may run only against the same still-current validated plan.
Record failure through `iasi __runtime workflow fail execute`.