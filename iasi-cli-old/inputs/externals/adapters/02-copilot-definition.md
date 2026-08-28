# GitHub Copilot adapter

## Purpose

The Copilot adapter currently projects two IASI artifact types:

```text
instructions
commands
```

It does not yet project:

```text
skills
MCP
custom agents
```

## Canonical descriptor

The canonical descriptor lives at:

```text
iasi/adapters/copilot/adapter.yml
```

Use this conceptual structure:

```yaml
schema_version: 1
id: copilot
platform: github-copilot

supports:
  instructions: true
  commands: true
  skills: false
  mcp: false
  agents: false

instructions:
  general:
    type: repository
    target: .github/copilot-instructions.md

  documentation:
    type: path
    target: .github/instructions/documentation.instructions.md
    applyTo: "**/*.md,**/*.qmd"

  code:
    type: path
    target: .github/instructions/code.instructions.md
    applyTo: "**/*.go,**/*.py,**/*.js,**/*.jsx,**/*.ts,**/*.tsx,**/*.java,**/*.kt,**/*.kts,**/*.c,**/*.h,**/*.cpp,**/*.hpp,**/*.cs,**/*.rs,**/*.rb,**/*.php,**/*.R,**/*.r,**/*.lua,**/*.sh,**/*.ps1,**/*.sql"

  diagrams:
    type: path
    target: .github/instructions/diagrams.instructions.md
    applyTo: "**/*.puml,**/*.plantuml,**/*.mmd,**/*.dot"

  knowledge:
    type: path
    target: .github/instructions/inputs.instructions.md
    applyTo: "**/inputs/**/*"

commands:
  type: prompt
  source: commands
  target_dir: .github/prompts
  suffix: .prompt.md
```

The exact descriptor schema may be refined during implementation, but the following behaviour is mandatory:

- instruction targets and `applyTo` values come from the descriptor rather than hard-coded Go paths;
- command projection maps an effective IASI command to a Copilot prompt file;
- `/validate` maps to `.github/prompts/validate.prompt.md`;
- the generated prompt is a projection of `iasi/commands/validate.md`, not an independent command definition.

## Canonical versus generated

```text
iasi/commands/validate.md
          ↓
Copilot adapter
          ↓
.github/prompts/validate.prompt.md
```

The file under `.github/prompts/` is generated and replaceable. The IASI command under `commands/` remains authoritative.
