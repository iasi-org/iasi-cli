# Focused changes

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.focused-change |
| Description | Keeps a change limited to the behavior it intends to modify. |
| Scope | coding |
| Level | must |
| Tags | coding; style; changes; refactoring; scope |

## Rule

A focused change must not include unrelated refactoring, formatting churn, renaming, or behavioral changes.

## Exceptions

None.

## Rationale

Keeping changes focused reduces review noise and makes causality, rollback, and validation clearer.

## Sources

IASI

## Examples

None.
