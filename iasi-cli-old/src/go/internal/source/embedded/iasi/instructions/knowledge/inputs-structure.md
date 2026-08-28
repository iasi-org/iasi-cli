---
id: knowledge.inputs-structure
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

# Inputs structure

## Purpose

Definir la estructura común utilizada por IASI para organizar la información que condiciona, orienta o alimenta un trabajo de ingeniería.

La estructura distingue la información según su origen y conserva explícitamente su naturaleza durante todo su ciclo de vida.

## Rules

La raíz de inputs MUST utilizar esta estructura:

```text
inputs/
├── externals/
├── internals/
└── obtained/
```

### `externals/`

Contiene información recibida desde fuera del proceso de trabajo actual.

La pertenencia a `externals/` describe el origen de la información, no su formato ni su calidad.

### `internals/`

Contiene información creada deliberadamente dentro del propio proceso para orientar, completar, aclarar o dirigir el trabajo.

La información situada en `internals/` forma parte de los inputs activos mientras no sea archivada explícitamente.

### `obtained/`

Contiene conocimiento obtenido durante el propio proceso de ingeniería mediante análisis, inspección, experimentación, ingeniería inversa, observación, pruebas, investigación técnica u otros mecanismos de descubrimiento.

`obtained/` distingue el conocimiento descubierto del contenido recibido externamente y del contenido creado deliberadamente como orientación interna.

## Archived state

Cada una de las tres categorías MAY contener un subdirectorio:

```text
archived/
```

Por tanto, son válidos:

```text
inputs/
├── externals/
│   └── archived/
├── internals/
│   └── archived/
└── obtained/
```

y:

```text
inputs/
├── externals/
│   └── archived/
├── internals/
│   └── archived/
└── obtained/
    └── archived/
```

`archived/` representa un estado de ciclo de vida y MUST NOT considerarse un cuarto tipo de input.

Un elemento archivado conserva siempre su naturaleza original:

```text
external → archived external
internal → archived internal
obtained → archived obtained
```

El directorio `archived/` SHOULD crearse únicamente cuando sea necesario archivar algún elemento.

No es obligatorio crear directorios `archived/` vacíos durante la inicialización de la estructura.

## Constraints

- IASI MUST NOT utilizar un directorio global `inputs/archived/`.
- Un elemento MUST NOT cambiar de `externals`, `internals` u `obtained` únicamente por ser archivado.
- El archivado MUST preservar la categoría de origen del elemento.
- La estructura MUST NOT utilizar categorías adicionales sin una necesidad metodológica explícita.
- La clasificación MUST basarse en el origen de la información, no en su formato.

## Validation

Una estructura de inputs es válida cuando:

- existen las categorías `externals/`, `internals/` y `obtained/`;
- cada elemento puede identificarse por su origen;
- cualquier `archived/` existente está contenido dentro de una de las tres categorías;
- no existe un `inputs/archived/` global;
- los elementos archivados conservan su categoría original;
- los directorios `archived/` solo existen cuando se han necesitado.
