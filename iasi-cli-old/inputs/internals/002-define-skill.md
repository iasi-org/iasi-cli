# Responsabilidad del skill define

## Propósito

El skill `define` transforma conocimiento contenido en `inputs/` en una representación estructurada dentro de `definitions/`.

Su trabajo es semántico: debe interpretar el significado de las entradas y decidir qué definitions son necesarias. No debe limitarse a transformar archivos mecánicamente ni asumir correspondencia 1:1 entre archivo de entrada y archivo de salida.

## Comportamiento esperado

El skill debe:

1. Examinar los inputs relevantes disponibles.
2. Identificar conceptos, requisitos, restricciones, decisiones y preguntas.
3. Compararlos con las definitions ya existentes.
4. Crear nuevas definitions cuando exista conocimiento nuevo que lo justifique.
5. Actualizar definitions existentes cuando un input aporte información adicional o más precisa.
6. Separar una definition cuando mezcle conceptos que deban representarse independientemente.
7. Consolidar conocimiento de varios inputs cuando semánticamente represente la misma definition.
8. Mantener la trazabilidad desde cada definition hasta los inputs que la sustentan.
9. Preservar refinamientos humanos que sigan siendo compatibles con los inputs.
10. Evitar sobrescrituras destructivas y duplicados semánticos.

## Reconciliación

La generación de `definitions/` debe entenderse como reconciliación del estado actual, no como regeneración ciega.

Si una definition ya existe, el skill debe razonar sobre ella y sobre la nueva evidencia antes de proponer o realizar cambios. La presencia de nuevo contenido en `inputs/` no justifica por sí sola reemplazar el contenido existente.

## Separación de responsabilidades

El skill decide qué debería cambiar desde el punto de vista semántico.

La mecánica concreta para descubrir archivos, aplicar operaciones, escribir cambios o integrar el proceso en comandos pertenece al runtime y no forma parte de la responsabilidad semántica del skill.
