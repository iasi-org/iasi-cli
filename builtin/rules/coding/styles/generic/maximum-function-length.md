# Maximum function length

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.maximum-function-length |
| Description | Keeps functions and methods small enough to understand as a unit. |
| Scope | coding |
| Level | must |
| Tags | coding; style; functions; methods; readability; complexity |

## Rule

A function or method must not exceed 50 physical lines from its declaration to its closing delimiter. When it approaches this size, review it for decomposition into focused subfunctions.

## Exceptions

Generated code or a concrete language/framework constraint may justify exceeding the limit.

## Rationale

A function should normally fit within a single screen-sized reading context so its purpose and flow can be understood without navigation.

## Sources

IASI

## Examples

A long function that loads data, validates it, chooses a processing path, applies a result, and post-processes it should be decomposed into a flow function and focused operations.
