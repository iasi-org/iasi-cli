# `/execute` acceptance criteria

The implementation is complete when:

1. `iasi/commands/execute.md` exists;
2. the platform adapter projects it without redefining semantics;
3. `/execute` refuses to run without current `PLAN_VALIDATED`;
4. no project mutation occurs before the gate passes;
5. execution uses the exact current validated plan;
6. scope expansion is prohibited;
7. unsupported material decisions cause failure rather than invention;
8. success transitions workflow state to `EXECUTED`;
9. failure blocks `/verify`;
10. a failed execution can be retried safely;
11. corrective execution after failed verification is supported for the same current validated plan.
