# Public contracts remain stable

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.public-contract-stability |
| Description | Protects public behavior from accidental changes. |
| Scope | coding |
| Level | must |
| Tags | coding; style; api; compatibility; contracts |

## Rule

Public interfaces and externally observable behavior must remain stable unless changing them is an explicit requirement of the task.

## Exceptions

None.

## Rationale

Unintended contract changes create downstream breakage that is often larger than the local implementation change.

## Sources

IASI

## Examples

None.
