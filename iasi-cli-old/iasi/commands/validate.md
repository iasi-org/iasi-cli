# /validate

Evaluate the current project's `inputs/` against the effective composed IASI
instructions and this command definition before the next workflow phase. Read
all applicable installed `.iasi` layers from parent to child; do not treat only
the nearest layer as authoritative.

Use these sources:

- the effective active IASI instructions;
- `<project>/inputs/`;
- this effective `/validate` command.

Ignore all historical subtrees named `archived/` below `inputs/externals/`,
`inputs/internals/`, and `inputs/obtained/`. Archived documents must not affect
validation, findings, input discovery, or active-input fingerprints.

Check for:

1. project requirements that conflict with effective active instructions;
2. material contradictions between inputs;
3. missing information that would force a material decision in the next phase;
4. material ambiguity with multiple plausible outcomes;
5. invalid required references;
6. incompatible constraints;
7. decisions explicitly left unresolved when the next phase needs them;
8. divergent duplicate definitions that obscure authority.

Classify each finding as `BLOCKER` or `WARNING`. A blocker means the next phase
would need an unsafe assumption, violate instructions, choose between
incompatible requirements, or use unusable information. Any blocker produces
`FAILED`; otherwise produce `PASSED`, or `PASSED WITH WARNINGS` when warnings
exist.

Do not modify inputs, instructions, or installed layers. Do not resolve
contradictions, select technologies, or invent material decisions to make
validation pass. Explain missing or conflicting information without rewriting
user material.

Persist local validation state at `.iasi/validation.json` for both successful
and semantic-failure outcomes. Use `schema_version: 1`, `status` (`passed` or
`failed`), `validated_at`, `instructions_hash`, `inputs_hash`, `command_hash`,
`blockers`, and `warnings`. Hash effective instructions, normalized project
input paths and contents, and this effective command deterministically with
SHA-256. A failed, missing, or stale local state blocks later IASI workflow
phases. Validation state is local to the project and is never inherited.

Persist shared `.iasi/workflow.json` atomically with the validation result. A
successful validation from `INPUTS` produces `INPUTS_VALIDATED`; a successful
validation after `PLANNED` produces `PLAN_VALIDATED`. A failure persists
`failed_command = validate` and must not leave a prior checkpoint usable.

After determining the result, invoke the shared runtime rather than writing
either state file directly:

```text
iasi __runtime validate <passed|failed> <blockers> <warnings>
```