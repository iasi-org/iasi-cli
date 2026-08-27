# Maximum nesting depth

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.maximum-nesting-depth |
| Description | Limits nested control flow to three levels. |
| Scope | coding |
| Level | must |
| Tags | coding; style; control-flow; nesting; complexity; readability |

## Rule

Control structures must not exceed three nested levels inside a function or method. The function body itself is not counted as a nesting level.

## Exceptions

A language construct or generated source may require deeper structural nesting when decomposition would not improve clarity.

## Rationale

Deep nesting hides the main flow and increases the number of conditions a reader must keep in working memory.

## Sources

IASI

## Examples

If a fourth nested `if`, `for`, `while`, or equivalent construct is required, extract the nested operation or simplify the preceding control flow.
