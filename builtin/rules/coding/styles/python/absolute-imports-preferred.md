# Absolute imports preferred

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.absolute-imports-preferred |
| Description | Prefers explicit absolute imports. |
| Scope | coding/python |
| Level | should |
| Tags | coding; python; style; imports |

## Rule

Absolute imports should be preferred when they remain clear and reasonably concise.

## Exceptions

Explicit relative imports may be used within packages when they make package-local relationships clearer or avoid unnecessary verbosity.

## Rationale

Absolute imports are generally easier to understand and produce clearer failures when package resolution is wrong.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

None.
