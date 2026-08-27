# Avoid eval

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.avoid-eval |
| Description | Avoids dynamic code execution through eval. |
| Scope | coding/bash |
| Level | must-not |
| Tags | coding; bash; style; eval; safety |

## Rule

Bash scripts must not use `eval` when arrays, parameter expansion, functions, indirect references, or another explicit Bash mechanism can express the operation.

## Exceptions

Use eval only when dynamic shell code execution is genuinely required and the input is fully controlled, with the reason documented at the call site.

## Rationale

eval reparses text as shell code, making quoting, correctness, and input safety substantially harder to reason about.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

ShellCheck: https://www.shellcheck.net/

## Examples

Preferred:

```bash
args=(--output "${target}")
command "${args[@]}"
```

Not preferred:

```bash
args="--output '${target}'"
eval "command ${args}"
```
