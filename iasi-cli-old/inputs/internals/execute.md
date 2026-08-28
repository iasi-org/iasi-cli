1. runtime compartido de workflow
   - .iasi/workflow.json
   - checkpoints
   - gates
   - persistencia de fallos
   - invalidación descendente
   - fingerprints
   - escrituras atómicas

2. corregir discovery/hash de archived
   - excluir SOLO:
     inputs/externals/archived/
     inputs/internals/archived/
     inputs/obtained/archived/

3. conectar /validate al workflow
   - pre-plan  → INPUTS_VALIDATED
   - post-plan → PLAN_VALIDATED

4. conectar /plan
   - requiere INPUTS_VALIDATED
   - produce PLANNED

5. implementar iasi/commands/execute.md
   - requiere PLAN_VALIDATED
   - produce EXECUTED

6. implementar iasi/commands/verify.md
   - requiere EXECUTED
   - produce VERIFIED

7. conectar /archive
   - invalida checkpoints dependientes del input archivado
   - no produce ningún checkpoint hacia delante

8. actualizar adapter Copilot
   - execute.prompt.md
   - verify.prompt.md

9. actualizar README.md y README_en.md

10. tests end-to-end del workflow completo