# Resource management

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.resources |
| Description | Uses deterministic resource management in Java. |
| Scope | coding/java |
| Level | must |
| Tags | coding; java; style; resources; try-with-resources |

## Rule

A resource implementing `AutoCloseable` must be managed with try-with-resources when its ownership and lifetime are local to the operation.

## Exceptions

Another lifecycle mechanism may be used when resource ownership intentionally extends beyond the local operation.

## Rationale

Try-with-resources guarantees deterministic closing and handles exception propagation correctly.

## Sources

Oracle Java Tutorials - The try-with-resources Statement: https://docs.oracle.com/javase/tutorial/essential/exceptions/tryResourceClose.html

## Examples

Preferred:

```java
try (BufferedReader reader = Files.newBufferedReader(path)) {
    return reader.readLine();
}
```
