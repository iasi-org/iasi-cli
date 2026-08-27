# Consistent abstraction level

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.abstraction-level |
| Description | Keeps statements in a function at a comparable level of abstraction. |
| Scope | coding |
| Level | should |
| Tags | coding; style; functions; abstraction; flow; readability |

## Rule

Statements in a flow function should operate at a comparable level of abstraction. Low-level implementation details should be delegated instead of being mixed directly with high-level operations.

## Exceptions

A small local expression may remain inline when extracting it would obscure rather than clarify the flow.

## Rationale

Mixing abstraction levels forces the reader to alternate between understanding the process and understanding its implementation details.

## Sources

IASI

## Examples

Preferred:

```text
load_configuration();
validate_configuration();
build_plan();
execute_plan();
return result;
```

Avoid inserting low-level parsing or file-manipulation details between those flow operations.
