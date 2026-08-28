# `/archive` command

## Purpose

`/archive` archives one specific active input document.

The document may exist anywhere below one of the three IASI input branches:

```text
inputs/externals/
inputs/internals/
inputs/obtained/
```

The command archives exactly the document explicitly supplied.

## Invocation

```text
/archive <input-document>
```

Examples:

```text
/archive inputs/externals/request.md
/archive inputs/internals/labs/graphics/scope.md
/archive inputs/obtained/research/notes.md
```

## Archive location

Each input branch owns its archive:

```text
inputs/externals/archived/
inputs/internals/archived/
inputs/obtained/archived/
```

The document always remains in its original logical branch:

```text
inputs/externals/... → inputs/externals/archived/...
inputs/internals/... → inputs/internals/archived/...
inputs/obtained/...  → inputs/obtained/archived/...
```

## Canonical command

The canonical IASI command is:

```text
iasi/commands/archive.md
```

Platform-specific integrations may expose `/archive`, but MUST delegate to this canonical command.

## Archived filename

The original filename receives a local-time timestamp in the exact format:

```text
YYYYMMDDhhmmss
```

The timestamp is inserted immediately before the final extension.

Example:

```text
inputs/internals/labs/graphics/create-mcp-graphics.md
```

becomes:

```text
inputs/internals/archived/create-mcp-graphics-20260817002245.md
```

Other examples:

```text
request.md
→ request-20260817002245.md

example.schema.json
→ example.schema-20260817002245.json

requirement
→ requirement-20260817002245
```

Directories below the branch are not reproduced inside `archived/`.

## Operation

On success, `/archive` MUST:

1. resolve the supplied path;
2. verify that it is a regular file;
3. verify that it is below `inputs/externals/`, `inputs/internals/` or `inputs/obtained/`;
4. reject files already below an `archived/` directory;
5. determine the source branch;
6. create that branch's `archived/` directory when necessary;
7. generate the timestamped filename;
8. move the document into that `archived/` directory;
9. preserve the document contents unchanged;
10. report source and destination.

## Validation relationship

Archived documents are not active inputs.

`/validate` MUST ignore the three `archived/` subtrees.

Moving an active input into `archived/` therefore changes the active input fingerprint and makes any previous validation that included it stale.

## Console output

```text
IASI archive

Branch      : internals
Source      : inputs/internals/labs/graphics/create-mcp-graphics.md
Destination : inputs/internals/archived/create-mcp-graphics-20260817002245.md
Status      : ARCHIVED
```

## Exit behaviour

```text
0 = archived successfully
1 = invalid archive request
2 = archive execution/configuration error
```
