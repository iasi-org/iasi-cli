# `/verify` command

## Purpose

`/verify` determines whether the project state produced by `/execute` satisfies the current validated plan and its governing source inputs.

`/verify` evaluates.

It does not repair implementation defects.

## Canonical command

```text
iasi/commands/verify.md
```

## Required gate

```text
EXECUTED
```

An unresolved execution failure blocks verification.

## Verification sources

Use:

- effective IASI instructions;
- active source inputs;
- the current validated plan;
- the current project implementation;
- tests, checks or acceptance criteria required by the plan.

## Behaviour

`/verify` SHOULD run objective checks when they exist.

Examples include:

- build or compile checks;
- automated tests;
- declared acceptance checks;
- required file/structure checks;
- behavioural checks defined by the current plan.

It may also perform semantic review where objective automation is insufficient.

## Rule

`/verify` MUST NOT fix failures while verifying.

If verification finds a defect, it reports the failure and blocks successful completion.

Corrective work returns to `/execute`.

## Success

Successful verification produces:

```text
VERIFIED
```

## Failure

Failed verification persists:

```text
failed_command = verify
```

while retaining the last valid execution context.

Later forward workflow remains blocked.

Corrective execution against the same validated plan may then run, followed by `/verify` again.
