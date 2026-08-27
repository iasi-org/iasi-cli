# Function variables are local

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.local-function-variables |
| Description | Keeps function-specific variables local to their function. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; variables; local; functions |

## Rule

Variables used only by a function must be declared with `local`.

When a local variable receives the result of command substitution and the command status matters, declaration and assignment must be separate statements.

## Exceptions

A value intentionally shared through the script or exported to child processes may be global.

## Rationale

Local variables reduce accidental state coupling and namespace pollution. Separating declaration from status-sensitive assignment preserves the command's exit status.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
load_name() {
  local name
  name="$(read_name)" || return
  printf '%s\n' "${name}"
}
```

Not preferred:

```bash
load_name() {
  name="$(read_name)"
}
```
