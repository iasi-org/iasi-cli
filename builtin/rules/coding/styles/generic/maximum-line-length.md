# Maximum line length

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.maximum-line-length |
| Description | Limits source lines to a readable maximum width. |
| Scope | coding |
| Level | must |
| Tags | coding; style; formatting; readability; line-length |

## Rule

A source-code line must not exceed 150 characters.

## Exceptions

Generated code, unavoidable machine identifiers, or literals whose splitting would reduce correctness or readability may exceed the limit.

## Rationale

A bounded line width keeps code readable without horizontal navigation and discourages statements that carry too much information at once.

## Sources

IASI

## Examples

When an expression approaches the limit, split it according to the language syntax or extract a named intermediate operation.
