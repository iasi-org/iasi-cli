# One public top-level type per file

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.one-public-type-per-file |
| Description | Keeps each Java compilation unit centered on one public top-level type. |
| Scope | coding/java |
| Level | should |
| Tags | coding; java; style; files; types |

## Rule

A Java source file should contain at most one public top-level class, interface, record, enum, or annotation, and the filename should match that public type.

## Exceptions

None.

## Rationale

A one-to-one relationship between public types and source files keeps navigation and ownership predictable.

## Sources

IASI

## Examples

None.
