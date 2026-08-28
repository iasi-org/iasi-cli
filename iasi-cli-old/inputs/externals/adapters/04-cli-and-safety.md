# `iasi adapt copilot` and safety contract

## Command

```bash
iasi adapt copilot
```

acts on the current working directory as the target project root.

No additional scope flag is required.

## Source

The command MUST resolve every applicable installed IASI layer and build the effective context described by `iasi-cli/03-resolution.md`.

It MUST NOT use only the nearest installation.

## Target boundary

All adapter-generated project artifacts and adapter metadata MUST remain inside:

```text
<cwd>/.github/
```

Managed metadata lives at:

```text
.github/.iasi/copilot-manifest.yml
```

Paths MUST be normalized and validated before any write. Absolute or traversing targets that escape `.github/` are invalid.

## Preflight and atomicity

Adaptation is logically:

```text
resolve
  ↓
preflight
  ↓
staging
  ↓
commit
```

Before writing, preflight MUST:

1. resolve and validate the effective IASI context;
2. resolve and validate the selected Copilot adapter;
3. resolve effective instructions and commands;
4. calculate every target file;
5. validate existing Copilot manifest if present;
6. inspect collisions and stale outputs;
7. establish ownership unambiguously.

Any preflight error leaves the project unchanged.

Staging generates complete new content in temporary storage.

Commit begins only after preflight and staging succeed. If commit fails, the implementation MUST attempt to restore the prior state of every affected file.

## Existing targets

If a target does not exist, create it.

If it exists and contains the IASI ownership marker, it may be replaced.

If it exists without the IASI ownership marker, fail without modifying any target.

No `--force` behaviour is required.

## Copilot manifest

The absence of the manifest means first adaptation and is valid.

Minimum conceptual manifest:

```yaml
schema_version: 1
adapter: copilot
context_fingerprint: <sha256>

generated:
  - .github/copilot-instructions.md
  - .github/prompts/validate.prompt.md
```

Do not store a single `iasi_version` as the identity of the effective context. Multiple inherited layers may have different versions.

`context_fingerprint` MUST deterministically identify the effective resolved source artifacts and selected adapter descriptor used to generate the projection.

An existing manifest MUST be validated before it is trusted for ownership or stale-output decisions.

## Stale outputs

A previously generated file may be deleted as stale only when:

- it appears in the prior valid Copilot manifest;
- its normalized path remains inside `.github/`;
- it exists;
- it still carries the IASI ownership marker.

Any ownership ambiguity MUST fail rather than delete.

## Runtime independence

Adaptation MUST continue to work when the source repository is absent.

It MUST reflect deliberate content changes made inside installed layers and MUST NOT silently fall back to the executable's embedded distribution.
