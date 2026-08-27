# Executable module main guard

## Metadata

| Field | Value |
|---|---|
| ID | coding.styles.python.main-guard |
| Description | Separates importable definitions from script execution. |
| Scope | coding/python |
| Level | should |
| Tags | coding; python; style; scripts; main; modules |

## Rule

A Python module that is both importable and directly executable should place execution entry logic behind `if __name__ == "__main__":` and delegate non-trivial work to a `main()` function.

## Exceptions

Tiny throwaway scripts that are intentionally never imported may omit the guard.

## Rationale

The guard prevents side effects during import and makes script logic reusable and testable.

## Sources

Python documentation - __main__: https://docs.python.org/3/library/__main__.html

## Examples

Preferred:

```python
def main():
    return run()

if __name__ == "__main__":
    raise SystemExit(main())
```
