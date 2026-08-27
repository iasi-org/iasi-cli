# Console message colors follow message type

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.console-message-colors |
| Description | Uses consistent colors for human-readable console message types. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; console; logging; color |

## Rule

When human-readable console output supports color, each message type must use a consistent semantic color.

IASI Bash scripts use:

- `INFO`: cyan
- `SUCCESS`: green
- `WARN`: yellow
- `ERROR`: red
- `DEBUG`: dim or neutral

Color must supplement the textual message type, never replace it. Color output must be disabled when the destination is not an interactive terminal or when color has been explicitly disabled.

## Exceptions

Machine-readable output must remain free of ANSI color sequences.

## Rationale

Stable semantic colors improve scanning while textual labels preserve meaning in logs, redirected output, accessibility tools, and terminals without color support.

## Sources

IASI

## Examples

Preferred:

```bash
log_info "Reading configuration"
log_success "Configuration loaded"
log_warn "Optional file not found"
log_error "Validation failed"
```

The rendered output includes both the textual type and, when appropriate, its semantic color.
