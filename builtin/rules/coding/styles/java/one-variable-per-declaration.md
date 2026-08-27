# One variable per declaration

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.one-variable-per-declaration |
| Description | Declares one Java variable per declaration statement. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; variables; declarations |

## Rule

Each field or local-variable declaration must declare only one variable.

## Exceptions

Multiple variables may appear in the control portion of a `for` loop when appropriate.

## Rationale

One variable per declaration keeps types, initialization, comments, and later edits unambiguous.

## Sources

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```java
int width;
int height;
```

Not allowed:

```java
int width, height;
```
