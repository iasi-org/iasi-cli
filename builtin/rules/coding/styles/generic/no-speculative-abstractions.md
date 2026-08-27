# No speculative abstractions

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.no-speculative-abstractions |
| Description | Prevents abstractions created only for hypothetical future needs. |
| Scope | coding |
| Level | must-not |
| Tags | coding; style; abstractions; simplicity; yagni |

## Rule

Code must not introduce an abstraction solely for a hypothetical future requirement. An abstraction must solve a concrete current problem such as reuse, clarity, isolation, substitution, or ownership of responsibility.

## Exceptions

None.

## Rationale

Premature abstractions increase conceptual surface without providing present value.

## Sources

IASI

## Examples

None.
