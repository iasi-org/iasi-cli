# Catch specific exceptions

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.specific-exceptions |
| Description | Avoids overly broad exception handlers. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; exceptions; errors |

## Rule

Exception handlers must catch the narrowest practical exception type. Bare `except:` handlers must not be used in ordinary application code.

## Exceptions

A top-level process boundary that must perform cleanup or logging before re-raising or terminating may deliberately catch broadly.

## Rationale

Specific handlers prevent unrelated programming errors from being silently mistaken for expected failures.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred:

```python
try:
    value = int(text)
except ValueError:
    return None
```
