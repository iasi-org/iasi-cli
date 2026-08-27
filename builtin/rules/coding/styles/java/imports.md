# Java imports

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.imports |
| Description | Keeps Java imports explicit and minimal. |
| Scope | coding/java |
| Level | should |
| Tags | coding; java; style; imports |

## Rule

Java imports should be explicit and limited to types actually used by the compilation unit. Wildcard imports should be avoided.

## Exceptions

A wildcard import may be used when required by generated code or an established repository convention.

## Rationale

Explicit imports make dependencies visible and avoid ambiguity when packages evolve.

## Sources

IASI

## Examples

None.
