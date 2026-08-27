# Do not invent dunder names

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.dunder-names-reserved |
| Description | Reserves double-underscore names for documented Python protocols. |
| Scope | coding/python |
| Level | must-not |
| Tags | coding; python; style; naming; dunder; api |

## Rule

Application code must not invent names with both leading and trailing double underscores. Such names must only be used when implementing documented Python protocols.

## Exceptions

None.

## Rationale

Dunder names are reserved for language and ecosystem protocols; inventing them risks collisions and misleading semantics.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Implement `__iter__` when implementing the iterator protocol; do not invent names such as `__process_data__`.
