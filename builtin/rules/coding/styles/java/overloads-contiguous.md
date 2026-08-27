# Keep overloads contiguous

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.overloads-contiguous |
| Description | Keeps overloaded Java methods and constructors together. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; methods; constructors; organization |

## Rule

Methods with the same name must appear in one contiguous group. Multiple constructors must likewise remain contiguous.

## Exceptions

None.

## Rationale

Keeping related overloads together makes the available variants visible as a single API concept.

## Sources

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```java
Result load(Path path) { ... }
Result load(Path path, Charset charset) { ... }
Result load(InputStream input) { ... }
```
