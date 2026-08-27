# Inline single statements

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.inline-single-statement |
| Description | Keeps simple single-statement control flow inline. |
| Scope | coding |
| Level | should |
| Tags | coding; style; formatting; control-flow; inline |

## Rule

When a control structure performs one simple statement and the language syntax allows it, the control structure should be written inline. Avoid expanding a single simple statement into a block without a readability or language reason.

## Exceptions

Use a block when the statement is complex, when additional statements are likely to be required immediately, or when the language or repository convention requires a block.

## Rationale

Simple control flow should remain visually simple.

## Sources

IASI

## Examples

Preferred:

```text
if (a > b) return value;
```
