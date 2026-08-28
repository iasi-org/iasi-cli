# Internal agentic runtime interface

## Decision

IASI agentic commands are not public CLI subcommands.

Commands such as:

```text
/validate
/plan
/execute
/verify
/archive
```

are invoked by agentic platforms through their native mechanisms.

However, those commands require deterministic operations implemented by the shared IASI Go runtime, including:

- workflow gate evaluation;
- checkpoint transitions;
- workflow failure persistence;
- fingerprint calculation;
- validation state persistence;
- safe archive operations;
- safe plan replacement;
- workflow invalidation.

Platform prompts and adapters MUST NOT implement those operations independently.

## Internal runtime interface

The `iasi` binary SHALL expose a reserved internal machine-oriented interface under:

```text
iasi __runtime ...
```

This interface is the bridge between agentic commands and the shared Go runtime.

It is not part of the public IASI CLI.

Conceptually:

```text
agentic platform
      ↓
platform adapter / prompt
      ↓
canonical iasi/commands/<command>.md
      ↓
iasi __runtime ...
      ↓
shared Go runtime
```

## Public CLI boundary

The public CLI remains:

```text
iasi install
iasi reinstall
iasi status
iasi version
iasi adapt copilot
```

Internal runtime operations MUST NOT be documented or presented as normal user-facing CLI commands.

The reserved namespace:

```text
__runtime
```

exists for IASI-owned adapters and agentic command execution.

## Responsibilities

The internal runtime interface MUST provide the operations required by canonical IASI commands without duplicating workflow logic in adapters.

It MUST expose enough functionality for the current workflow:

```text
/validate
/plan
/execute
/verify
/archive
```

The exact internal subcommand structure is an implementation detail.

For example, implementations MAY use operations conceptually similar to:

```text
iasi __runtime workflow status
iasi __runtime workflow require <checkpoint>
iasi __runtime workflow transition <checkpoint>
iasi __runtime workflow fail <command>

iasi __runtime validate ...
iasi __runtime plan ...
iasi __runtime archive <document>
```

These examples are illustrative, not a required final syntax.

## Ownership

The dependency direction is:

```text
Copilot / Codex / other platform
              ↓
            adapter
              ↓
      canonical IASI command
              ↓
       iasi __runtime interface
              ↓
      shared internal Go packages
```

Never:

```text
adapter
   ↓
independent gate logic
independent state logic
independent filesystem semantics
```

## Platform independence

All agentic platforms MUST use the same internal runtime semantics.

A Copilot prompt and a future Codex adapter may invoke the runtime differently at the platform layer, but they MUST ultimately reach the same shared IASI implementation.

Changing agentic platform MUST NOT change:

- workflow checkpoints;
- validation semantics;
- archive semantics;
- plan lifecycle;
- workflow state format;
- path-safety rules.

## Failure behaviour

Internal runtime operations MUST return deterministic process exit status.

Recommended convention:

```text
0 = operation succeeded
1 = workflow or semantic operation rejected
2 = runtime/configuration/internal execution error
```

Machine-readable output MAY be added where useful.

Human-readable prompt behaviour MUST NOT depend on parsing unstable free-form console text when structured output is required.

## Security and path safety

The internal namespace is not a security boundary.

Every runtime operation MUST independently enforce the same path, symlink, overwrite and workflow-safety rules defined by the normative IASI inputs.

Adapters are not trusted to pre-validate filesystem operations.

## Implementation rule

Do not solve this by making `/validate`, `/plan`, `/execute`, `/verify` or `/archive` public CLI subcommands.

Do not solve it by embedding Go workflow behaviour inside Copilot prompt files.

Implement a shared internal interface in the `iasi` binary and make the canonical agentic commands/adapters use it.

## Acceptance criteria

This decision is implemented when:

1. the `iasi` binary exposes a reserved `__runtime` namespace;
2. `__runtime` is not presented as part of the public CLI;
3. agentic commands can invoke shared Go runtime behaviour through that interface;
4. workflow gates and transitions are implemented once in shared Go code;
5. filesystem mutation rules are implemented once in shared Go code;
6. Copilot prompts do not persist workflow state themselves;
7. adapters do not duplicate workflow or archive logic;
8. the same runtime interface can be used by future Codex or other adapters;
9. runtime failures produce deterministic exit status;
10. public CLI behaviour remains unchanged.
