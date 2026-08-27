# Guard clauses

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.guard-clauses |
| Description | Uses early exits to keep the main execution path shallow and visible. |
| Scope | coding |
| Level | should |
| Tags | coding; style; control-flow; guard-clause; nesting; readability |

## Rule

When an exceptional, invalid, or terminating condition can be handled immediately, prefer a guard clause or early exit instead of nesting the remaining function body inside another control block.

## Exceptions

Do not introduce early exits when the language, resource-management model, or local control flow becomes less clear as a result.

## Rationale

Guard clauses remove unnecessary nesting and leave the normal execution path visible.

## Sources

IASI

## Examples

Preferred:

```text
if (!valid(input)) return error;
process(input);
return result;
```

Instead of:

```text
if (valid(input)) {
    process(input);
    return result;
}
return error;
```
