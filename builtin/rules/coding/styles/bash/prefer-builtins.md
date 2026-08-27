# Prefer Bash builtins and expansions

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.prefer-builtins |
| Description | Uses Bash builtins and parameter expansion for simple shell-native operations. |
| Scope | coding/bash |
| Level | should |
| Tags | coding; bash; style; builtins; performance |

## Rule

Prefer Bash builtins, parameter expansion, pattern matching, and arithmetic constructs over spawning external processes for operations Bash expresses clearly.

Do not replace a clearer external command with obscure shell syntax solely to avoid a process.

## Exceptions

Use an external tool when it expresses the operation more clearly, robustly, or efficiently for non-trivial data processing.

## Rationale

Shell-native operations avoid unnecessary process creation and often preserve quoting and error semantics more directly.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
trimmed="${value#prefix}"
total=$(( x + y ))
```

Avoid for trivial shell-native work:

```bash
trimmed="$(printf '%s' "${value}" | sed 's/^prefix//')"
total="$(expr "${x}" + "${y}")"
```
