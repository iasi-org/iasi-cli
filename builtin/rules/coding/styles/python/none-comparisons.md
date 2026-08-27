# Compare None by identity

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.none-comparisons |
| Description | Uses identity operators for None. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; comparisons; none |

## Rule

Comparisons with `None` must use `is None` or `is not None`, never equality operators.

## Exceptions

None.

## Rationale

`None` is a singleton sentinel and identity expresses the intended test directly.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred: `if value is None:`

Avoid: `if value == None:`
