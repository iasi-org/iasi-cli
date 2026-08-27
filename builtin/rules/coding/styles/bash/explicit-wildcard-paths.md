# Wildcards use explicit paths

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.explicit-wildcard-paths |
| Description | Uses explicit path prefixes for filesystem wildcard expansion. |
| Scope | coding/bash |
| Level | should |
| Tags | coding; bash; style; glob; files; safety |

## Rule

Filesystem wildcard expansion should use an explicit path such as `./*` or `"${dir}"/*` rather than a bare `*` when filenames are passed to commands.

## Exceptions

A controlled context where filenames cannot be interpreted as options may justify a simpler form.

## Rationale

A filename beginning with `-` can be interpreted as a command option. Explicit paths reduce that risk.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
rm -- ./*
```

Not preferred:

```bash
rm *
```
