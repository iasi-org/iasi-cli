# Java package names

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.package-names |
| Description | Keeps Java package names lowercase and conventional. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; naming; packages |

## Rule

Package names must use lowercase letters and digits. CamelCase and underscores must not be used except when required to legalize an otherwise invalid package component.

## Exceptions

None.

## Rationale

Lowercase package names prevent visual ambiguity with type names and follow Java platform conventions.

## Sources

Java Language Specification, Chapter 6 - Names: https://docs.oracle.com/javase/specs/jls/se24/html/jls-6.html

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```java
package org.example.billing;
```

Not preferred:

```java
package org.example.BillingCore;
```
