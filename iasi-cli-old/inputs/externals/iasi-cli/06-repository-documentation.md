# Repository documentation

## Purpose

Repository-level documentation MUST describe the same current architecture as the normative IASI inputs.

Documentation is not an independent architectural source of truth. The normative architecture is defined by the documents under:

```text
inputs/externals/iasi-cli/
```

When repository documentation conflicts with these inputs, the documentation MUST be updated.

## Required README files

The repository currently maintains:

```text
README.md
README_en.md
```

Both files MUST be kept aligned with the canonical IASI structure.

## Canonical structure

Repository documentation MUST describe IASI-owned artifacts under:

```text
iasi/
├── instructions/
├── commands/
├── skills/
├── mcp/
└── adapters/
```

The former architectural model based on `agentics/` is superseded.

`agentics/` MUST NOT be presented as the current canonical location for instructions, commands, adapters, skills or MCP artifacts.

Historical references are acceptable only when explicitly identified as historical context and when they cannot be mistaken for current implementation guidance.

## Installed structure

Documentation describing an installed IASI layer MUST use `.iasi/manifest.yml` as the marker of a valid installed layer and MUST reflect the same artifact categories used by the canonical source tree.

A generated `.iasi/validation.json` is local workflow state and MUST NOT be described as an installed layer marker.

## CLI documentation

Public CLI examples MUST reflect the current interface:

```text
iasi install
iasi reinstall
iasi status
iasi version
iasi adapt copilot
```

Do not document obsolete scope flags such as:

```text
--workspace
--project
```

Do not describe `workspace` as a current installation profile unless such a profile is reintroduced by a future normative specification.

## Agentic commands

Repository documentation MAY describe IASI agentic commands separately from the CLI.

The first canonical agentic command is:

```text
/validate
```

Its canonical definition lives in:

```text
iasi/commands/validate.md
```

Platform-specific representations, such as a Copilot prompt file under `.github/prompts/`, are adapters or generated integration artifacts and MUST NOT be presented as the canonical command definition.

## Acceptance criteria

Repository documentation is aligned when:

1. `README.md` and `README_en.md` describe `iasi/` as the canonical IASI source root;
2. neither README presents `agentics/` as the current architecture;
3. CLI examples match the current public CLI;
4. obsolete workspace/project scope flags are absent from current usage instructions;
5. agentic commands, when documented, point to `iasi/commands/` as their canonical definitions;
6. documentation does not create a second architectural contract that conflicts with the normative inputs.
