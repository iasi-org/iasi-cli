# Quote variable expansions

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.bash.quote-variable-expansions |
| Description | Quotes variable expansions unless shell semantics deliberately require otherwise. |
| Scope | coding/bash |
| Level | must |
| Tags | coding; bash; style; quoting; variables |

## Rule

Variable expansions must be double-quoted unless unquoted expansion is explicitly required by Bash semantics.

For named variables, prefer the form `"${name}"` over `"$name"`.

## Exceptions

Arithmetic contexts, carefully designed pattern contexts, and other Bash constructs with different expansion rules may omit quotes when that behavior is intentional.

## Rationale

Quoting prevents unintended word splitting and pathname expansion, which are common sources of shell bugs.

## Sources

Google Shell Style Guide: https://google.github.io/styleguide/shellguide.html

## Examples

Preferred:

```bash
file="${base_dir}/${name}"
rm -- "${file}"
printf '%s\n' "${message}"
```

Not preferred:

```bash
rm $file
printf '%s\n' $message
```
