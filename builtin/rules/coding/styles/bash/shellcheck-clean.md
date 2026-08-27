# Bash scripts pass ShellCheck

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.shellcheck-clean |
| Description | Requires Bash scripts to be checked with ShellCheck. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; shellcheck; static-analysis |

## Rule

Bash scripts must be analyzed with ShellCheck and should have no unresolved actionable diagnostics.

A suppressed ShellCheck diagnostic must include a concrete justification close to the suppression.

## Exceptions

A diagnostic may be suppressed when ShellCheck cannot model a deliberate and correct construct, but the reason must be documented.

## Rationale

ShellCheck detects common quoting, expansion, pipeline, test, and control-flow defects that are easy to miss during review.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

ShellCheck: https://www.shellcheck.net/

## Examples

Preferred:

```bash
# shellcheck disable=SC1091  # Loaded from deployment-provided path.
source "${external_config}"
```

Not preferred:

```bash
# shellcheck disable=SC1091
source $external_config
```
