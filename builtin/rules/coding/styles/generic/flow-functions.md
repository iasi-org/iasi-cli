# Flow functions

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.generic.flow-functions |
| Description | Makes orchestration functions express sequence while focused subfunctions perform the work. |
| Scope | coding |
| Level | must |
| Tags | coding; style; flow; structured-programming; decomposition; readability |

## Rule

A function that performs several conceptually distinct operations must express those operations as a clear sequence of calls to focused subfunctions rather than implementing all details inline.

The same decomposition rule applies recursively: if a delegated function still performs several conceptually distinct operations, it must be decomposed again until each function reaches an appropriate level of detail and has a clear responsibility.

A flow function should read close to pseudocode.

## Exceptions

Do not extract a subfunction when the extraction would only rename a trivial expression or would make the code harder to understand.

## Rationale

Separating flow from operational detail makes the sequence immediately visible, limits local complexity, and lets each level of the program be understood independently.

## Sources

IASI

## Examples

Preferred:

```text
foo(...) {
    cargar_cosas();
    validar_entradas();
    if (A) proceso_1(); else proceso_2();
    aplicar_algo();
    postprocesar();
    return resultado;
}
```

If `validar_entradas()` still contains several conceptual operations, decompose it again:

```text
validar_entradas(...) {
    validar_formato();
    validar_obligatorios();
    validar_consistencia();
    return resultado;
}
```
