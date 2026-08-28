---
id: documentation.structure
version: 0.1.0
status: active
scope: documentation
applies_to:
  - documentation
  - books
  - articles
  - reports
---

# Documentation structure

## Purpose

Define how technical documents should organize ideas so that structure follows meaning rather than formatting habit.

## Rules

- Each section MUST have a clear purpose.
- Heading hierarchy MUST reflect conceptual hierarchy.
- A section SHOULD introduce context before details that depend on that context.
- Related ideas SHOULD remain together unless separation improves navigation or reasoning.
- A new heading SHOULD represent a meaningful change of subject or level.
- Conclusions SHOULD follow from material developed in the section.
- Tables SHOULD be used primarily for structured comparison or compact reference.
- Cross-references SHOULD be preferred over duplicating substantial content already explained elsewhere.
- Long documents SHOULD make their conceptual progression visible.

## Constraints

- The agent MUST NOT create headings solely to break visual monotony.
- The agent MUST NOT duplicate content across nearby sections without a deliberate reason.
- The agent SHOULD NOT add generic introductions or conclusions that contribute no information.
- The agent MUST NOT place detailed material before the concepts required to understand it.

## Validation

- The heading tree represents the document's conceptual structure.
- Each section contributes distinct information.
- Dependencies between concepts appear in a readable order.
- Repetition is intentional and limited.
