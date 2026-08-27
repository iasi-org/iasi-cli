# UTF-8 Java source files

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.utf8-source-files |
| Description | Uses UTF-8 for Java source code. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; encoding; source-files |

## Rule

Java source files must be encoded as UTF-8.

## Exceptions

None.

## Rationale

A single explicit encoding avoids platform-dependent interpretation of source files.

## Sources

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

A source file containing `España`, `μs`, or other non-ASCII text is stored directly as UTF-8.
