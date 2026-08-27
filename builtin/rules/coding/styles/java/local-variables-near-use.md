# Declare local variables near use

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.local-variables-near-use |
| Description | Keeps Java local variables close to where they are first needed. |
| Scope | coding/java |
| Level | should |
| Tags | coding; java; style; variables; scope |

## Rule

Local variables should be declared close to their first use and should have the smallest practical scope.

## Exceptions

None.

## Rationale

Shorter variable lifetimes reduce the amount of state a reader must track and make accidental reuse less likely.

## Sources

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```java
validate(input);
Result result = process(input);
return result;
```

Avoid declaring `result` many lines before it is used.
