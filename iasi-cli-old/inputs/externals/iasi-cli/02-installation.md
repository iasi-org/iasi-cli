# Installation model

## `iasi install`

`iasi install` always acts on the current working directory.

It exposes no scope flags such as:

```text
--workspace
--project
```

Running:

```bash
cd <directory>
iasi install
```

installs a local IASI layer at:

```text
<directory>/.iasi/
```

Expected installed structure:

```text
.iasi/
├── manifest.yml
├── instructions/
├── commands/
├── skills/
├── mcp/
└── adapters/
```

The installed tree is copied from the distribution embedded in the executable and preserves relative structure and file contents.

No `profile` or installation `type` field is required. There is only one installation operation: install at the active directory.

## Installation identity

A `.iasi` directory is an **installed IASI layer** only when it contains a valid:

```text
.iasi/manifest.yml
```

The existence of a `.iasi/` directory by itself does not make it an installation layer. This distinction allows project-local IASI state, such as validation state, to coexist with inherited installations.

## Manifest

Minimum manifest:

```yaml
schema_version: 1
version: <embedded IASI VERSION>

installed:
  instructions: all
  commands: all
  skills: all
  mcp: all
  adapters: all
```

Do not add `profile: workspace` or equivalent fields.

The manifest version MUST equal the distribution version embedded in the executable that performed the installation.

## Existing `.iasi`

If `.iasi/manifest.yml` already exists, `iasi install` MUST fail because a local installation already exists.

If `.iasi/` exists without `manifest.yml`, installation MAY proceed only when existing content is recognized IASI runtime state and no managed installation artifact would be overwritten ambiguously.

For the current model, this specifically allows an existing local:

```text
.iasi/validation.json
```

to be preserved while installing the local layer.

Unknown or conflicting content MUST cause preflight failure rather than silent overwrite.

## `iasi reinstall`

`iasi reinstall` also acts only on the current working directory.

It MUST replace the managed artifacts of the **local** installed layer from the current executable distribution:

```text
manifest.yml
instructions/
commands/
skills/
mcp/
adapters/
```

It MUST NOT reinstall or modify any parent layer.

It MUST preserve local runtime state such as:

```text
.iasi/validation.json
```

If no local `.iasi/manifest.yml` exists, `iasi reinstall` MUST fail even when a parent installation is inherited.

## Atomicity

Install and reinstall SHOULD use preflight plus staging before replacing managed content. Any ambiguity about ownership or target safety MUST fail rather than partially alter the installation.
