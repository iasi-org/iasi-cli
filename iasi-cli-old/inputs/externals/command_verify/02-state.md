# Verify state

## Detailed state

IASI MAY maintain:

```text
.iasi/verification.json
```

for verification-specific evidence.

Suggested minimum:

```json
{
  "schema_version": 1,
  "status": "passed",
  "verified_at": "...",
  "instructions_hash": "...",
  "source_inputs_hash": "...",
  "plan_hash": "...",
  "checks": []
}
```

The workflow checkpoint remains authoritative in:

```text
.iasi/workflow.json
```

## Currentness

Verification is current only for the same:

- effective instructions;
- source inputs;
- current plan;
- execution context.

A new `/execute` invalidates previous verification.

## No repair during verification

Verification output may describe required corrections.

It MUST NOT apply them.

Corrections belong to `/execute`.
