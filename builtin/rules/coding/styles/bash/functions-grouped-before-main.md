# Functions are grouped before execution

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.functions-grouped-before-main |
| Description | Keeps function definitions together before executable script flow. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; functions; structure |

## Rule

Function definitions must be grouped together before the script's executable flow.

Includes, shell option setup, and constant initialization may appear before function definitions. Executable business logic must not be interleaved between function declarations.

## Exceptions

Code sourced specifically for immediate initialization may follow an external library contract.

## Rationale

Grouping functions makes structure predictable and avoids hidden execution while the reader is still scanning declarations.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred structure:

```bash
#!/usr/bin/env bash

readonly VERSION="1.0"

log_info() {
  ...
}

main() {
  ...
}

main "$@"
```
