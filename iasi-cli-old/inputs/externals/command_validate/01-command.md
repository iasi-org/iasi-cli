# `/validate` command

## Purpose

`/validate` is the semantic quality gate of the IASI workflow.

It has two workflow modes selected from current workflow state:

```text
pre-plan validation
post-plan validation
```

The caller does not choose a mode flag.

IASI determines the required mode from the current checkpoint and current content.

## Canonical definition

```text
iasi/commands/validate.md
```

## Pre-plan validation

Purpose:

> determine whether the current source inputs are coherent and sufficient for `/plan` to derive an executable plan without inventing material decisions.

Sources:

```text
effective IASI instructions
+
active inputs/externals
+
active inputs/internals
```

The current `inputs/obtained/` plan is NOT part of pre-plan semantic validation.

This is intentional: it may belong to the previous iteration and `/plan` is responsible for archiving it before producing the next plan.

Successful pre-plan validation produces:

```text
INPUTS_VALIDATED
```

Failure blocks `/plan`.

## Post-plan validation

Purpose:

> determine whether the newly obtained current plan is coherent with the source inputs and instructions and is safe to execute.

Sources:

```text
effective IASI instructions
+
active inputs/externals
+
active inputs/internals
+
active inputs/obtained
```

Successful post-plan validation produces:

```text
PLAN_VALIDATED
```

Failure blocks `/execute`.

## Archived content

Archived content is historical and never participates in either mode.

## Outcomes

```text
PASSED
PASSED WITH WARNINGS
FAILED
```

Any blocker means `FAILED`.

Warnings do not block progression.

## Fundamental rule

The validator MUST NOT invent a material product, architecture, design or implementation decision to make validation pass.

If the next stage would need such an invention, report the gap.

## Execution failure versus semantic failure

A semantic validation failure and inability to execute validation are distinct conditions.

The workflow gate is satisfied only by a successfully executed semantic validation whose result is `passed`.
