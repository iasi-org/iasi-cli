# Composed `.iasi` resolution

## Purpose

IASI uses **composed inheritance**.

A local installation complements installations found in ancestor directories. Resolution MUST NOT stop at the first `.iasi` found while walking upward.

## Layer discovery

Starting at the current working directory, inspect the current directory and every ancestor up to the filesystem root.

A directory contributes an installed layer only when this file exists and is valid:

```text
<directory>/.iasi/manifest.yml
```

Collect every valid layer.

For composition, order them from least specific to most specific:

```text
filesystem root
      ↓
parent installation
      ↓
child installation
      ↓
nearest installation
      ↓
current working directory
```

The nearest layer has the highest precedence.

If no installation layer exists, commands that require installed IASI MUST fail.

## Fundamental merge rule

Composition is additive with deterministic override:

```text
parent provides baseline
child adds new artifacts
child overrides collisions
```

A child artifact never causes the resolver to ignore unrelated parent artifacts.

There is currently no tombstone or explicit removal mechanism. A child may add or replace an artifact, but cannot hide an inherited artifact without providing its replacement.

## Instructions

Instruction identity is the declared instruction `id`.

Within a single layer:

- instruction IDs MUST be unique;
- duplicate IDs are invalid and MUST fail resolution.

Across different layers:

- the instruction with the same `id` from the nearest layer wins;
- different IDs are combined.

The relative file path is not the precedence key when an instruction has a valid semantic ID.

## Commands

Command identity is its normalized relative path under `commands/` without the final `.md` extension.

Examples:

```text
commands/validate.md       → validate
commands/review/api.md     → review/api
```

Across layers, the nearest command with the same identity replaces the parent command completely.

Different command identities are combined.

## Adapters

Adapter identity is the adapter `id` declared by its `adapter.yml` and represented by its top-level directory under:

```text
adapters/<id>/
```

The reserved `adapters/schema/` directory is metadata and is not an adapter.

Adapters are **atomic units**. They are never merged file by file.

If the child layer contains:

```text
adapters/copilot/
```

that adapter completely replaces an inherited `copilot` adapter.

If the child layer does not contain `copilot`, the nearest ancestor `copilot` adapter is inherited.

If the nearest selected adapter is invalid, resolution MUST fail. The resolver MUST NOT silently fall back to an older parent adapter.

## Skills

Until a richer skill schema is introduced, skill identity is the first path component under:

```text
skills/<id>/
```

A skill is an atomic subtree. The nearest layer containing a skill with the same `<id>` replaces the parent skill completely.

## MCP

Until a richer MCP schema is introduced, MCP identity is the first path component under:

```text
mcp/<id>/
```

An MCP artifact is an atomic subtree. The nearest layer containing the same `<id>` replaces the parent artifact completely.

## Schema and support content

Reserved schema/readme support content is not a runtime semantic artifact. When tooling requires such content, use the nearest available support subtree for that category rather than merging support files across layers.

## Effective context

The result of layer composition is the **effective IASI context** for the current location.

Commands that consume installed IASI, including:

```text
iasi status
iasi adapt copilot
```

MUST use this effective context rather than a single nearest installation.

## Determinism

Given the same ordered layers and file contents, effective resolution MUST produce the same artifact set regardless of filesystem enumeration order.

Artifact discovery and output ordering MUST therefore be explicitly sorted.
