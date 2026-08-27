# UTF-8 Python source

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.utf8-source |
| Description | Uses UTF-8 for Python source files. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; encoding; utf8 |

## Rule

Python source files must use UTF-8. An encoding declaration should not be added for ordinary UTF-8 source files.

## Exceptions

Tests explicitly exercising alternate encodings may use the encoding required by the test.

## Rationale

UTF-8 is Python 3's default source encoding and avoids unnecessary encoding declarations.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

None.
