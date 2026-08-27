# Command failures are checked

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.check-command-status |
| Description | Makes failure handling explicit for commands whose success matters. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; errors; exit-status |

## Rule

The exit status of every command whose failure affects correctness must be handled explicitly or propagated.

A failure must not be accidentally hidden by subsequent commands, assignments, pipelines, or logging.

## Exceptions

Commands whose failure is intentionally irrelevant may ignore the status when that intent is obvious or documented.

## Rationale

Bash composes programs through exit statuses. Explicit handling prevents scripts from reporting success after partial or failed work.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

ShellCheck: https://www.shellcheck.net/

## Examples

Preferred:

```bash
if ! copy_files; then
  log_error "Copy failed"
  return 1
fi
```

or:

```bash
copy_files || return
```

Not preferred:

```bash
copy_files
log_success "Copy complete"
```
