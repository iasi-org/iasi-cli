# Java method names

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.method-names |
| Description | Uses verb-oriented names for Java methods. |
| Scope | coding/java |
| Level | should |
| Tags | coding; java; style; naming; methods |

## Rule

Method names should be verbs or verb phrases written in lowerCamelCase and should describe the operation performed or result obtained.

## Exceptions

None.

## Rationale

Verb-oriented method names make calls read as operations and improve the readability of flow functions.

## Sources

Java Language Specification, Chapter 6 - Names: https://docs.oracle.com/javase/specs/jls/se24/html/jls-6.html

Oracle Java Code Conventions - Naming Conventions: https://www.oracle.com/java/technologies/javase/codeconventions-namingconventions.html

## Examples

Preferred:

```java
loadConfiguration();
validateInputs();
createSpecifications();
```
