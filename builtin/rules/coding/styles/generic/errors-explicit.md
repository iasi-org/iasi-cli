# Errors remain explicit

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.errors-explicit |
| Description | Prevents failures from being silently hidden. |
| Scope | coding |
| Level | must |
| Tags | coding; style; errors; failure; diagnostics |

## Rule

Failure modes must remain explicit. Errors must not be silently ignored, broadly swallowed, or converted into successful outcomes without a documented and justified fallback.

## Exceptions

None.

## Rationale

Visible failures preserve causality and make incorrect states diagnosable.

## Sources

IASI

## Examples

None.
