# Archive rules

## 1. One invocation archives one document

`/archive` MUST operate on exactly one explicitly supplied document.

It MUST NOT select sibling files, referenced files or other apparently completed inputs.

## 2. Valid source branches

A valid source MUST resolve below exactly one of:

```text
inputs/externals/
inputs/internals/
inputs/obtained/
```

## 3. Archived documents cannot be archived again

A source already below:

```text
inputs/externals/archived/
inputs/internals/archived/
inputs/obtained/archived/
```

MUST be rejected.

## 4. Preserve the logical branch

The destination branch is determined by the source branch.

An input from `internals` can only be moved to:

```text
inputs/internals/archived/
```

The same rule applies to `externals` and `obtained`.

## 5. Flatten nested paths

Directories below the logical branch are not preserved.

Example:

```text
inputs/internals/labs/graphics/request.md
```

becomes:

```text
inputs/internals/archived/request-20260817002245.md
```

## 6. Filename format

The target filename is:

```text
<original-stem>-<YYYYMMDDhhmmss><original-final-extension>
```

The timestamp MUST contain exactly 14 decimal digits and no separators.

## 7. Preserve contents exactly

`/archive` is a move operation.

It MUST NOT rewrite, summarize, normalize, reformat or otherwise change the document contents.

## 8. No completion inference

The caller decides which document to archive.

The command MUST NOT infer that an iteration or another input is complete.

## 9. Never overwrite an archive

If the generated destination already exists, `/archive` MUST fail clearly.

It MUST NOT overwrite an existing archived document.

## 10. Path safety

The resolved source path must remain inside one of the three valid active input branches.

Path traversal and symbolic links MUST NOT allow files outside those branches to be moved.

## 11. Validation freshness

The active input fingerprint MUST exclude archived documents.

Therefore moving an active document into its branch `archived/` directory naturally invalidates any validation fingerprint that included that document.
