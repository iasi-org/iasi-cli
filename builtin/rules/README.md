# Built-in rules

This directory contains the rules distributed with IASI.

Rules are Markdown documents with a normalized `Metadata` table followed by `Rule`, `Exceptions`, `Rationale`, `Sources`, and `Examples` sections. `Examples` is always the final section.

Current initial library: 47 rules.

## Structure

```text
rules/
├── coding/
│   └── styles/
│       ├── generic/
│       └── java/
└── docs/
    ├── styles/
    └── books/
```

The normalized rule template is available in `TEMPLATE.md`.
