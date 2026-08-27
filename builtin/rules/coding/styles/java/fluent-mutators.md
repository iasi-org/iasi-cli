# Fluent mutators

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.fluent-mutators |
| Description | Prefers returning the current object from mutating instance methods when no more meaningful result exists. |
| Scope | coding/java |
| Level | should |
| Tags | coding; style; java; methods; fluent-api; return-value |

## Rule

A mutating Java instance method that would otherwise return `void` should return the current instance when doing so is compatible with the public contract and provides a useful result.

## Exceptions

Do not change a required interface, framework callback, overridden signature, or established public contract solely to make a method fluent.

## Rationale

Returning the current object preserves composability without inventing an artificial result value.

## Sources

IASI

## Examples

Preferred:

```java
Widget setValue(int value) {
    this.value = value;
    return this;
}
```
