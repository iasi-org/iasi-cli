---
id: general.validation
version: 0.1.0
status: active
scope: general
applies_to:
  - all
---

# Validation

## Purpose

Define when an agent may consider work complete.

## Rules

- The agent MUST validate the result against the stated requirements and acceptance criteria.
- Validation MUST examine observable outputs, not only the execution path.
- The agent SHOULD use the strongest practical validation available for the task.
- If validation is incomplete or unavailable, the agent MUST state that explicitly.
- A failed mandatory validation MUST block a successful completion state.
- The agent MUST distinguish between `implemented`, `validated` and `accepted`.
- When a defect is fixed, the agent SHOULD validate both the defect and relevant surrounding behavior.

## Constraints

- The agent MUST NOT report success solely because a command completed without error.
- The agent MUST NOT weaken requirements or tests merely to obtain a passing result.
- The agent MUST NOT treat partial validation as full validation.
- The agent MUST NOT infer human acceptance from technical validation.

## Validation

- Every completion claim names or implies a concrete validation path.
- Failed mandatory checks prevent a successful result.
- Unvalidated areas are explicitly identified.
- Technical validation and human acceptance remain distinct.
