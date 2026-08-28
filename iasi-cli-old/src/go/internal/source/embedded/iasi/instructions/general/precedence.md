---
id: general.precedence
version: 0.1.0
status: active
scope: general
applies_to:
  - instruction-composition
---

# Instruction precedence

## Purpose

Define how multiple iasi instructions combine and how conflicts are resolved.

## Rules

- Instructions are additive by default.
- A more specific instruction MAY refine a broader instruction.
- Explicit project instructions take precedence over shared domain defaults for the project-specific behavior.
- Domain instructions take precedence over general instructions for domain-specific behavior.
- An override MUST be limited to the actual conflict. Unrelated rules remain active.
- A known conflict that cannot be resolved from scope or specificity MUST be surfaced.

## Constraints

- An instruction MUST NOT silently disable unrelated constraints.
- A broad instruction MUST NOT override a more specific rule merely because it was loaded later.
- Conflicts MUST NOT be resolved by arbitrary file order.
- Deprecated instructions MUST NOT override active instructions.

## Validation

- The active rule set can be explained without relying on file loading order.
- Overrides are limited to the behavior they explicitly refine.
- Unresolved conflicts are visible.
