# Modelo de inputs y definitions

## Contexto

IASI necesita separar lo que entra en el sistema de lo que el sistema ha entendido.

`inputs/` contiene las entradas originales. Una entrada puede tener cualquier forma y proceder de cualquier fuente admitida por IASI. Las entradas son evidencia de origen y no se editan para adaptarlas al modelo interno.

`definitions/` contiene la representación estructurada y canónica que IASI deriva de esas entradas.

## Principios

1. La relación entre `inputs/` y `definitions/` no es 1:1.
2. Una entrada puede originar varias definiciones.
3. Varias entradas pueden contribuir a una misma definición.
4. Una definición puede consolidar, separar o relacionar conceptos encontrados en distintas entradas.
5. Toda definición derivada debe conservar trazabilidad suficiente hacia los inputs que la sustentan.
6. `definitions/` es editable por humanos. Un humano puede matizar, precisar o refinar una definición después de su generación.
7. Regenerar o reconciliar `definitions/` a partir de `inputs/` no debe destruir refinamientos humanos válidos.
8. Los inputs originales se conservan como fuente; las definitions expresan el entendimiento actual del sistema.

## Tipos de definition

En esta fase existen exactamente cinco tipos de definition:

- `definition`: concepto o significado que debe quedar establecido.
- `requirement`: comportamiento o capacidad que el sistema debe proporcionar.
- `constraint`: límite o condición que restringe una solución.
- `decision`: elección adoptada entre alternativas o posibilidades.
- `question`: cuestión relevante todavía no resuelta.

No se deben crear tipos adicionales sin una decisión explícita posterior.

## Dogfooding

IASI debe construirse utilizando IASI. El desarrollo del propio sistema debe servir como caso real de uso de su metodología y como prueba continua del workflow.
