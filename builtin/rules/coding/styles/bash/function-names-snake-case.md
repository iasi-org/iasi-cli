# Bash function names use snake_case

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.function-names-snake-case |
| Description | Uses lowercase snake_case names for Bash functions. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; naming; functions; snake-case |

## Rule

Bash function names must be lowercase and use underscores to separate words.

A function name must describe the operation it performs. Package or library namespaces may prefix the function name when a project convention requires it.

## Exceptions

An externally imposed API or sourced library contract may require a different name.

## Rationale

Consistent snake_case names make shell functions easy to scan and align with established Bash style.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
load_configuration() {
  ...
}

validate_inputs() {
  ...
}
```

Not preferred:

```bash
loadConfiguration() {
  ...
}

ValidateInputs() {
  ...
}
```
