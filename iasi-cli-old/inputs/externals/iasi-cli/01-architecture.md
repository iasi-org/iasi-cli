# IASI architecture

## Principle

IASI owns its behaviour independently of any specific agentic platform.

All canonical IASI artifacts MUST live under:

```text
iasi/
```

Target source structure:

```text
iasi/
├── instructions/
├── commands/
├── skills/
├── mcp/
└── adapters/
    ├── schema/
    └── copilot/
```

The previous `agentics/...` source layout is superseded. New implementation work MUST use the structure above.

## Ownership

```text
iasi/          canonical IASI behaviour and platform adapters
inputs/        project knowledge and implementation inputs
.iasi/         installed IASI layer and local IASI runtime state
.github/       Copilot-native generated projection
```

Platform-native generated files are not a second source of truth.

## Commands

Canonical IASI agentic commands live in:

```text
iasi/commands/
```

The current command set is:

```text
iasi/commands/validate.md
iasi/commands/plan.md
iasi/commands/execute.md
iasi/commands/verify.md
iasi/commands/archive.md
```

which exposes:

```text
/validate
/plan
/execute
/verify
/archive
```

Agentic commands are not CLI subcommands. The CLI installs and adapts command definitions while the agentic platform invokes them through its native mechanism.

Command semantics and workflow gates belong to IASI, never to the platform adapter.

## Adapters

Canonical platform adapters live in:

```text
iasi/adapters/<platform>/
```

Adapters translate IASI artifacts into platform-native representations. They MUST NOT redefine IASI semantics.

Dependency direction:

```text
platform-native entry point
          ↓
       adapter
          ↓
 canonical IASI artifact
          ↓
     IASI workflow
```

## Distribution

The standalone executable contains the complete current IASI distribution required by `iasi install`:

```text
VERSION
+
iasi/
```

The source tree remains declarative data. Go code MUST NOT contain duplicated methodological rules.
