# Triple double-quoted docstrings

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.docstring-triple-double-quotes |
| Description | Uses the standard quoting form for docstrings. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; documentation; docstrings; quotes |

## Rule

Docstrings must use triple double quotes.

## Exceptions

Raw triple double quotes may be used when a docstring contains significant backslashes.

## Rationale

A single quoting convention improves consistency and matches PEP 257.

## Sources

PEP 257 - Docstring Conventions: https://peps.python.org/pep-0257/

## Examples

Preferred:

```python
def load():
    """Load and return the configuration."""
```
