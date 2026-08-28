# IASI adapter model

## Purpose

An adapter projects canonical IASI artifacts into a platform-native representation.

Adapters are part of IASI and live canonically under:

```text
iasi/adapters/
```

They contain platform mapping rules, not independent IASI methodology.

## Minimum adapter structure

```text
iasi/adapters/
├── schema/
│   └── adapter.md
└── copilot/
    ├── README.md
    └── adapter.yml
```

`adapters/schema/` is reserved support content and is not an adapter.

## Descriptor

Each adapter descriptor MUST contain:

```yaml
schema_version: 1
id: <adapter-id>
platform: <platform-id>
```

`schema_version` is mandatory and versions the descriptor format.

Do not add an independent adapter release `version`. Adapters ship with the IASI distribution version.

## Supported artifact types

The descriptor declares which canonical IASI artifact types it can project.

For the current Copilot adapter:

```yaml
supports:
  instructions: true
  commands: true
  skills: false
  mcp: false
  agents: false
```

Unsupported artifact types MUST remain unsupported. The adapter MUST NOT approximate them using another platform concept.

## Source resolution

`iasi adapt <adapter>` operates on the **effective composed IASI context** defined by `iasi-cli/03-resolution.md`.

The requested adapter is itself resolved by adapter ID using the same precedence rules.

An adapter found only in a parent layer remains available to a child project.

A nearer adapter with the same ID replaces the parent adapter as a complete atomic unit.

## Validity

A selected adapter is invalid when, among other failures:

- its descriptor is missing or malformed;
- `schema_version` is missing or unsupported;
- `id` or `platform` is missing;
- descriptor `id` does not match the requested adapter;
- required mappings for declared supported types are missing;
- a configured target escapes the platform-owned target boundary.

An invalid selected adapter MUST fail preflight. Do not fall back to an older inherited adapter.
