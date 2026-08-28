# Execute state

## Detailed state

IASI MAY maintain:

```text
.iasi/execution.json
```

for execution-specific evidence.

When present, its minimum useful fields are:

```json
{
  "schema_version": 1,
  "status": "passed",
  "executed_at": "...",
  "instructions_hash": "...",
  "source_inputs_hash": "...",
  "plan_hash": "...",
  "command_hash": "..."
}
```

Workflow permission remains authoritative in:

```text
.iasi/workflow.json
```

## Currentness

A previous execution cannot authorize verification if:

- effective instructions changed;
- source inputs changed;
- the current plan changed;
- the effective `/execute` command semantics changed in a way the execution state tracks.

## Corrective execution after failed verification

If `/verify` fails while the same validated plan remains current, `/execute` MAY be re-entered to perform corrective work required by that same plan.

It MUST NOT broaden the plan.

Successful corrective execution returns workflow state to:

```text
EXECUTED
```

and `/verify` must run again.
