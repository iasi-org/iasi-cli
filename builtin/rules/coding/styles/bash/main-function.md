# Non-trivial scripts use main

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.main-function |
| Description | Provides an explicit main flow for non-trivial Bash scripts. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; main; flow-functions |

## Rule

A Bash script containing multiple functions or non-trivial control flow must define a `main()` function.

`main` must be the bottom-most function and the last executable statement of the file must invoke it with `main "$@"`.

## Exceptions

A short linear script with no meaningful decomposition may omit `main`.

## Rationale

An explicit main makes the script entry point obvious, enables local variables, and naturally implements the IASI flow-function pattern.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

IASI

## Examples

Preferred:

```bash
main() {
  load_configuration
  validate_inputs
  execute_work
}

main "$@"
```
