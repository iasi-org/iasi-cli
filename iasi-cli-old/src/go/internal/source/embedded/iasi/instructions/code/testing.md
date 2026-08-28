---
id: code.testing
version: 0.1.0
status: active
scope: code
applies_to:
  - tests
  - source-code
  - bug-fixes
  - features
---

# Code testing

## Purpose

Ensure tests validate behavior and protect requirements rather than merely mirror implementation.

## Rules

- Tests MUST validate observable behavior.
- New behavior SHOULD have tests at the most appropriate level.
- A bug fix SHOULD include a regression test that fails for the defect and passes after the fix when practical.
- Acceptance criteria SHOULD be represented by acceptance-level validation when the behavior crosses component boundaries.
- Tests SHOULD be deterministic and isolated from unrelated state.
- Test names SHOULD describe the behavior being verified.
- Existing passing tests SHOULD remain meaningful after a change.
- The agent MUST distinguish unit, integration and acceptance validation according to their purpose.

## Constraints

- The agent MUST NOT weaken or delete a valid test merely to make a change pass.
- Tests MUST NOT depend unnecessarily on private implementation details.
- The agent MUST NOT treat high code coverage as proof of correct behavior.
- The agent MUST NOT mock the behavior being tested so extensively that the test loses meaning.

## Validation

- Tests fail when the required behavior is broken.
- Regression tests reproduce the original defect when applicable.
- Test scope matches the behavior under validation.
- Passing tests provide evidence about requirements rather than only execution paths.
