# No else after terminal branch

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.terminal-branch-no-else |
| Description | Avoids unnecessary else blocks after a branch already terminates control flow. |
| Scope | coding |
| Level | should |
| Tags | coding; style; control-flow; else; return; readability |

## Rule

When an `if` branch terminates the current flow with `return`, `throw`, `break`, `continue`, or an equivalent operation, the subsequent path should continue after the conditional instead of being wrapped in an unnecessary `else`.

## Exceptions

Use `else` when the explicit symmetry of the alternatives materially improves understanding.

## Rationale

Removing redundant `else` blocks reduces nesting and keeps the continuing execution path at the shallowest indentation level.

## Sources

IASI

## Examples

Preferred:

```text
if (!valid) return error;
process();
return result;
```

Instead of:

```text
if (!valid) return error;
else {
    process();
    return result;
}
```
