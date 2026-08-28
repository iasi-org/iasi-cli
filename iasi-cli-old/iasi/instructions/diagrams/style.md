---
id: diagrams.style
version: 0.1.0
status: active
scope: diagrams
applies_to:
  - diagrams
  - architecture-views
  - process-flows
  - technical-figures
---

# Diagram style

## Purpose

Define a simple, repeatable visual language for technical diagrams produced under iasi.

## Rules

- A diagram SHOULD communicate one primary idea.
- The agent SHOULD prefer the simplest representation that preserves the required meaning.
- Labels MUST be concise and unambiguous.
- Relationships MUST be visually distinguishable from components.
- Direction and flow SHOULD remain consistent within a diagram.
- Similar concepts SHOULD use consistent shapes and notation.
- Decorative elements SHOULD be subordinate to information.
- Complex subjects SHOULD be split into multiple diagrams when a single figure becomes difficult to read.
- A diagram SHOULD remain understandable without relying on visual ornament.

## Constraints

- The agent MUST NOT add elements that do not contribute meaning.
- The agent SHOULD NOT encode too many independent dimensions through shape, color, line style or position.
- The agent MUST NOT use inconsistent notation for the same concept within one diagram set.
- The agent MUST NOT sacrifice readability merely to fit everything into one figure.

## Validation

- The main idea of the diagram can be identified quickly.
- Every element has an informational purpose.
- Labels and relationships are readable and consistent.
- Complexity is appropriate for the intended audience and medium.
