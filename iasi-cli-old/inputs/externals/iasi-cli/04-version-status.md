# Version and status

## Version source

The repository-level:

```text
VERSION
```

is the single source of truth for the IASI distribution version.

The standalone executable MUST embed that version and the complete canonical `iasi/` distribution.

Runtime commands MUST NOT require the source repository.

## `iasi version`

```bash
iasi version
```

prints the version embedded in the executable:

```text
IASI <version>
```

## Layer versions

Every installed layer records the version of the executable that installed it in its `manifest.yml`.

Composed inheritance may therefore contain layers from different IASI versions.

A multi-layer effective context has **no single installed version**.

Do not invent one by choosing the nearest layer, the oldest layer or the binary version.

Version differences between layers are informational unless a specific schema or artifact validation rule makes them incompatible.

## `iasi status`

`iasi status` MUST show the complete resolved installation chain, not only the nearest `.iasi`.

Example:

```text
IASI

Binary : 0.3.0

Layers (low → high precedence):
  1. C:/workspace/.iasi          0.2.0
  2. C:/workspace/project/.iasi  0.3.0

Effective:
  Instructions : 18
  Commands     : 1
  Skills       : 0
  MCP          : 0
  Adapters     : 1
```

Counts refer to the **effective resolved artifacts after overrides**, not the sum of raw files in every layer.

If only one layer exists, show one layer.

If no installed layer exists, return non-zero and report that IASI is not installed for the current location.

## Local validation state

If `<cwd>/.iasi/validation.json` exists, `status` MAY display its current declared status as supplementary information, but validation-state reporting is not required for the first implementation of this specification.

It MUST NOT treat validation state as an installation layer.
