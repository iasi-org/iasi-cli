# Java constant names

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.constant-names |
| Description | Uses uppercase underscore-separated names for constants. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; naming; constants |

## Rule

Constants must use uppercase words separated by underscores.

## Exceptions

None.

## Rationale

The convention distinguishes constants immediately from fields and local variables.

## Sources

Java Language Specification, Chapter 6 - Names: https://docs.oracle.com/javase/specs/jls/se24/html/jls-6.html

Oracle Java Code Conventions - Naming Conventions: https://www.oracle.com/java/technologies/javase/codeconventions-namingconventions.html

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```java
static final int MAX_RETRIES = 3;
```

Not preferred:

```java
static final int maxRetries = 3;
```
