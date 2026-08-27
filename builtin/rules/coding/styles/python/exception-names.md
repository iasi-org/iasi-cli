# Exception names end in Error

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.exception-names |
| Description | Names error exception classes consistently. |
| Scope | coding/python |
| Level | should |
| Tags | coding; python; style; naming; exceptions |

## Rule

Exception classes that represent errors should use CapWords and end with `Error`.

## Exceptions

Exception classes that represent non-error control conditions may use another descriptive name.

## Rationale

The suffix communicates the role of the class and matches Python conventions.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred: `class ConfigurationError(Exception): ...`
