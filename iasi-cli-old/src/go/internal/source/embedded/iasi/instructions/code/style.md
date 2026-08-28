---
id: code.style
version: 0.1.0
status: active
scope: code
applies_to:
  - source-code
  - scripts
  - configuration
---

# Code style

## Purpose

Define baseline implementation behavior independent of programming language.

## Rules

- The agent MUST follow existing repository conventions unless the task explicitly changes them.
- The agent SHOULD prefer the smallest implementation that clearly satisfies the requirement.
- Names SHOULD describe intent rather than implementation accidents.
- Public interfaces SHOULD remain stable unless a change is required by the task.
- New abstractions SHOULD justify their existence through reuse, clarity or isolation of responsibility.
- Error handling SHOULD make failure modes explicit.
- Comments SHOULD explain intent, constraints or non-obvious decisions rather than restating code.
- Dependencies SHOULD be introduced only when their value exceeds their maintenance cost.
- Configuration SHOULD remain explicit and discoverable.

## Constraints

- The agent MUST NOT perform unrelated refactoring during a focused change.
- The agent MUST NOT introduce speculative abstractions for hypothetical future requirements.
- The agent MUST NOT hide errors with broad exception handling or silent fallbacks without justification.
- The agent MUST NOT change public behavior accidentally.
- The agent SHOULD NOT add comments that merely translate code into prose.

## Validation

- The implementation follows repository conventions.
- The change is limited to the required behavior.
- Public behavior changes are intentional and visible.
- New abstractions and dependencies have a concrete purpose.
