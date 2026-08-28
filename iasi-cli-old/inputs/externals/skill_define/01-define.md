# define

Estamos desarrollando IASI y quiero incorporar un nuevo skill llamado `define`.

## Objetivo

Crear:

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

Si las templates ya existen en el repositorio, reutilízalas y no las sobrescribas innecesariamente. Revísalas para entender el modelo que debe seguir el skill.

El fichero principal será exactamente:

```text
define/skill.md
```

## Qué es `define`

`define` transforma los `inputs/` actuales de un proyecto IASI en una representación estructurada y canónica dentro de:

```text
definitions/
```

Conceptualmente:

```text
inputs/      = lo que entra en IASI
definitions/ = lo que IASI ha entendido
```

La transformación es semántica, no mecánica.

NO existe una relación 1:1 entre inputs y definitions.

Un input puede producir varias definitions.

Varios inputs pueden contribuir a una misma definition.

Una nueva entrada puede simplemente modificar o enriquecer una definition existente.

## Inputs

El skill debe trabajar conceptualmente con:

```text
inputs/
├── externals/
├── internals/
└── obtained/
```

Los inputs son evidencia.

No deben ser modificados por `define`.

No hay que copiar o reformatear inputs dentro de `definitions/`.

## Definitions existentes

Antes de generar cambios, `define` debe considerar las definitions que ya existan.

`definitions/` es estado canónico editable.

Un humano puede modificar una definition para:

* corregirla;
* matizarla;
* ampliarla;
* mejorar su representación.

Por tanto, una nueva ejecución de `define` NO puede asumir que puede destruir y regenerar todo desde cero.

Debe reconciliar:

```text
inputs actuales
+
definitions existentes
+
cambios humanos existentes
```

La regeneración debe ser no destructiva.

## Responsabilidad del skill

El skill es responsable de la interpretación semántica.

Debe:

1. leer los inputs relevantes;
2. leer las definitions existentes;
3. entender conjuntamente la información;
4. identificar unidades semánticas de conocimiento;
5. detectar relaciones entre ellas;
6. detectar requisitos;
7. detectar restricciones;
8. detectar decisiones;
9. detectar cuestiones todavía abiertas;
10. detectar contradicciones;
11. detectar información necesaria que falte;
12. clasificar el conocimiento;
13. reconciliarlo con las definitions existentes;
14. seleccionar la template apropiada;
15. proponer las definitions resultantes.

No debe limitarse a resumir archivos.

## Workflow conceptual

El skill debe documentar explícitamente un workflow similar a:

```text
Discover
   ↓
Read
   ↓
Analyse
   ↓
Classify
   ↓
Reconcile
   ↓
Draft
   ↓
Validate
   ↓
Propose operations
```

### Discover

Obtener los inputs actuales, definitions existentes y templates disponibles mediante los mecanismos proporcionados por el runtime de IASI.

### Read

Leer la información necesaria.

No asumir que cada input puede interpretarse correctamente de forma aislada.

### Analyse

Construir una comprensión semántica conjunta.

Identificar, entre otras cosas:

* conceptos;
* objetivos;
* requisitos;
* restricciones;
* decisiones;
* actores;
* entidades;
* comportamientos;
* relaciones;
* dependencias;
* contradicciones;
* incertidumbres;
* información ausente.

### Classify

Clasificar las unidades SEMÁNTICAS obtenidas durante `Analyse`.

No clasificar archivos.

Tipos iniciales:

```text
definition
requirement
constraint
decision
question
```

El significado aproximado es:

```text
definition
    conocimiento descriptivo:
    qué es algo, cómo funciona, qué significa,
    qué componentes tiene o cómo se relaciona.

requirement
    algo que el sistema/proyecto/proceso debe
    conseguir o satisfacer.

constraint
    un límite, prohibición, frontera o condición
    dentro de la cual debe cumplirse lo anterior.

decision
    una elección deliberada ya realizada.

question
    una cuestión todavía no resuelta.
```

Regla útil:

```text
requirement = qué debe ocurrir
constraint  = dentro de qué límites puede ocurrir
```

La clasificación debe depender del significado y del contexto, no de palabras clave.

Ejemplo:

```text
The system must generate HTML and PDF.
Websites only generate HTML.
We decided to use Quarto.
EPUB support has not yet been resolved.
```

Puede producir:

```text
The system must generate HTML and PDF
    → requirement

Websites only generate HTML
    → constraint

We decided to use Quarto
    → decision

EPUB support has not yet been resolved
    → question
```

`Classify` es una fase interna de `define`.

NO crear:

```text
/classify
```

NO crear:

```text
classifications/
```

NO crear un artefacto intermedio obligatorio para la clasificación.

### Reconcile

Comparar lo entendido con las definitions actuales.

Para cada unidad semántica decidir si corresponde:

```text
create
update
split
consolidate
reclassify
unchanged
```

También debe poder detectar contradicciones o información insuficiente.

No hacer cambios estructurales sin beneficio semántico.

### Draft

Representar el conocimiento usando las templates disponibles en:

```text
define/templates/
```

La selección de template debe realizarse según la naturaleza semántica del conocimiento.

No según:

* el nombre del fichero de origen;
* la carpeta del input;
* palabras clave;
* comodidad de implementación.

### Validate

Antes de proponer cambios comprobar al menos que:

* no se ha inventado información;
* la clasificación es coherente;
* la template elegida corresponde al conocimiento;
* las afirmaciones importantes tienen trazabilidad;
* las contradicciones no se han resuelto arbitrariamente;
* las cuestiones abiertas siguen explícitamente abiertas;
* no se han destruido silenciosamente refinamientos humanos;
* el conjunto resultante de definitions es coherente.

## Templates

Las templates representan cómo se escribe cada clase de conocimiento.

Debe existir esta separación conceptual:

```text
skill
    → entiende y clasifica

templates
    → dan forma a la representación

runtime
    → realiza mecánica determinista y valida estructura
```

`define` debe utilizar inicialmente exclusivamente:

```text
definition.md
requirement.md
constraint.md
decision.md
question.md
```

No añadas nuevos tipos por ahora.

Si una pieza de conocimiento no encaja correctamente en ninguno, debe indicarse el problema en lugar de inventar silenciosamente una nueva categoría.

## Contradicciones

Si dos inputs se contradicen:

1. identificar ambas posiciones;
2. mantener su trazabilidad;
3. comprobar si existen reglas explícitas de precedencia;
4. si no existen, dejar la contradicción sin resolver.

No elegir arbitrariamente una versión.

Una contradicción también es información.

## Información ausente

No inventar información para poder continuar.

Si falta algo necesario:

* indicar qué falta;
* indicar por qué importa;
* determinar si bloquea realmente el siguiente paso;
* utilizar `question` cuando corresponda una respuesta humana.

No pedir información que todavía no sea necesaria.

## Trazabilidad

Las definitions deben conservar referencias a los inputs que soportan materialmente su contenido.

Una definition puede depender de varios inputs.

Un input puede soportar varias definitions.

La trazabilidad no implica una relación de propiedad 1:1.

## Granularidad

No crear automáticamente una definition por input.

No crear una definition por cada frase.

La unidad debe ser semánticamente útil:

* entendible por sí sola;
* mantenible;
* referenciable;
* capaz de evolucionar independientemente cuando tenga sentido.

## Reclassification

El tipo de una definition puede cambiar si nueva evidencia demuestra que su naturaleza era diferente o ha evolucionado.

Por ejemplo:

```text
question:
Which language should implement the IASI runtime?
```

y posteriormente:

```text
We decided to implement the runtime in Go.
```

puede provocar conceptualmente:

```text
question → decision
```

No reclasificar por motivos meramente estilísticos.

## Frontera con el runtime

No implementar ahora código del runtime.

El skill debe asumir que el runtime de IASI se encargará de aspectos deterministas como:

* discovery de archivos;
* fingerprints;
* detección de inputs modificados;
* workflow gates;
* archivado;
* escritura segura;
* estado de ejecución;
* validaciones estructurales deterministas.

El skill se encarga del significado.

No trasladar razonamiento semántico al runtime.

## Trabajo solicitado

1. Inspecciona la estructura actual del repositorio para seguir sus convenciones.
2. Crea `define/skill.md`.
3. Revisa las cinco templates existentes si ya están presentes.
4. Si no están presentes, crea `define/templates/` con las cinco templates indicadas.
5. Mantén el skill declarativo y orientado al comportamiento del agente.
6. No implementes todavía comandos, runtime Go ni integración con `/define`.
7. No introduzcas nuevas categorías.
8. Al terminar, muéstrame:

   * archivos creados o modificados;
   * decisiones de diseño relevantes;
   * cualquier contradicción que hayas encontrado.

No hagas cambios fuera de este alcance sin indicarlo.
