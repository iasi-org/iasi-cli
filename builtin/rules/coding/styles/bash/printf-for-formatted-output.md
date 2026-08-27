# Use printf for formatted output

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.printf-for-formatted-output |
| Description | Uses printf for predictable formatted console output. |
| Scope | coding/bash |
| Level | should |
| Tags | coding; bash; style; console; printf; output |

## Rule

Use `printf` instead of `echo` when producing formatted, escaped, colored, or structured output.

The format string should be a literal and dynamic values should be passed as arguments.

## Exceptions

Simple human-readable output without escapes or formatting may use `echo` when the project convention permits it.

## Rationale

printf has predictable formatting semantics and avoids ambiguities around echo options, escape handling, and dynamic format strings.

## Sources

ShellCheck: https://www.shellcheck.net/

IASI

## Examples

Preferred:

```bash
printf '[%s] [%s] %s\n' "${timestamp}" "${level}" "${message}"
```

Not preferred:

```bash
printf "${message}\n"
echo -e "\033[32m${message}\033[0m"
```
