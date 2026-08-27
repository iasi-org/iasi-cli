# Avoid Boolean literal comparisons

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.boolean-comparisons |
| Description | Uses truth-value testing instead of explicit True or False comparisons. |
| Scope | coding/python |
| Level | should-not |
| Tags | coding; python; style; comparisons; boolean |

## Rule

Boolean expressions should not normally be compared explicitly with `True` or `False`.

## Exceptions

An API that distinguishes Boolean objects from other truthy or falsy values may require an explicit identity check.

## Rationale

Direct truth-value testing is idiomatic and removes redundant syntax.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred: `if enabled:`

Avoid: `if enabled == True:`
