# Console output uses logging functions

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.console-logging-functions |
| Description | Centralizes operational console output in dedicated logging functions. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; console; logging; functions |

## Rule

Operational console messages must be emitted through focused logging functions such as `log_info`, `log_success`, `log_warn`, `log_error`, and `log_debug`.

Flow and operation functions should call these logging functions instead of reproducing timestamp, type, color, and stream-selection logic inline.

## Exceptions

Data deliberately written to standard output for consumption by another program may use direct `printf` output.

## Rationale

Centralizing console formatting makes message behavior consistent and keeps business flow free of presentation details.

## Sources

IASI

## Examples

Preferred:

```bash
process_project() {
  log_info "Processing ${project}"
  validate_project "${project}"
  log_success "Processed ${project}"
}
```

Not preferred:

```bash
process_project() {
  printf '\033[36m[%s] INFO Processing %s\033[0m\n' "$(date)" "${project}"
  validate_project "${project}"
}
```
