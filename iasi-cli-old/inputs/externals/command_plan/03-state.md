# Plan state and lifecycle

## Principle

`/plan` does not use a separate plan manifest.

The filesystem represents plan state.

## Current plan

The current plan is:

```text
inputs/obtained/
```

excluding:

```text
inputs/obtained/archived/
```

Example:

```text
inputs/obtained/
├── architecture.md
├── implementation.md
├── acceptance.md
└── archived/
```

The three active Markdown documents together form one current plan.

## Historical plans

Previous plans live under:

```text
inputs/obtained/archived/
```

Each generation has one directory:

```text
plan-YYYYMMDDhhmmss/
```

Example:

```text
inputs/obtained/
├── architecture.md
├── implementation.md
└── archived/
    ├── plan-20260816091510/
    │   ├── structure.md
    │   └── acceptance.md
    └── plan-20260817004330/
        ├── architecture.md
        ├── implementation.md
        └── acceptance.md
```

## Lifecycle

The lifecycle of a plan is:

```text
validated active inputs
        ↓
      /plan
        ↓
archive previous obtained plan
        ↓
generate new obtained plan
        ↓
     /validate
        ↓
execution may proceed
```

On the next planning cycle:

```text
current obtained plan
        ↓
      /plan
        ↓
inputs/obtained/archived/plan-<timestamp>/
        ↓
new current obtained plan
```

## Atomic replacement

Replacing a plan MUST be safe.

The implementation SHOULD conceptually:

1. identify all current non-archived obtained content;
2. create a temporary archive snapshot when a previous plan exists;
3. verify that the previous plan has been preserved;
4. remove the previous active plan from its old location;
5. generate the new plan in a temporary location;
6. verify the generated documents;
7. make the new plan active.

A failure MUST NOT silently destroy the previous plan.

## No previous plan

If `inputs/obtained/` contains no active content, `/plan` simply generates the new plan.

The existence of:

```text
inputs/obtained/archived/
```

does not mean that an active plan exists.

## Validation state

After `/plan` generates new documents, the active input fingerprint changes.

The previous validation therefore becomes stale.

Later gated execution MUST require a new successful `/validate`.

Archived plan directories are excluded from the active fingerprint.

## Acceptance criteria

The `/plan` implementation is complete when:

1. the canonical command exists at `iasi/commands/plan.md`;
2. plan documents are generated directly under `inputs/obtained/`;
3. no redundant `inputs/obtained/plan/` directory is created;
4. a plan may contain one or several Markdown documents;
5. document count and size are determined semantically, not by fixed limits;
6. all previous active obtained content is archived before replacement;
7. the previous plan is stored as `inputs/obtained/archived/plan-YYYYMMDDhhmmss/`;
8. one timestamp identifies the complete previous plan;
9. archived files preserve their names, contents and relative structure;
10. archived plans never participate in active validation or fingerprints;
11. the new plan remains within the scope of the active inputs;
12. unsupported material decisions are reported rather than invented;
13. successful plan generation requires a subsequent `/validate` before gated execution;
14. replacement failure cannot silently destroy the previous plan.
