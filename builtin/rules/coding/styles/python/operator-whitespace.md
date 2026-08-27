# Operator whitespace

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.operator-whitespace |
| Description | Uses consistent whitespace around Python operators. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; whitespace; operators |

## Rule

Binary comparison, assignment, augmented-assignment, and Boolean operators must normally have a single space on each side. Extra whitespace used only to align neighboring statements must not be added.

## Exceptions

PEP 8 permits selective omission of spaces around higher-precedence arithmetic operators when that improves readability.

## Rationale

Consistent spacing makes expressions easier to parse visually without decorative alignment.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred:

```python
count = count + 1
ready = enabled and count > 0
```
