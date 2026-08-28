# IASI external inputs

## Normative status

The documents in this package describe the **current target behaviour** of IASI.

They are normative inputs for implementation.

The existing codebase is implementation state, not architectural authority. If the current implementation conflicts with these inputs, the implementation MUST be changed to satisfy these inputs.

Do not preserve superseded behaviour merely because it already exists in code.

## Package application

This package is a **replacement** for the previous `inputs/externals/` tree.

Do not overlay it onto an older external-input tree while preserving files that are absent from this package.

Before applying it, replace or remove:

```text
inputs/externals/
```

and then copy this package's complete `inputs/externals/` tree.

Repository-level documentation such as `README.md` and `README_en.md` is not bundled here, but it MUST be synchronized with the normative architecture as specified by `iasi-cli/06-repository-documentation.md`.

## Current public CLI

```text
iasi install
iasi reinstall
iasi status
iasi version
iasi adapt copilot
```

## Current canonical agentic commands

```text
/validate
/plan
/execute
/verify
/archive
```

Canonical definitions live under:

```text
iasi/commands/
```

Agentic commands are not CLI subcommands.

## Normative input groups

```text
inputs/externals/
├── iasi-cli/
├── adapters/
├── iasi-workflow/
├── command_validate/
├── command_plan/
├── command_execute/
├── command_verify/
├── command_archive/
└── skill_define/
```

The nine groups form one normative contract:

- `iasi-cli/` defines distribution, installation, resolution and common architecture.
- `adapters/` defines platform projection, currently GitHub Copilot.
- `iasi-workflow/` defines shared workflow state, gates and runtime services.
- `command_validate/` defines `/validate`.
- `command_plan/` defines `/plan`.
- `command_execute/` defines `/execute`.
- `command_verify/` defines `/verify`.
- `command_archive/` defines `/archive`.
- `skill_define/` defines the normative semantic contract for the `define` skill.

There is no historical-precedence chain between documents in this package. They MUST be mutually consistent.

## Implementation order

Implement in this order because later behaviour depends on earlier runtime contracts:

```text
1. canonical iasi/ source layout and embedded distribution
2. local install/reinstall semantics
3. composed .iasi resolver
4. status over the effective context
5. shared IASI runtime:
   - exact active-input discovery
   - deterministic fingerprints
   - workflow.json
   - gate evaluation
   - atomic state writes
   - safe filesystem operations
6. canonical /validate with pre-plan and post-plan validation modes
7. canonical /plan and its atomic previous-plan archive
8. canonical /execute
9. canonical /verify
10. canonical /archive
11. Copilot projection of all effective canonical commands
12. end-to-end workflow and recovery tests
```

Do not claim a command is implemented merely because its canonical Markdown file or Copilot prompt exists.

A command is implemented only when its required shared runtime behaviour, gates, state transitions, filesystem effects and tests are also implemented.
