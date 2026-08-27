# Four-space indentation

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.four-space-indentation |
| Description | Uses four spaces for each Python indentation level. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; formatting; indentation |

## Rule

Each Python indentation level must use four spaces.

## Exceptions

Continuation lines may use alignment or hanging indentation when it clearly represents expression structure.

## Rationale

Four-space indentation is the standard Python layout and makes nesting predictable.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred:

```python
if ready:
    process()
```
