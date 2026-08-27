# Prefer Bash double-bracket tests

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.double-bracket-tests |
| Description | Uses Bash [[ ... ]] for string and file tests. |
| Scope | coding/bash |
| Level | should |
| Tags | coding; bash; style; conditions; tests |

## Rule

In Bash scripts, use `[[ ... ]]` instead of `[ ... ]` or `test` for string, file, pattern, and regular-expression conditions when Bash-specific syntax is acceptable.

## Exceptions

Use another form when required by POSIX compatibility or by a construct whose natural Bash form is arithmetic `(( ... ))`.

## Rationale

`[[ ... ]]` avoids pathname expansion and word splitting inside the condition and provides clearer Bash-specific matching semantics.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
if [[ -n "${name}" ]]; then
  process_name "${name}"
fi
```

Not preferred:

```bash
if [ -n "$name" ]; then
  process_name "$name"
fi
```
