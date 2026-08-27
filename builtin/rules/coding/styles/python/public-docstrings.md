# Public API docstrings

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.public-docstrings |
| Description | Documents public Python modules, functions, classes, and methods. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; documentation; docstrings; api |

## Rule

Public modules, functions, classes, methods, and constructors must normally have docstrings that describe their externally relevant behavior.

## Exceptions

Trivial overrides whose behavior is completely inherited may omit duplicated documentation when the inherited contract is clear.

## Rationale

Docstrings are Python's native API documentation mechanism and are available at runtime through `__doc__`.

## Sources

PEP 257 - Docstring Conventions: https://peps.python.org/pep-0257/

## Examples

None.
