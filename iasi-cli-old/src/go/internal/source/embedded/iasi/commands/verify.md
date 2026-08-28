# /verify

Require the current shared IASI workflow checkpoint to be `EXECUTED` for the
exact active instructions, source inputs, plan, and execution context. An
unresolved execution failure blocks verification.

Use the shared runtime gate and transition:

```text
iasi __runtime workflow require verify EXECUTED
...
iasi __runtime workflow transition verify VERIFIED
```

Evaluate whether the project state satisfies the validated plan and governing
source inputs. Run objective checks when available, including builds, tests,
acceptance checks, required file structures, or behaviours defined by the
plan. Semantic review may complement those checks.

Do not repair defects while verifying. Report failures and return corrective
work to `/execute` for the same current validated plan.

On success, persist shared workflow state with checkpoint `VERIFIED`. On
failure, retain the last valid execution checkpoint, persist
`failed_command = verify`, and block later forward workflow until corrective
execution succeeds and `/verify` runs again. Record failure through
`iasi __runtime workflow fail verify`.