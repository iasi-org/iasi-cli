# No wildcard imports

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.no-wildcard-imports |
| Description | Avoids importing an unknown set of names. |
| Scope | coding/python |
| Level | must-not |
| Tags | coding; python; style; imports; namespace |

## Rule

Production Python code must not use `from module import *`.

## Exceptions

A module explicitly designed to republish an internal interface as a public API may use a wildcard import when the exported set cannot reasonably be stated explicitly.

## Rationale

Wildcard imports hide namespace contents and make static analysis and maintenance harder.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Avoid:

```python
from service import *
```
