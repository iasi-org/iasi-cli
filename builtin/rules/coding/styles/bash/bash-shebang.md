# Bash scripts declare Bash explicitly

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.bash-shebang |
| Description | Declares Bash explicitly when Bash syntax or behavior is required. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; shebang; interpreter |

## Rule

An executable script that relies on Bash features must declare Bash in its shebang and must not advertise itself as POSIX `sh`.

Use the interpreter path convention required by the target environment consistently across the project.

## Exceptions

A script intentionally restricted to POSIX shell must follow POSIX shell rules instead of Bash-specific rules.

## Rationale

The declared interpreter is part of the execution contract. A Bash script run as sh can fail or behave differently.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred examples, depending on project policy:

```bash
#!/bin/bash
```

or:

```bash
#!/usr/bin/env bash
```

Not preferred for Bash-specific code:

```bash
#!/bin/sh
```
