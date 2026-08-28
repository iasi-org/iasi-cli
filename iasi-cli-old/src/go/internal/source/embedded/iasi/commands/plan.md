# /plan

Produce the executable plan for the current IASI iteration from active inputs
only. Use `inputs/externals/`, `inputs/internals/`, and `inputs/obtained/`,
but ignore every `archived/` subtree. Read the effective composed IASI
instructions and remain within the declared iteration scope.

Require shared workflow checkpoint `INPUTS_VALIDATED` before archiving or
generating anything. If the checkpoint is absent, failed, or stale, stop
without changing the existing plan.

Check and transition the shared gate through the runtime, never by editing
workflow state directly:

```text
iasi __runtime workflow require plan INPUTS_VALIDATED
...
iasi __runtime workflow transition plan PLANNED
```

Generate the minimum coherent set of Markdown planning documents directly in
`inputs/obtained/`. Do not create `inputs/obtained/plan/`. Document boundaries
are semantic: do not split material merely to meet a size limit and do not
combine unrelated concerns merely to reduce the number of documents.

Before generating a replacement plan, archive every active item under
`inputs/obtained/`, excluding its `archived/` subtree, as one preserved
snapshot at `inputs/obtained/archived/plan-YYYYMMDDhhmmss/`. Use one local-time
timestamp containing exactly 14 digits for the snapshot. Preserve original
filenames, contents, and relative structure. Never overwrite an existing
snapshot. If replacement fails, do not silently destroy the previous plan.

The plan may derive execution detail but must not invent material product,
architectural, or behavioural decisions unsupported by active inputs. Report
such gaps instead of choosing on the project's behalf. State out-of-scope work
when it prevents accidental expansion.

The generated plan is active derived input. Its creation changes the active
input fingerprint, making prior validation stale. Report `PLAN GENERATED` and
require `/validate` again before any later gated execution phase. Do not create
a separate plan manifest or other plan state. Only after archive and generation
succeed, persist shared workflow checkpoint `PLANNED`; on failure persist
`failed_command = plan` without destroying the previous active plan.

On failure, record the shared failure through:

```text
iasi __runtime workflow fail plan
```