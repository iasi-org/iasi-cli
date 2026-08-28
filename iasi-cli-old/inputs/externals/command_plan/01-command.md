# `/plan` command

## Purpose

`/plan` produces the executable plan for the current IASI iteration.

It reads the current active inputs and derives the minimum set of planning documents required to make the current iteration executable without exceeding its declared scope.

The generated plan is itself input to IASI.

Therefore all plan documents are written directly under:

```text
inputs/obtained/
```

## Canonical command

The canonical IASI command is:

```text
iasi/commands/plan.md
```

Platform-specific integrations may expose `/plan`, but MUST delegate to this canonical command.

## Input sources

`/plan` consumes the current **pre-plan validated source inputs**:

```text
active inputs/externals/
active inputs/internals/
effective IASI instructions
```

The existing active `inputs/obtained/` content is the previous plan.

It is not semantic source material for the new plan.

Before generating the new plan, `/plan` archives that previous obtained plan according to this contract.

## Output

`/plan` generates one or more Markdown documents directly in:

```text
inputs/obtained/
```

Example:

```text
inputs/obtained/
├── architecture.md
├── implementation.md
├── acceptance.md
└── archived/
```

The command MUST NOT create a redundant `plan/` subdirectory.

The complete set of active documents in `inputs/obtained/`, excluding `archived/`, represents the current plan.

## Number of documents

The number of plan documents is not fixed.

`/plan` may generate:

```text
1 document
```

or:

```text
several documents
```

depending on the semantic structure of the current iteration.

The command MUST generate the minimum coherent set of documents necessary to describe executable work.

It MUST NOT split content merely to satisfy an arbitrary size limit.

It MUST NOT combine unrelated planning concerns merely to reduce file count.

## Document size

There is no fixed word, page or token limit.

Document boundaries are semantic.

A document should contain material that belongs together as one coherent planning unit.

If two parts have clearly different responsibilities or can evolve independently, they SHOULD be separate documents.

If they form one inseparable planning unit, they SHOULD remain together.

## Previous plan

Before generating a new plan, `/plan` MUST archive the complete previous active plan when one exists.

All active content below:

```text
inputs/obtained/
```

excluding:

```text
inputs/obtained/archived/
```

belongs to the previous plan.

It is archived as one unit under:

```text
inputs/obtained/archived/plan-YYYYMMDDhhmmss/
```

Example:

Before:

```text
inputs/obtained/
├── architecture.md
├── implementation.md
├── acceptance.md
└── archived/
```

After archiving the previous plan:

```text
inputs/obtained/
└── archived/
    └── plan-20260817004330/
        ├── architecture.md
        ├── implementation.md
        └── acceptance.md
```

The original filenames and internal directory structure of the previous plan are preserved inside the archived plan directory.

## New plan generation

After the previous plan has been archived, `/plan` generates the new plan directly under:

```text
inputs/obtained/
```

Example:

```text
inputs/obtained/
├── mcp-structure.md
├── bootstrap.md
└── archived/
    └── plan-20260817004330/
        └── ...
```

## Planning objective

The plan MUST describe only the work required by the current iteration.

Example active inputs:

```text
Create an MCP for IASI Graphics.
```

and:

```text
For this iteration, only create the project skeleton.
```

A valid plan may define:

- the project structure;
- the minimum MCP bootstrap;
- the required configuration;
- a minimal startup test;
- explicit out-of-scope work.

It MUST NOT expand the iteration into implementing graphics tools when the inputs explicitly limit the work to the project skeleton.

## Fundamental rule

`/plan` may derive execution detail.

It MUST NOT invent material product, architectural or behavioural decisions that are not justified by the active inputs.

If planning exposes a decision that cannot be made safely from the available information, `/plan` MUST report the gap instead of silently deciding it.

## Validation after planning

The generated plan becomes active input under:

```text
inputs/obtained/
```

Therefore the plan MUST be validated before a later gated execution phase uses it.

Conceptually:

```text
inputs
   ↓
/validate
   ↓
/plan
   ↓
inputs/obtained/*.md
   ↓
/validate
   ↓
execution
```

The second validation verifies that the derived plan remains coherent with the original inputs and applicable instructions.

## Console output

Example:

```text
IASI plan

Previous plan : archived as plan-20260817004330
Generated     : 3 documents

inputs/obtained/
├── architecture.md
├── implementation.md
└── acceptance.md

Status        : PLAN GENERATED
Next          : /validate
```

If no previous plan exists:

```text
IASI plan

Previous plan : none
Generated     : 2 documents

Status        : PLAN GENERATED
Next          : /validate
```
