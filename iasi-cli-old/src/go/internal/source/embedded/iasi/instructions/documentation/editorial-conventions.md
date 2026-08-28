---
id: documentation.editorial-conventions
version: 0.1.0
status: active
scope: documentation
applies_to:
  - documentation
  - books
  - articles
  - reports
  - technical-prose
---

# Editorial conventions

## Purpose

Complement the baseline documentation style with conventions for developing a continuous argument and handling technical terminology in Spanish documentation.

## Rules

- The agent SHOULD develop related ideas within the same paragraph when they form part of one line of reasoning.
- The agent SHOULD avoid successive isolated one-line sentences when they can be integrated naturally into a paragraph.
- Short isolated sentences MAY be used when they have a clear structural, expressive or didactic purpose.
- The agent SHOULD avoid recurring emphatic or lapidary statements as a substitute for developing an argument.
- The agent SHOULD prefer a developed explanation over a list when the ideas have an argumentative relationship.
- Lists SHOULD be used when their elements are genuinely parallel, independent or sequential.
- Tables SHOULD be used when structured information or comparison is clearer in tabular form.

### English terminology in Spanish documentation

When the document is in Spanish, the agent MAY use an English technical term when it is common in the field or when translation would be artificial, ambiguous or less precise.

- An English term integrated into a Spanish sentence MUST be written in italics.
- Proper names, products, technologies, protocols, commands, identifiers, filenames, APIs and code elements MUST preserve their original spelling and do not require italics merely because they are in English.

Examples:

```markdown
Realizar un *backup* antes de modificar la configuración.

El proceso forma parte del *pipeline* de integración.

GitHub Actions ejecuta el proceso definido en `.github/workflows/`.
```

## Constraints

- The agent MUST NOT use isolated sentences repeatedly merely to create visual emphasis.
- The agent MUST NOT use slogans or sentence fragments as a substitute for explanatory prose.
- When the document is in Spanish, the agent MUST NOT translate a technical English term when the translation reduces precision or is not established in common usage.
- When the document is in Spanish, the agent MUST NOT italicize proper names, technologies, protocols, commands, filenames, APIs or technical identifiers.

## Validation

- Related ideas are developed continuously when they belong to the same argument.
- Isolated sentences have a clear structural, expressive or didactic justification.
- Lists and tables are used because they improve comprehension.
- English technical terms in Spanish prose follow the typography rules above.
- Technical names and code elements retain their canonical spelling.