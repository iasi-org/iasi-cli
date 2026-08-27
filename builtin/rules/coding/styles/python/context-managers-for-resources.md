# Context managers for resources

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.context-managers-for-resources |
| Description | Uses context managers for deterministic resource cleanup. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; resources; context-manager; cleanup |

## Rule

Resources that implement the context manager protocol, such as files, should be acquired with a `with` statement when deterministic cleanup is required.

## Exceptions

Manual lifecycle management may be used when the resource must intentionally outlive a lexical `with` block or an API does not expose context management.

## Rationale

Context managers make acquisition and cleanup explicit and reliable across normal and exceptional control flow.

## Sources

Python documentation - contextlib: https://docs.python.org/3/library/contextlib.html

## Examples

Preferred:

```python
with open(path, encoding="utf-8") as stream:
    content = stream.read()
```
