# Resource management

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.resources |
| Description | Uses deterministic resource management in Java. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; resources; try-with-resources |

## Rule

Java resources implementing `AutoCloseable` must be managed with try-with-resources when their lifetime is local to the operation.

## Exceptions

Another lifecycle mechanism may be used when resource ownership intentionally extends beyond the local operation.

## Rationale

Try-with-resources makes cleanup deterministic and preserves exception semantics.

## Sources

IASI

## Examples

None.
