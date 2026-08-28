# Archived inputs

## Exact historical roots

Only these exact subtrees are excluded from active-input discovery and fingerprints:

```text
inputs/externals/archived/
inputs/internals/archived/
inputs/obtained/archived/
```

The exclusion is path-specific.

Do NOT exclude every directory whose basename is `archived`.

For example:

```text
inputs/externals/design/archived/
```

is active content and MUST be discovered and hashed.

## Validation modes

Pre-plan validation ignores the complete active `inputs/obtained/` branch semantically because that branch contains the current/previous plan.

Post-plan validation includes active `inputs/obtained/`.

Both modes always ignore the exact historical archive roots above.
