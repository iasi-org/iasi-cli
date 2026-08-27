# Python naming conventions

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.naming |
| Description | Uses Python naming conventions consistently. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; naming |

## Rule

Functions, methods, local variables, parameters, and instance variables must use snake_case. Classes and exceptions must use CapWords. Constants must use UPPER_SNAKE_CASE.

## Exceptions

Existing external APIs may retain their established names when compatibility requires it.

## Rationale

Conventional casing makes identifier roles immediately recognizable to Python developers.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred:

```python
MAX_RETRIES = 3

class CustomerOrder:
    def calculate_total(self, item_count):
        return item_count * 10
```
