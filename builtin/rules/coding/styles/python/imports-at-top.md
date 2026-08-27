# Imports at the top

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.imports-at-top |
| Description | Places imports in a predictable location. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; imports |

## Rule

Imports must normally appear at the top of the file, after the module docstring and before module globals and constants.

## Exceptions

A local import may be used to break a dependency cycle, avoid an optional dependency at startup, or defer an expensive import when there is a justified reason.

## Rationale

Centralized imports make dependencies immediately visible.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

None.
