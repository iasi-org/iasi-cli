# Use Bash arithmetic syntax

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.bash-arithmetic |
| Description | Uses Bash arithmetic constructs instead of legacy external arithmetic forms. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; arithmetic |

## Rule

Integer arithmetic must use `(( ... ))` or `$(( ... ))`.

Do not use legacy `$[ ... ]`, `let`, or external `expr` for ordinary Bash arithmetic.

## Exceptions

An external command may be appropriate for arithmetic Bash cannot represent safely, such as arbitrary precision or floating-point computation.

## Rationale

Bash arithmetic syntax is clearer, avoids extra processes, and has well-defined shell semantics.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
total=$(( count + 1 ))

if (( total > limit )); then
  return 1
fi
```

Not preferred:

```bash
total="$(expr "${count}" + 1)"
```
