# Java exceptions remain visible

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.exceptions |
| Description | Prevents Java exceptions from being silently discarded. |
| Scope | coding/java |
| Level | must-not |
| Tags | coding; java; style; exceptions; errors |

## Rule

Java code must not catch an exception and silently discard it. A caught exception must be handled, translated with preserved context, propagated, or explicitly documented as intentionally ignorable.

## Exceptions

None.

## Rationale

Silent exception handling destroys failure information and can leave the application in an invalid state.

## Sources

IASI

## Examples

None.
