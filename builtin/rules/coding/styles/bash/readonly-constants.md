# Constants are readonly uppercase names

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.readonly-constants |
| Description | Uses uppercase readonly names for constants. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; naming; constants; readonly |

## Rule

Script constants and environment variables defined by the script must use uppercase names with underscores between words.

A value intended to remain constant must be marked `readonly` as soon as its final value is known.

## Exceptions

Externally defined environment variable names must retain the names required by their external contract.

## Rationale

Uppercase names make global constants visible and readonly prevents accidental mutation.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
readonly DEFAULT_TIMEOUT=30
readonly OUTPUT_DIR="${project_root}/output"
```

Not preferred:

```bash
defaultTimeout=30
output_dir="${project_root}/output"
```
