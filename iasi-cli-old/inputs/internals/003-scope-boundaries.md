# Límites de alcance de define en esta fase

Las siguientes cuestiones están deliberadamente fuera del alcance de la definición actual del skill `define`:

1. No existe todavía un comando canónico `/define`.
2. No se define todavía un gate, checkpoint o transición obligatoria del workflow que ejecute `define`.
3. No se define la interfaz concreta del runtime para descubrir `inputs/` o `definitions/`.
4. No se define la interfaz concreta del runtime para calcular, presentar o aplicar cambios.
5. No se define un contrato estructurado de operaciones entre skill y runtime, como `create`, `update`, `split` o `consolidate` expresado en JSON o YAML.
6. No se implementa runtime en Go como parte de este trabajo.

Estos puntos no son carencias accidentales. Son límites conscientes para separar primero la semántica del skill de la mecánica de ejecución.

Cuando el comportamiento semántico esté validado con casos reales, estas decisiones podrán abordarse como trabajo posterior y explícito.
