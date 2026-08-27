# Python maximum line length

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.maximum-line-length |
| Description | Sets the default maximum width for Python source lines. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; formatting; line-length |

## Rule

Python code lines must not exceed 99 characters. Comments and docstrings should not exceed 72 characters.

## Exceptions

Long URLs, generated content, machine identifiers, or literals that become less readable when split may exceed the limit.

## Rationale

PEP 8 permits teams to extend code lines to 99 characters while retaining a narrower limit for prose. This Python-specific rule overrides the generic IASI 150-character limit.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Break long expressions using implicit continuation inside parentheses, brackets, or braces rather than extending the line.
