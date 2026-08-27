# Access static members through their class

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.static-members-through-class |
| Description | Makes static-member ownership explicit in Java code. |
| Scope | coding/java |
| Level | should |
| Tags | coding; java; style; static; readability |

## Rule

A static field or method should be accessed through its declaring class or directly when already in that class, not through an object instance.

## Exceptions

None.

## Rationale

Using an instance to access a static member falsely suggests instance-specific behavior.

## Sources

Oracle Java Code Conventions - Programming Practices: https://www.oracle.com/java/technologies/javase/codeconventions-programmingpractices.html

## Examples

Preferred:

```java
Clock.systemUTC();
```

Avoid:

```java
clock.systemUTC();
```
