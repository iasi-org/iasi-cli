# Java CamelCase naming

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.camel-case |
| Description | Uses conventional CamelCase naming for Java identifiers. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; naming; camel-case |

## Rule

Java identifiers must follow conventional CamelCase naming. Type names use UpperCamelCase. Method names, field names, local variables, and parameters use lowerCamelCase.

## Exceptions

None.

## Rationale

Java casing conventions make the role of an identifier visible at a glance and align code with the Java ecosystem.

## Sources

Java Language Specification, Chapter 6 - Names: https://docs.oracle.com/javase/specs/jls/se24/html/jls-6.html

Oracle Java Code Conventions - Naming Conventions: https://www.oracle.com/java/technologies/javase/codeconventions-namingconventions.html

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Preferred:

```java
class CustomerOrder {
    private int itemCount;

    CustomerOrder addItem(Item item) {
        itemCount++;
        return this;
    }
}
```

Not preferred:

```java
class customer_order {
    private int Item_Count;
}
```
