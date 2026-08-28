# `/validate` acceptance criteria

The `/validate` implementation and its first Copilot projection are complete when the following behaviour is demonstrated.

## Canonical command

- `iasi/commands/validate.md` exists as the canonical command definition.
- `iasi install` installs it under `.iasi/commands/validate.md`.
- composed command resolution selects the nearest `validate` command when multiple layers define it.

## Copilot adapter

- `iasi adapt copilot` generates `.github/prompts/validate.prompt.md` from the effective canonical command;
- the generated prompt carries IASI ownership metadata;
- the prompt does not contain a separately maintained validation contract.

## Semantic validation

Test representative cases for:

- contradiction between inputs;
- input versus effective instruction conflict;
- missing material information;
- material ambiguity;
- invalid required reference;
- incompatible constraints;
- unresolved required decision;
- warning-only outcome.

The validator MUST not silently fill blockers with invented design decisions.

## Persistent gate

Verify that semantic validation writes:

```text
<project>/.iasi/validation.json
```

on both pass and semantic failure.

Verify that a failed state blocks the next phase.

Verify that a warning-only state permits the next phase.

## Staleness

After a successful validation, independently modify:

1. a project input;
2. an effective instruction;
3. the effective `validate` command.

Each change MUST make the previous validation state stale and block the next phase until `/validate` runs again.

## Inheritance boundary

Verify that:

- effective instructions include inherited parent layers plus local overrides;
- validation state is local to the project and is never inherited;
- parent validation state cannot authorize a child project.

## Installation coexistence

Verify that a project with only `.iasi/validation.json` is not reported as having a local installed layer.

Verify that `iasi install` can create the local installation while preserving the existing validation state, subject to the installation safety rules.
