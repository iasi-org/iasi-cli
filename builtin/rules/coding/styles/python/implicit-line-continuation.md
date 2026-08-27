# Implicit line continuation

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.implicit-line-continuation |
| Description | Uses Python delimiters rather than backslashes to wrap expressions. |
| Scope | coding/python |
| Level | should |
| Tags | coding; python; style; formatting; line-continuation |

## Rule

Long Python expressions should be wrapped using implicit continuation inside parentheses, brackets, or braces instead of backslash line continuation.

## Exceptions

A backslash may be used when the syntax offers no clearer implicit-continuation form.

## Rationale

Implicit continuation is safer to edit and preserves expression structure visibly.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred:

```python
result = calculate_total(
    subtotal,
    taxes,
    discount,
)
```
