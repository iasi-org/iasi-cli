---
name: define
description: Build and maintain IASI definitions from the current project inputs.
---

# Define

## Purpose

Transform the current set of IASI `inputs/` into a structured, canonical and maintainable representation under `definitions/`.

Definitions represent what IASI has understood from the available inputs.

The transformation is semantic, not mechanical.

There is no required 1:1 relationship between inputs and definitions.

A definition may:

- consolidate information from multiple inputs;
- separate several concepts found in a single input;
- refine an existing definition;
- reveal contradictions between inputs;
- expose missing information;
- remain unchanged when new inputs add no relevant knowledge.

## Source model

Read the current valid inputs from:

- `inputs/externals/`
- `inputs/internals/`
- `inputs/obtained/`

Treat inputs as evidence.

Do not rewrite, normalize or modify them.

Inputs are not definitions.

## Core rules

### Never invent

Do not create facts, requirements, decisions, constraints or assumptions that are not supported by the available inputs.

When relevant information is missing, expose the gap.

When a decision is required but has not been made, represent it as unresolved.

### Understand before writing

Do not process inputs independently.

First understand the complete current input set and the relationships between its contents.

Look for:

- concepts;
- goals;
- requirements;
- constraints;
- decisions;
- rules;
- actors;
- entities;
- relationships;
- dependencies;
- contradictions;
- uncertainties;
- missing information.

Only then determine the appropriate definition structure.

### Organize semantically

Definitions are organized by meaning, not by source file.

Do not reproduce the directory structure or filenames of `inputs/` unless that organization is itself semantically relevant.

Prefer cohesive definitions with a clear responsibility.

Split a definition when it contains independent concepts that should evolve separately.

Merge knowledge when several inputs describe the same concept.

### Maintain a canonical representation

`definitions/` is the editable canonical representation understood by IASI.

Definitions may be edited by humans.

Human edits are therefore part of the current canonical state and must not be silently destroyed during regeneration.

New input processing must reconcile existing definitions with new evidence.

### Preserve intent

Regeneration is non-destructive.

When an existing definition contains information that cannot be derived from the current inputs, do not automatically remove it.

Determine whether it is:

- a valid human refinement;
- knowledge originating from an older input;
- obsolete information;
- contradictory information;
- an unsupported statement.

Preserve it unless there is sufficient evidence and workflow authority to replace or remove it.

### Maintain traceability

Every material statement in a definition must be traceable to its supporting inputs whenever practical.

Each definition must contain a `Sources` section.

Sources identify the input files that support the definition.

Source references establish provenance, not ownership.

A definition may reference many inputs and an input may support many definitions.

## Workflow

### 1. Discover

Obtain the set of current valid inputs, existing definitions and available definition templates from the IASI runtime.

Do not bypass runtime discovery rules.

### 2. Read

Read all inputs relevant to the current definition operation.

Read existing definitions before proposing changes.

Do not assume that an input can be understood correctly in isolation.

### 3. Analyse

Build a semantic model of the available information.

Identify:

- existing concepts;
- new concepts;
- refinements;
- overlaps;
- contradictions;
- unresolved questions;
- dependencies between definitions.

### 4. Classify

Classify semantic units, not source files, as one of:

- `definition`: descriptive knowledge about a concept, component, meaning, or relationship;
- `requirement`: something the system, project, or process must satisfy;
- `constraint`: a boundary, prohibition, or condition within which work must occur;
- `decision`: a deliberate choice already made;
- `question`: information that remains unresolved.

Choose the type from meaning and context, never from a filename, directory, or
keyword alone. If a unit does not fit these types, report the problem rather
than inventing a new category.

### 5. Reconcile

Compare the semantic model with the current contents of `definitions/`.

For each concept determine whether to:

- create a definition;
- update a definition;
- split a definition;
- consolidate definitions;
- reclassify a definition when new evidence changes its semantic nature;
- leave the current definition unchanged;
- flag a contradiction;
- report missing information.

Do not perform structural churn without semantic benefit.

### 6. Draft

Produce the proposed canonical definitions in structured Markdown.

Write for future use by both humans and intelligent systems.

Definitions should be explicit, concise and unambiguous.

Avoid narrative repetition of the original inputs.

### 7. Validate

Before writing, verify that:

- every definition has a clear purpose;
- no unsupported facts have been introduced;
- material claims are traceable;
- contradictory inputs have not been silently resolved;
- unresolved information remains explicitly unresolved;
- human refinements have not been silently discarded;
- the resulting definition set is internally coherent.

If the available information is insufficient for the next workflow step, report the gaps and stop.

"Sufficient" means sufficient to continue, not complete in an absolute sense.

### 8. Write

Return the proposed definition operations to the IASI runtime.

The runtime is responsible for safe filesystem changes.

The skill must not implement its own file mutation mechanism when runtime operations are available.

## Templates

Templates use structured Markdown.

A template should normally contain:

```markdown
# <Definition name>

## Definition

<Canonical description of the concept.>

## Details

<Relevant structured knowledge, when needed.>

## Constraints

<Constraints that apply to the definition, when present.>

## Relationships

<Relationships or dependencies with other definitions, when relevant.>

## Open questions

<Information that remains unresolved, when present.>

## Sources

- <input reference>
````

Sections that carry no information may be omitted except `Sources`.

The exact structure may evolve when a definition type requires a more appropriate representation.

Do not force information into irrelevant sections merely to satisfy a template.

## Contradictions

When inputs contradict each other:

1. do not choose a winner without evidence;
2. identify the conflicting statements;
3. identify their sources;
4. determine whether precedence can be established from explicit project rules;
5. otherwise mark the issue as unresolved.

A contradiction is information.

Do not hide it.

## Missing information

When the inputs are insufficient:

* identify exactly what is missing;
* explain why it matters for the next workflow step;
* avoid requesting information that is not currently necessary;
* do not fabricate defaults to make the workflow continue.

The result of `define` may therefore be a request for additional `inputs/internals/`.

## Responsibility boundary

The `define` skill owns semantic interpretation.

It is responsible for:

* understanding inputs;
* relating information;
* consolidating concepts;
* separating concepts;
* detecting contradictions;
* detecting missing information;
* maintaining semantic coherence;
* proposing canonical definitions.

The IASI runtime owns deterministic mechanics.

It is responsible for:

* discovering files;
* calculating fingerprints;
* determining changed inputs;
* enforcing workflow gates;
* managing archived inputs;
* applying filesystem operations;
* protecting writes;
* recording execution state.

Do not move semantic reasoning into the runtime merely because it can be implemented deterministically.

Do not move deterministic infrastructure into the skill merely because an agent can perform it.

## Principle

`inputs/` contains what enters IASI.

`definitions/` contains what IASI understands.



