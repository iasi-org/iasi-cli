# Bash source filenames use snake_case

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.source-filenames-snake-case |
| Description | Uses lowercase snake_case for Bash source file names. |
| Scope | coding/bash |
| Level | should |
| Tags | coding; bash; style; naming; files; snake-case |

## Rule

Bash source filenames should be lowercase and use underscores to separate words.

Executable commands intentionally installed on `PATH` may use a user-facing command name without a `.sh` suffix when that is part of the interface.

## Exceptions

An existing external naming contract may require another filename.

## Rationale

Predictable lowercase names avoid case ambiguity and align source naming with function naming.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```text
build_project.sh
validate_inputs.sh
```

Avoid:

```text
BuildProject.sh
validate-inputs.SH
```
