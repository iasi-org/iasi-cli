# Use arrays for lists

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.arrays-for-lists |
| Description | Represents lists and command argument collections as Bash arrays. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; arrays; arguments; quoting |

## Rule

A collection of independent values or command arguments must be represented as a Bash array rather than as a whitespace-delimited string.

Arrays must be expanded with quoted `"${array[@]}"` syntax when passing individual elements.

## Exceptions

A single scalar string that is intentionally parsed as text is not a list and need not use an array.

## Rationale

Arrays preserve element boundaries and avoid unreliable nested quoting, word splitting, and eval-based argument construction.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
args=(--quiet --output "${target}")
command "${args[@]}"
```

Not preferred:

```bash
args="--quiet --output ${target}"
command ${args}
```
