# Workflow gates

## Principle

Every forward IASI command MUST check its gate before performing work.

Gate enforcement belongs to the shared IASI runtime.

Platform adapters MUST NOT implement independent or weaker sequencing rules.

## Gate matrix

| Command | Required checkpoint | Successful result |
|---|---|---|
| pre-plan `/validate` | `INPUTS` or stale/revalidation state | `INPUTS_VALIDATED` |
| `/plan` | `INPUTS_VALIDATED` | `PLANNED` |
| post-plan `/validate` | `PLANNED` | `PLAN_VALIDATED` |
| `/execute` | `PLAN_VALIDATED` | `EXECUTED` |
| `/verify` | `EXECUTED` | `VERIFIED` |

Recovery exception:

- after failed `/verify`, `/execute` MAY be re-entered for corrective execution against the same still-current `PLAN_VALIDATED` context;
- successful corrective execution restores `EXECUTED`;
- `/verify` must then be retried.

## Gate failure

A command invoked without its required current checkpoint MUST fail before its main work.

Example:

```text
IASI workflow blocked

Command   : /execute
Required  : PLAN_VALIDATED
Current   : PLANNED

Run /validate successfully before /execute.
```

## Failed predecessor

An unresolved failed forward stage blocks later stages.

## Downstream invalidation

- source-input or effective-instruction changes invalidate `INPUTS_VALIDATED` and everything after it;
- successful `/plan` invalidates prior plan validation, execution and verification;
- current-plan changes invalidate `PLAN_VALIDATED`, execution and verification;
- successful `/execute` invalidates prior verification;
- `/archive` invalidates every checkpoint whose fingerprint depended on the archived active document.

## Shared implementation

Commands MUST use one shared gate implementation conceptually equivalent to:

```text
load_workflow()
refresh_currentness()
require_checkpoint(...)
record_success(...)
record_failure(...)
invalidate_downstream(...)
```

The exact API is implementation detail.

Duplicating independent gate logic inside command adapters or command prompts is prohibited.

## `/archive`

`/archive` is not a forward checkpoint.

It never grants permission for a later forward stage.

It may invalidate current workflow state because it changes active inputs.
