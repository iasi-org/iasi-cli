# Keep comprehensions simple

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.comprehensions-remain-simple |
| Description | Uses comprehensions only when they remain easy to read. |
| Scope | coding/python |
| Level | should |
| Tags | coding; python; style; comprehensions; readability |

## Rule

Comprehensions should be used for simple transformations and filtering. Complex nested logic, multiple conceptual operations, or substantial side effects should be extracted into named functions or explicit flow.

## Exceptions

A nested comprehension may be used when its structure is immediately obvious and simpler than the equivalent loop.

## Rationale

Python comprehensions are compact, but excessive density hides flow and conflicts with IASI flow-function principles.

## Sources

IASI; PEP 8 readability principles

## Examples

Preferred:

```python
active_names = [user.name for user in users if user.active]
```

Extract complex processing into a named function instead of embedding it in a comprehension.
