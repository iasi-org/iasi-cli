# Python module and package names

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.module-package-names |
| Description | Keeps Python module and package names short and lowercase. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; naming; modules; packages |

## Rule

Module names must be short and lowercase; underscores may be used when they improve readability. Package names must be short and lowercase, and underscores should be avoided unless needed for clarity or compatibility.

## Exceptions

Existing package names and third-party integration boundaries may preserve established naming.

## Rationale

Predictable module and package names match Python ecosystem conventions and improve import readability.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred: `order_service.py`, package `billing`.

Avoid: `OrderService.py`, package `Billing_System`.
