Estamos desarrollando IASI.

El skill `define` ya ha sido creado y dispone de las cinco templates iniciales:

```text
define/
├── skill.md
└── templates/
    ├── definition.md
    ├── requirement.md
    ├── constraint.md
    ├── decision.md
    └── question.md
```

La revisión realizada confirma que las templates corresponden exactamente a los cinco tipos semánticos actuales:

```text
definition
requirement
constraint
decision
question
```

y que `define/skill.md` respeta las reglas principales:

* interpretación semántica de los inputs;
* ausencia de relación 1:1 entre inputs y definitions;
* trazabilidad;
* reconciliación no destructiva de `definitions/`;
* preservación de refinamientos humanos;
* clasificación de unidades semánticas y no de archivos;
* uso de templates según la naturaleza del conocimiento;
* separación entre razonamiento del skill y mecánica determinista del runtime.

## Objetivo de este cambio

Registrar formalmente `define` como **skill normativo de IASI** dentro de los inputs del propio proyecto.

IASI debe utilizar su propia metodología para desarrollarse.

Por tanto, la definición y evolución del skill `define` debe quedar recogida como input normativo del proyecto y no existir únicamente como implementación dentro de `define/skill.md`.

## Naming

Actualmente puede existir o haberse considerado un nombre como:

```text
command_define
```

No utilizar ese nombre para este input.

`define` todavía NO ha sido definido como comando.

Lo que acabamos de diseñar es un **skill**.

Utilizar como nombre canónico:

```text
skill_define
```

De esta forma distinguimos explícitamente:

```text
skill_define
    → contrato semántico y comportamiento del skill

command_define
    → futuro contrato del comando /define
```

Estos dos conceptos no deben mezclarse.

## Trabajo solicitado

Inspecciona primero la estructura actual de:

```text
inputs/externals/
```

y especialmente:

```text
inputs/externals/README.md
```

para respetar las convenciones existentes.

Después:

1. crea o renombra, según corresponda, el grupo normativo del skill como:

```text
skill_define/
```

2. registra `skill_define` en:

```text
inputs/externals/README.md
```

siguiendo exactamente el formato utilizado por los demás grupos normativos;

3. conserva dentro de ese grupo el material de entrada que ha originado el diseño actual de `define`;

4. asegúrate de que el input describe el comportamiento y principios del **skill**, no de un comando `/define`;

5. no dupliques innecesariamente información que ya exista en el input original;

6. no conviertas `define/skill.md` en el input normativo. El input es evidencia/origen; `define/skill.md` es el artefacto derivado que implementa esa comprensión.

## Distinción fundamental

Mantener explícitamente este modelo:

```text
inputs/externals/skill_define/
        ↓
        evidencia normativa

define/skill.md
define/templates/
        ↓
        representación/implementación actual del skill
```

El input no debe depender de que la implementación actual permanezca idéntica.

La implementación podrá evolucionar posteriormente mediante nuevas entradas y nuevas ejecuciones del workflow.

## No hacer todavía

Este cambio NO debe definir ni implementar:

```text
/define
```

Tampoco debe definir:

* command routing;
* workflow gate;
* checkpoint de ejecución;
* fingerprints;
* detección incremental de cambios;
* interfaz runtime para discovery;
* formato estructurado de operaciones;
* aplicación de `create`, `update`, `split`, `consolidate`, etc.;
* implementación Go;
* integración entre agente y runtime.

Todos esos elementos pertenecen a una etapa posterior.

Especialmente, no crear todavía:

```text
command_define/
```

porque todavía no hemos definido el contrato del comando `/define`.

## Principio metodológico

Este cambio forma parte del uso de IASI para desarrollar IASI.

Conceptualmente:

```text
input normativo
      ↓
define
      ↓
definition / artefacto
      ↓
implementación
      ↓
validación
```

El desarrollo del propio skill debe poder servir como caso real de aplicación de la metodología.

## Al terminar

Muéstrame:

1. archivos creados;
2. archivos modificados;
3. cualquier rename realizado;
4. cómo ha quedado registrado `skill_define` en `inputs/externals/README.md`;
5. cualquier inconsistencia que hayas encontrado entre el input normativo y `define/skill.md`.

No hagas cambios fuera de este alcance.
