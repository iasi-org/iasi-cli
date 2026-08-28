# `/validate` workflow integration

`/validate` participates in two workflow checkpoints.

## Pre-plan validation

When validating the active iteration before a new plan is generated, successful validation produces:

```text
INPUTS_VALIDATED
```

Failure blocks `/plan`.

## Post-plan validation

After `/plan` generates the current active plan under:

```text
inputs/obtained/
```

successful validation produces:

```text
PLAN_VALIDATED
```

Failure blocks `/execute`.

## State coordination

`/validate` MUST update both:

```text
.iasi/validation.json
```

and:

```text
.iasi/workflow.json
```

consistently.

A validation success that cannot be persisted into workflow state MUST NOT be treated as a completed workflow gate.

## Failure

A failed validation MUST persist the failure and MUST NOT leave a previous validation checkpoint usable for changed content.

No downstream stage may execute from a stale validation success.
