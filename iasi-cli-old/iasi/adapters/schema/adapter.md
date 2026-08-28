# IASI adapter schema

An adapter projects IASI artifacts into a platform-specific representation.

Adapters define platform mappings and generation constraints. They MUST NOT duplicate IASI instruction content or redefine IASI behavior. Generated files are projections and are not a source of truth.

## Descriptor

An adapter descriptor is YAML and contains:

- `schema_version`
- `id`
- `platform`
- `supports`
- `instructions`

The distribution version is provided by IASI `VERSION`; adapters have no independent release version.
