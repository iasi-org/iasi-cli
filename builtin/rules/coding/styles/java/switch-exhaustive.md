# Exhaustive switch

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.switch-exhaustive |
| Description | Requires Java switch logic to cover all possible inputs. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; switch; control-flow |

## Rule

Every `switch` statement or expression must be exhaustive. It must cover all possible values explicitly or provide an appropriate `default` branch.

## Exceptions

None.

## Rationale

Exhaustive switches make unhandled states explicit and prevent silent behavior when the domain evolves.

## Sources

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```java
return switch (status) {
    case NEW -> create();
    case ACTIVE -> update();
    case CLOSED -> archive();
};
```
