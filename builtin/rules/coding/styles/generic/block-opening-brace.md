# Opening brace placement

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.block-opening-brace |
| Description | Places opening braces on the line that starts the block. |
| Scope | coding |
| Level | must |
| Tags | coding; style; formatting; blocks; braces |

## Rule

When a language uses `{` to open a block, the opening brace must be placed on the same line as the construct that starts the block.

## Exceptions

None.

## Rationale

A consistent opening-brace position keeps the start of a block visually attached to the construct that owns it.

## Sources

IASI

## Examples

Preferred:

```text
if (condition) {
    action();
}
```

Not preferred:

```text
if (condition)
{
    action();
}
```
