# Consistent return statements

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.consistent-return-statements |
| Description | Keeps return behavior explicit within a function. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; functions; return-value |

## Rule

When any return statement in a function returns an expression, all reachable return paths should return an explicit expression. Use `return None` when an explicit no-result path is required.

## Exceptions

A function that only performs side effects may use bare `return` or implicit `None` when this is required by its API contract; the generic IASI preference for meaningful return values still applies when practical.

## Rationale

Consistent return forms make a function contract easier to understand and analyze.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred:

```python
def find(value):
    if value:
        return value
    return None
```
