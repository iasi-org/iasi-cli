# Use modern command substitution

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.command-substitution |
| Description | Uses the readable $(...) form for command substitution. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; command-substitution |

## Rule

Command substitution must use `$(command)` syntax.

Legacy backtick command substitution must not be used.

## Exceptions

None.

## Rationale

The $(...) form is easier to read, easier to nest, and avoids the escaping complexity of backticks.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
version="$(get_version)"
```

Not preferred:

```bash
version=`get_version`
```
