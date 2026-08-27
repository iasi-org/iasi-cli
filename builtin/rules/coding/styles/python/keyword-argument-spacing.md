# Keyword argument spacing

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.keyword-argument-spacing |
| Description | Uses Python spacing conventions for keyword arguments and defaults. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; whitespace; functions |

## Rule

Do not place spaces around `=` in keyword arguments or unannotated default parameter values.

## Exceptions

Annotated parameters with defaults must use spaces around `=` as prescribed by PEP 8.

## Rationale

This convention visually distinguishes argument binding from ordinary assignment.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred:

```python
def connect(timeout=30):
    return open_connection(retries=3)
```
