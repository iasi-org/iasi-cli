---
id: knowledge.inputs-handling
version: 0.1.0
status: active
scope: knowledge
applies_to:
  - inputs
  - requirements
  - specifications
  - analysis
  - reverse-engineering
---

# Inputs handling

## Purpose

Definir cómo deben tratar los agentes la información almacenada en `inputs/`, preservando evidencia, trazabilidad y control humano sobre su ciclo de vida.

## Rules

### General

- El agente MUST preservar la naturaleza y procedencia de cada input.
- El agente MUST distinguir entre información recibida, información creada internamente y conocimiento obtenido durante el trabajo.
- El agente MUST utilizar los inputs activos como evidencia y contexto sin modificarlos para hacerlos encajar con una solución.
- Cuando cambie la interpretación de un input, el agente SHOULD registrar nuevo conocimiento o una nueva aclaración sin alterar la evidencia original.

### `externals/`

- El contenido de `inputs/externals/` MUST tratarse como inmutable.
- El agente MAY leer, analizar, citar, referenciar o derivar conocimiento a partir de un external.
- El agente MUST NOT modificar, corregir, completar, reescribir ni sustituir el contenido de un external.
- Si un external contiene errores, contradicciones o información incompleta, el agente MUST preservar el original y registrar el análisis correspondiente en la ubicación adecuada.

La inmutabilidad se aplica al contenido del elemento. Una orden explícita de archivado MAY cambiar su ubicación desde:

```text
inputs/externals/<element>
```

a:

```text
inputs/externals/archived/<element>
```

sin modificar su contenido.

### `internals/`

- El contenido activo de `inputs/internals/` MAY evolucionar mientras forme parte del trabajo.
- Las modificaciones SHOULD mantener la trazabilidad necesaria para comprender cambios relevantes.
- Cuando un internal deje de formar parte del contexto activo y deba conservarse, MUST archivarse en lugar de eliminarse.

### `obtained/`

- El agente MAY crear y actualizar información en `inputs/obtained/` como resultado del análisis o descubrimiento.
- El conocimiento obtenido SHOULD indicar suficientemente su origen o método de obtención cuando resulte relevante para su validación.
- Un obtained MUST NOT presentarse como external ni como información proporcionada originalmente.
- Cuando un conocimiento obtenido deje de formar parte del contexto activo, MAY archivarse mediante una decisión explícita.

## Archiving

El archivado MUST producirse únicamente como consecuencia de una decisión explícita.

Ejemplos:

```text
archiva este documento
archiva los inputs anteriores
este input ya no está activo, archívalo
```

El agente MUST NOT archivar automáticamente un input por antigüedad, aparente obsolescencia, contradicción con información posterior, falta de uso reciente o conveniencia para simplificar el contexto.

Si el directorio `archived/` correspondiente no existe, el agente MAY crearlo en el momento de ejecutar la orden de archivado.

Archivar significa retirar un elemento del conjunto activo de inputs conservándolo para trazabilidad.

Archivar MUST NOT significar borrar, invalidar históricamente, reescribir, corregir o cambiar la naturaleza original del elemento.

## Constraints

- El agente MUST NOT borrar un input archivado como parte del proceso normal de trabajo.
- El agente MUST NOT alterar el contenido de un external, esté activo o archivado.
- El agente MUST NOT mover información entre `externals/`, `internals/` y `obtained/` para resolver contradicciones o simplificar la estructura.
- El agente MUST NOT archivar elementos sin una instrucción o decisión explícita.
- El agente MUST NOT tratar un elemento archivado como input activo salvo que exista una decisión explícita para recuperarlo.
- El agente MUST NOT modificar evidencia original para reflejar interpretaciones posteriores.

## Validation

El tratamiento de inputs es válido cuando:

- los externals conservan intacto su contenido;
- las interpretaciones posteriores no sustituyen la evidencia original;
- los internals y obtained mantienen su categoría de origen;
- ningún elemento ha sido archivado automáticamente;
- todo elemento archivado permanece dentro de la categoría a la que pertenecía;
- el archivado conserva el contenido y la trazabilidad;
- los elementos archivados no participan en el contexto activo salvo decisión explícita;
- cualquier discrepancia entre inputs permanece visible o queda explicada mediante nuevo conocimiento, sin alterar los originales.
