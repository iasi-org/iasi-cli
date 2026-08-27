# Console messages include timestamps

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.console-message-timestamps |
| Description | Prefixes operational console messages with a consistent timestamp. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; console; logging; timestamp |

## Rule

Human-readable operational console messages must include a timestamp.

The timestamp must use one consistent, sortable representation throughout the script. Prefer an ISO 8601-like form including date, time, and timezone offset: `YYYY-MM-DDTHH:MM:SS±ZZZZ`.

## Exceptions

Pure data output intended for another command or program must not be prefixed with timestamps.

## Rationale

Timestamps make execution order, duration, and failure diagnosis visible when output is reviewed interactively or from logs.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

IASI

## Examples

Preferred:

```bash
timestamp="$(date +'%Y-%m-%dT%H:%M:%S%z')"
printf '[%s] [INFO] %s\n' "${timestamp}" "Starting build"
```
