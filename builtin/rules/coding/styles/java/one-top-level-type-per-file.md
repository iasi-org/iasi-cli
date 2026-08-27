# One top-level type per file

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.one-top-level-type-per-file |
| Description | Keeps each Java source file centered on one top-level type. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; files; types |

## Rule

A normal Java source file must contain exactly one top-level type declaration, and the filename must match that type name.

## Exceptions

`package-info.java` and `module-info.java` follow their language-defined structures.

## Rationale

A one-to-one relationship between source files and top-level types keeps navigation, ownership, and tooling predictable.

## Sources

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```text
CustomerService.java  ->  class CustomerService
```
