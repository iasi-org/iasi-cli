# `/archive` workflow integration

## Lifecycle operation

`/archive` is not a forward workflow checkpoint.

It does not produce permission to run `/plan`, `/execute` or `/verify`.

## Effect on workflow state

Archiving an active input removes that document from the active input context.

Any workflow checkpoint whose stored input fingerprint included that document becomes stale.

Therefore, after `/archive`, later forward work requires the workflow to establish a valid checkpoint again, normally beginning with `/validate`.

## No bypass

`/archive` MUST NOT clear a failed workflow stage in a way that allows later stages to execute without satisfying their gates.

If archiving changes the context so that the old failure no longer applies, the new context still requires fresh validation before forward progression.
