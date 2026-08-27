# Java type names

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.type-names |
| Description | Uses descriptive noun-oriented names for Java types. |
| Scope | coding/java |
| Level | should |
| Tags | coding; java; style; naming; classes; interfaces; records |

## Rule

Class and record names should be descriptive nouns or noun phrases in UpperCamelCase. Interface names should be descriptive nouns, noun phrases, or adjectives that express their role.

## Exceptions

None.

## Rationale

Type names should communicate what an abstraction represents rather than how it happens to be implemented.

## Sources

Java Language Specification, Chapter 6 - Names: https://docs.oracle.com/javase/specs/jls/se24/html/jls-6.html

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```java
class PaymentRequest {}
interface Readable {}
record CustomerAddress(String city) {}
```
