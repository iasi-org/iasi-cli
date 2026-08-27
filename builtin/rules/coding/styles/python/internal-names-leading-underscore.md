# Leading underscore for internal names

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.internal-names-leading-underscore |
| Description | Marks intentionally non-public names with a leading underscore. |
| Scope | coding/python |
| Level | should |
| Tags | coding; python; style; naming; api; internal |

## Rule

Module-level names and attributes intended for internal use should begin with a single leading underscore when the public/private distinction is useful.

## Exceptions

Purely local implementation variables do not need an underscore solely because they are internal to a function.

## Rationale

The leading underscore is Python's conventional weak indicator of non-public API.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Example: `_load_internal_state()` for a module helper that is not part of the public API.
