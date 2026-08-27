# Raise exceptions for exceptional failures

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.raise-not-return-error |
| Description | Uses Python exceptions rather than sentinel error codes for exceptional failures. |
| Scope | coding/python |
| Level | should |
| Tags | coding; python; style; exceptions; errors; api |

## Rule

Exceptional failures should normally be represented with exceptions rather than magic return codes such as `-1` or error strings.

## Exceptions

Protocols, performance-sensitive APIs, or external interfaces that explicitly define sentinel return values may use them.

## Rationale

Exceptions separate failure control flow from valid return values and compose naturally with Python APIs.

## Sources

IASI; Python language conventions

## Examples

None.
