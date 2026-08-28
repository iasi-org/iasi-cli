# iasi instructions

`instructions` contains persistent rules that define how an AI agent must behave and how it must produce work inside the iasi methodology.

Instructions are independent of any specific model, agent runtime or vendor. Codex, Copilot, Claude or any other platform may require an adapter, but the source rule belongs here.

## Principles

An instruction should be:

- **Atomic**: focused on one behavioral concern.
- **Declarative**: describe the rule, not a vendor-specific implementation.
- **Composable**: usable together with other instructions.
- **Observable**: compliance should be reviewable.
- **Stable**: avoid temporary project details unless the instruction is explicitly project-scoped.
- **Minimal**: add a rule only when it changes behavior in a useful way.

## Structure

```text
instructions/
├── README.md
├── schema/
│   └── instructions.md
├── general/
│   ├── behavior.md
│   ├── human-control.md
│   ├── precedence.md
│   ├── tool-use.md
│   ├── uncertainty.md
│   └── validation.md
├── documentation/
│   ├── sources.md
│   ├── structure.md
│   └── style.md
├── code/
│   ├── style.md
│   └── testing.md
└── diagrams/
    └── style.md
```

## Composition

General instructions apply by default.

Domain instructions are added when the task belongs to that domain. For example, a documentation task loads `general/*` plus `documentation/*`.

Project-specific instructions may later refine these rules without changing the shared baseline.

## Lifecycle

Instructions use one of these states:

- `draft`: being designed or tested.
- `active`: part of the current iasi baseline.
- `deprecated`: retained for traceability but should not be used for new work.

Do not delete an instruction merely because it is superseded. Deprecate it when traceability matters.
