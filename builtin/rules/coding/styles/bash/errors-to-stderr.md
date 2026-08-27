# Errors go to STDERR

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.errors-to-stderr |
| Description | Routes error messages to the standard error stream. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; console; stderr; errors |

## Rule

Error messages must be written to `STDERR`.

Normal data and successful command output must not be mixed with error diagnostics on `STDOUT`.

## Exceptions

None.

## Rationale

Separating data from diagnostics allows pipelines and callers to consume successful output independently from failures.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
log_error() {
  printf '%s\n' "$*" >&2
}
```

Not preferred:

```bash
echo "ERROR: failed"
```
