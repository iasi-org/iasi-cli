# /archive

Archive exactly one explicitly supplied active input document.

Accept one path below exactly one of these active branches:

- `inputs/externals/`
- `inputs/internals/`
- `inputs/obtained/`

Resolve the path and verify it is a regular file inside its branch. Reject path
traversal, symbolic-link escapes, files outside those branches, and files
already below any `archived/` directory. Do not infer that sibling or related
files should also be archived.

Move the supplied file without changing its content to the matching branch
archive:

- `inputs/externals/archived/`
- `inputs/internals/archived/`
- `inputs/obtained/archived/`

Flatten directories below the branch. Insert a local-time timestamp immediately
before the final extension using exactly `YYYYMMDDhhmmss`. For example,
`inputs/internals/labs/graphics/scope.md` becomes
`inputs/internals/archived/scope-20260817002245.md`. Never overwrite an
existing archive destination; fail clearly on a collision.

Report the logical branch, source, destination, and `ARCHIVED` status. On
success use exit status `0`; use `1` for an invalid archive request and `2` for
an execution or configuration error.

Do not create archive metadata or iteration state. Archived inputs are
historical and excluded from `/validate` input discovery and active-input
fingerprints. Moving an active input to `archived/` therefore makes any prior
validation state stale. `/archive` is not a forward checkpoint: it must not
grant permission for `/plan`, `/execute`, or `/verify`. Shared workflow state
must invalidate checkpoints whose active-input fingerprint changed, without
clearing a failed stage to bypass its gate.