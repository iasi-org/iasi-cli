# Logical class member order

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.java.logical-member-order |
| Description | Requires Java class members to follow an explainable logical order. |
| Scope | coding/java |
| Level | should |
| Tags | coding; java; style; classes; organization |

## Rule

Members within a Java type should be arranged in a logical order that reflects the structure or use of the type. New members should not be appended mechanically to the end of the file.

## Exceptions

None.

## Rationale

Logical ordering improves learnability and keeps the source structure meaningful as the class evolves.

## Sources

Google Java Style Guide: https://google.github.io/styleguide/javaguide.html

## Examples

Examples of valid ordering strategies include public API before implementation helpers, lifecycle order, or grouping by responsibility. The chosen strategy should be consistent within the type.
