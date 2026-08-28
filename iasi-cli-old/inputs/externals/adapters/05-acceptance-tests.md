# Copilot adapter acceptance tests

## Composed resolution

Create:

```text
workspace/.iasi/
workspace/project/.iasi/
workspace/project/src/
```

Run from `workspace/project/src/` or another chosen target root with known layers.

Verify that the adapter uses the complete effective context rather than only the nearest layer.

Test at least:

- parent instruction inherited when absent locally;
- child instruction overrides parent instruction with the same ID;
- child-only instruction is added;
- Copilot adapter inherited when it exists only in the parent;
- child Copilot adapter replaces the parent adapter atomically when both exist;
- invalid nearest Copilot adapter fails instead of falling back to the parent.

## Instructions

Verify generation of repository-wide and path-specific files according to `adapter.yml`.

Verify:

- only effective `active` instructions are generated;
- `draft` and `deprecated` remain valid but are omitted;
- schema/readme support files are excluded;
- unknown statuses fail;
- active unknown scopes fail;
- IASI front matter is not copied into generated instruction content;
- output ordering is deterministic.

## Commands

Given an effective:

```text
commands/validate.md
```

verify creation of:

```text
.github/prompts/validate.prompt.md
```

Verify that:

- it contains the IASI generated ownership marker;
- its semantics originate from the effective canonical command;
- a nearer `commands/validate.md` overrides an inherited one;
- the adapter does not maintain a separate hand-written validation definition.

## Human-owned collisions

Pre-create a required target without the IASI ownership marker.

Verify that adaptation fails and leaves every existing project file unchanged.

## First run and regeneration

Verify first adaptation without a prior Copilot manifest.

Run adaptation again unchanged and verify byte-for-byte idempotency.

Change an effective installed instruction or command and verify regeneration from installed state.

## Stale outputs

Verify that stale files are removed only when both prior manifest ownership and generated marker ownership are established.

## Context manifest

Verify that the generated Copilot manifest records a deterministic `context_fingerprint` rather than assuming a single installed IASI version.

## Independence

Run acceptance tests with a standalone executable and installed `.iasi` layers while the source repository is unavailable.


## All current canonical commands

With the current canonical command set installed:

```text
commands/validate.md
commands/plan.md
commands/execute.md
commands/verify.md
commands/archive.md
```

`iasi adapt copilot` MUST project the corresponding prompt files:

```text
.github/prompts/validate.prompt.md
.github/prompts/plan.prompt.md
.github/prompts/execute.prompt.md
.github/prompts/verify.prompt.md
.github/prompts/archive.prompt.md
```

Each prompt remains a thin projection of the effective canonical command and MUST NOT implement independent workflow gates or filesystem semantics.
