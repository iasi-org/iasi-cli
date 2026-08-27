# One import per line

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.imports-one-per-line |
| Description | Keeps direct imports on separate lines. |
| Scope | coding/python |
| Level | must |
| Tags | coding; python; style; imports |

## Rule

Direct module imports must normally be written one per line.

## Exceptions

Multiple names may be imported from the same module in a single `from ... import ...` statement.

## Rationale

One module per import line improves scanning, editing, and diffs.

## Sources

PEP 8 - Style Guide for Python Code: https://peps.python.org/pep-0008/

## Examples

Preferred:

```python
import os
import sys
from subprocess import PIPE, Popen
```

Avoid: `import os, sys`.
