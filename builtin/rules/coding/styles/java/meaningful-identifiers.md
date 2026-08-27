# Meaningful Java identifiers

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.meaningful-identifiers |
| Description | Prefers meaningful identifiers over cryptic abbreviations. |
| Scope | coding/java |
| Level | should |
| Tags | coding; java; style; naming; readability |

## Rule

Identifiers should use clear, meaningful words. Cryptic abbreviations and one-character names should be avoided except for established idioms or very small local scopes where the meaning is immediately obvious.

## Exceptions

None.

## Rationale

Meaningful identifiers reduce the amount of surrounding code a reader must inspect to understand an operation.

## Sources

Java Language Specification, Chapter 6 - Names: https://docs.oracle.com/javase/specs/jls/se24/html/jls-6.html

Oracle Java Code Conventions - Naming Conventions: https://www.oracle.com/java/technologies/javase/codeconventions-namingconventions.html

## Examples

Preferred:

```java
int retryCount;
Customer customer;
```

Avoid:

```java
int rc;
Customer c;
```
