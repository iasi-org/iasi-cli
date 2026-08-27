# No wildcard imports

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.no-wildcard-imports |
| Description | Keeps Java imports explicit. |
| Scope | coding/java |
| Level | must-not |
| Tags | coding; java; style; imports |

## Rule

Java source files must not use wildcard imports, including static wildcard imports.

## Exceptions

Generated source code may use wildcard imports when the generator cannot reasonably be changed.

## Rationale

Explicit imports keep dependencies visible and avoid ambiguity when packages evolve.

## Sources

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```java
import java.util.List;
import java.util.Map;
```

Not allowed:

```java
import java.util.*;
```
