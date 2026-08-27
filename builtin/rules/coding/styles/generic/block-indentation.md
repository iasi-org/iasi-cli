# Block indentation

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.block-indentation |
| Description | Aligns the contents of parenthesized control blocks from the opening parenthesis. |
| Scope | coding |
| Level | must |
| Tags | coding; style; formatting; indentation; blocks |

## Rule

For control structures such as `if`, `while`, `for`, or equivalent constructs that use parentheses, the contents of the block must begin at the indentation column established by the opening parenthesis of the control structure.

## Exceptions

A language formatter or mandatory repository convention may define a different indentation scheme.

## Rationale

Indenting from the opening parenthesis makes the visual hierarchy of the block depend on the syntactic construct that introduces it rather than on an arbitrary fixed indentation width.

## Sources

IASI

## Examples

```text
if (condition) {
   action();
}

while (condition) {
      action();
}
```
