# Functions return a value

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.functions-return-value |
| Description | Prefers functions and methods that return a meaningful value. |
| Scope | coding |
| Level | must |
| Tags | coding; style; functions; methods; return-value |

## Rule

A function or method must return a meaningful value whenever the language and its external contract permit it. Procedure-style or `void` operations are avoided by default.

When a mutating instance method has no more meaningful result, returning the current object is preferred when appropriate.

## Exceptions

A function or method may return no value when required by the language, framework, callback signature, interface, interoperability contract, or another concrete constraint.

## Rationale

Returning values makes operations composable and avoids an unnecessary distinction between functions and procedures.

## Sources

IASI

## Examples

Preferred:

```text
Object setValue(int value) {
    this.value = value;
    return this;
}
```
