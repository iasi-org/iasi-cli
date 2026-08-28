# Archive state

## Principle

`/archive` creates no separate iteration state.

There is no `iteration.json`, archive manifest or global `archive/` directory.

The historical state is represented directly by the timestamped document inside the `archived/` directory of its original branch.

## Input tree

The three branches are:

```text
inputs/
├── externals/
│   └── archived/
├── internals/
│   └── archived/
└── obtained/
    └── archived/
```

Documents outside `archived/` are active inputs.

Documents inside `archived/` are historical inputs.

## Example

Before:

```text
inputs/
└── internals/
    └── labs/
        └── graphics/
            └── scope.md
```

After:

```text
inputs/
└── internals/
    ├── archived/
    │   └── scope-20260817002245.md
    └── labs/
        └── graphics/
```

Empty source directories may remain unless another specification defines directory cleanup.

## Validation state

`/archive` does not need a separate archive flag.

The existing validation freshness mechanism remains authoritative.

Before archive:

```text
active inputs hash = A
```

After archive:

```text
active inputs hash = B
```

Because `archived/` is excluded from active inputs:

```text
A != B
```

and the previous validation is stale.

## Acceptance criteria

The implementation is complete when:

1. `iasi/commands/archive.md` exists;
2. `/archive` accepts exactly one document;
3. the document may be anywhere below the three input branches;
4. documents already under `archived/` are rejected;
5. the destination is `<same-branch>/archived/`;
6. nested source directories are flattened;
7. filenames receive a `YYYYMMDDhhmmss` timestamp before the final extension;
8. file contents remain unchanged;
9. existing archive targets are never overwritten;
10. no iteration metadata is created;
11. `/validate` excludes all three `archived/` subtrees from active inputs and fingerprints.
