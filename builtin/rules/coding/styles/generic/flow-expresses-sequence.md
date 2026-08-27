# Flow expresses sequence

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.flow-expresses-sequence |
| Description | Keeps flow functions focused on orchestration rather than implementation details. |
| Scope | coding |
| Level | should |
| Tags | coding; style; flow; orchestration; responsibility |

## Rule

A function whose responsibility is flow or orchestration should express the sequence of operations explicitly and delegate the work of each operation to a focused function.

The flow should read close to pseudocode.

## Exceptions

None.

## Rationale

Separating sequence from implementation makes the process visible without forcing the reader to inspect operational details.

## Sources

IASI
