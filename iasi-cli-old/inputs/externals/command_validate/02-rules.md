# Validation rules

## Severity

Every finding is either:

```text
BLOCKER
WARNING
```

A blocker means continuing would force the next phase to make a material unsafe assumption, violate instructions, choose between incompatible requirements or operate on unusable information.

Any blocker makes validation fail.

A warning identifies a real issue that does not prevent safe progression.

## Required checks

### 1. Compliance with effective IASI instructions

Check project inputs against the effective resolved active instructions.

A project input that requires behaviour forbidden by an applicable instruction is a blocker unless an explicit IASI precedence rule says otherwise.

### 2. Contradictions between inputs

Detect incompatible statements concerning the same subject, including incompatible behaviours, constraints, formats, interfaces or architectural decisions.

Material contradictions are blockers.

### 3. Missing required information

Report information whose absence would force the next phase to invent a material decision.

Do not report every unspecified detail. Missing information is a finding only when it matters to safe execution of the next phase.

### 4. Material ambiguity

Detect statements with multiple plausible interpretations when those interpretations could produce materially different results.

Severity depends on whether the ambiguity blocks safe progression.

### 5. Invalid references

Detect required references that cannot be resolved, such as nonexistent files, undefined components, missing APIs or required documents.

Only references necessary for the next phase are blockers.

### 6. Incompatible constraints

Detect sets of constraints that cannot all be satisfied simultaneously.

### 7. Unresolved decisions

Detect decisions explicitly left open when the next phase requires them.

Example:

```text
Database: PostgreSQL or SQLite, TBD.
```

If the next phase depends on that choice, it is a blocker.

### 8. Duplicate or divergent definitions

Duplication alone is not an error.

Report duplication when copies diverge, create ambiguity or leave authority unclear.

## Validator boundaries

`/validate` MUST NOT:

- modify project inputs;
- modify installed IASI instructions;
- resolve contradictions on behalf of the project;
- fill missing requirements with invented choices;
- select technologies merely to eliminate a blocker;
- design or implement the solution;
- rewrite user material as part of validation.

It MAY explain what information is missing or why two sources conflict.

## Finding structure

Each finding SHOULD contain:

```text
id
severity
type
message
sources
```

Example:

```json
{
  "id": "VAL-002",
  "severity": "blocker",
  "type": "contradiction",
  "message": "Authentication requirements are incompatible.",
  "sources": [
    "inputs/api.md",
    "inputs/security.md"
  ]
}
```

Finding IDs need only be unique within one validation execution.
