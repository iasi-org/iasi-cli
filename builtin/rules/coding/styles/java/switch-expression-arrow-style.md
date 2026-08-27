# Use arrow-style switch expressions

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.switch-expression-arrow-style |
| Description | Uses modern arrow syntax for Java switch expressions. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; switch; expressions |

## Rule

Java switch expressions must use arrow-style switch rules (`->`) rather than colon-style statement groups.

## Exceptions

None.

## Rationale

Arrow-style switch expressions avoid fall-through semantics and make each mapping explicit.

## Sources

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```java
return switch (size) {
    case 0 -> "";
    case 1 -> values.getFirst();
    default -> String.join(", ", values);
};
```
