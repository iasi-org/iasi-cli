# `/validate` platform integration

## Canonical ownership

IASI owns `/validate` at:

```text
iasi/commands/validate.md
```

The command-specific implementation requirements are these inputs:

```text
inputs/externals/command_validate/
```

There is no authoritative command definition under `.github/` or another platform-native directory.

## Installation

`iasi install` installs canonical commands into the local installed layer:

```text
.iasi/commands/
```

Composed resolution may inherit or override commands from parent layers as defined by `iasi-cli/03-resolution.md`.

## Copilot projection

The Copilot adapter supports IASI commands.

The effective command:

```text
commands/validate.md
```

is projected to:

```text
.github/prompts/validate.prompt.md
```

The generated Copilot prompt MUST remain a thin platform projection of the effective canonical command.

It MAY include Copilot-specific invocation framing required by the platform, but it MUST NOT independently redefine:

- what inputs are validated;
- blocker/warning semantics;
- the workflow gate;
- validation state format;
- fingerprint/staleness rules.

## Portability

A future Codex or other adapter may expose `/validate` differently, but the canonical command and validation semantics remain unchanged.

Changing platform means changing projection, not redefining the command.

## Current implementation consequence

Any older Copilot-adapter behaviour that explicitly excludes commands or prompt files is superseded by this specification.

The current Copilot adapter MUST be extended so that `/validate` is exposed through its native prompt mechanism.
