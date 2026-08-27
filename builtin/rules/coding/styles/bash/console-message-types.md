# Console messages have explicit types

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.console-message-types |
| Description | Classifies console output with explicit semantic message types. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; console; logging; messages |

## Rule

Operational console messages must identify their semantic type explicitly.

The standard message types are `INFO`, `SUCCESS`, `WARN`, `ERROR`, and `DEBUG`. The type label must be present in text and must not be communicated only through color.

## Exceptions

A script whose output is intentionally machine-readable must not inject human-oriented message labels into that output stream.

## Rationale

Explicit message types let humans, logs, and automation distinguish normal progress, successful completion, warnings, errors, and diagnostic output without relying on presentation.

## Sources

IASI

## Examples

Preferred:

```text
[2026-08-27T20:45:12+0200] [INFO] Reading configuration
[2026-08-27T20:45:13+0200] [SUCCESS] Configuration loaded
[2026-08-27T20:45:13+0200] [WARN] Optional file not found
[2026-08-27T20:45:14+0200] [ERROR] Validation failed
```

Not preferred:

```text
Reading configuration
Done
Something happened
```
