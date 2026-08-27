# Avoid pipe-fed stateful read loops

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.process-substitution-for-read-loops |
| Description | Avoids losing variable changes in subshells created by pipelines. |
| Scope | coding/bash |
| Level | should |
| Tags | coding; bash; style; loops; subshell; process-substitution |

## Rule

When a `while read` loop updates state that must remain available after the loop, feed it with redirection or process substitution rather than piping input into the loop.

## Exceptions

A pipe is acceptable when no state from the loop is needed afterwards.

## Rationale

A pipeline may run the loop in a subshell, causing variable updates to disappear when the pipeline ends.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
last=""
while read -r line; do
  last="${line}"
done < <(generate_lines)

printf '%s\n' "${last}"
```

Avoid when state is needed later:

```bash
generate_lines | while read -r line; do
  last="${line}"
done
```
