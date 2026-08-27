# Array brackets belong to the type

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.array-brackets-on-type |
| Description | Uses Java type-style array declarations. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; arrays; declarations |

## Rule

Array brackets must be written with the type, not after the variable name.

## Exceptions

None.

## Rationale

The brackets are part of the Java type and should be presented as such.

## Sources

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```java
String[] args;
```

Not allowed:

```java
String args[];
```
