# Build and CLI acceptance

## Standalone executable

The distributed executable MUST be self-contained.

It must embed:

```text
VERSION
iasi/
```

and must not require at runtime:

- the source repository;
- an external `iasi/` or `agentics/` directory;
- auxiliary files beside the executable.

`go:embed` or an equivalent deterministic build preparation step may be used. If build staging is necessary because of Go embed path restrictions, the staged copy is build material only and MUST be regenerated from the canonical repository source.

## Required CLI behaviour

The following commands are current scope:

```text
iasi install
iasi reinstall
iasi status
iasi version
iasi adapt copilot
```

Do not remove implemented commands merely because older milestone documents once excluded them.

## Minimum tests

### Install

Verify that `iasi install`:

- acts on the current directory;
- creates a valid local `.iasi/manifest.yml`;
- installs instructions, commands, skills, MCP and adapters from the embedded distribution;
- preserves relative file structure and contents;
- does not write a `profile` field;
- fails when a local installation already exists;
- can coexist with a pre-existing `.iasi/validation.json` without treating it as an installed layer.

### Reinstall

Verify that `iasi reinstall`:

- requires a local installation;
- does not modify parent installations;
- replaces local managed distribution artifacts;
- preserves `.iasi/validation.json`.

### Resolution

Create at least two installation layers in ancestor/child directories and verify that:

- both are discovered;
- parent-only artifacts are inherited;
- child-only artifacts are added;
- child collisions override parent artifacts according to category rules;
- an adapter present only in the parent remains available;
- a child adapter with the same ID replaces the parent adapter atomically.

### Status

Verify that status:

- shows all resolved layers in low-to-high precedence order;
- reports effective counts after overrides;
- does not confuse a state-only `.iasi` directory with an installed layer.

### Version

Verify that `iasi version` uses the embedded distribution version and does not read `VERSION` from the runtime filesystem.

### Independence

Acceptance tests MUST work with a copied standalone executable and temporary directories without access to the source repository.

## Compatibility

The CLI must work at least on Windows and Linux and use platform-safe path operations.
