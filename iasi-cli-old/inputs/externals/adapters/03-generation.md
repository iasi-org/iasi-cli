# Copilot generation rules

## Effective source

Generation MUST use the effective composed IASI context for the current working directory.

Conceptually:

```text
all applicable .iasi layers
          ↓
composed effective context
          ↓
selected Copilot adapter
          ↓
project .github/
```

The adapter MUST NOT read the source repository or silently substitute the methodology embedded in the executable.

## Instruction generation

Instruction candidates come from the effective resolved instruction set after layer precedence has been applied.

Exclude support content before semantic validation:

```text
schema/**
README.md
README_*.md
```

A candidate instruction MUST have valid IASI front matter including:

```yaml
id:
status:
scope:
```

Valid status values:

```text
active
draft
deprecated
```

Behaviour:

```text
active       → validate and generate
draft        → validate, do not generate
deprecated   → validate, do not generate
```

Unknown or missing status is an error.

An active instruction with a scope not mapped by the Copilot adapter is a preflight error.

### Content transformation

Generation is deterministic and mechanical:

1. remove IASI YAML front matter;
2. preserve Markdown body;
3. preserve normative terms and constraints;
4. do not paraphrase;
5. group by mapped scope;
6. order generated instruction sections lexicographically by instruction ID.

Path-specific outputs begin with the configured Copilot `applyTo` front matter.

Empty mapped groups do not create empty files.

## Command generation

Command candidates come from the effective resolved command set after layer precedence has been applied.

For each command supported by the adapter:

```text
commands/<identity>.md
      ↓
.github/prompts/<identity>.prompt.md
```

For the first command:

```text
commands/validate.md
      ↓
.github/prompts/validate.prompt.md
```

Generation MUST preserve the canonical command semantics. Platform-specific framing MAY be added only when required to invoke the command correctly in Copilot.

The generated prompt MUST NOT introduce independent validation rules, alter blocker/warning semantics or choose a different workflow state contract.

## Generated ownership marker

Every generated file MUST contain a stable machine-recognizable IASI ownership marker near the beginning.

For example:

```html
<!-- IASI-GENERATED: copilot; schema=1 -->
```

Do not encode a single "installed IASI version" in the marker because a composed effective context may contain several installed layer versions.

A short human-readable warning that the file is generated SHOULD also be present.

## Determinism

Given the same effective resolved artifact set and the same selected adapter descriptor, generated target contents MUST be byte-for-byte equivalent.

Absolute installation paths MUST NOT be embedded into generated Copilot artifacts.
