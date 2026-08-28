# `/verify` acceptance criteria

The implementation is complete when:

1. `iasi/commands/verify.md` exists;
2. the platform adapter projects it without redefining semantics;
3. `/verify` refuses to run without current `EXECUTED`;
4. objective plan-defined checks are run where available;
5. semantic verification may supplement objective checks;
6. `/verify` never repairs defects itself;
7. success transitions workflow state to `VERIFIED`;
8. failure persists and blocks later forward progression;
9. failed verification can return to corrective `/execute`;
10. a new execution invalidates previous verification.
