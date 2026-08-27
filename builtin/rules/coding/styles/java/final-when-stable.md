# Final for stable references

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.final-when-stable |
| Description | Prefers final fields when their reference does not change after initialization. |
| Scope | coding/java |
| Level | should |
| Tags | coding; java; style; final; immutability |

## Rule

Instance and class fields whose reference is not expected to change after initialization should be declared `final`.

## Exceptions

None.

## Rationale

Stable references reduce mutable state and communicate lifecycle intent directly in the type.

## Sources

IASI

## Examples

None.
