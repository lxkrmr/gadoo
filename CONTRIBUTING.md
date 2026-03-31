# Contributing

## Commits

Use Conventional Commits.

Format:

```text
type(scope): short description
```

Examples:

```text
feat(upgrade): add upgrade command
fix(install): return clear error when addon not found
docs(adr): add decision for json-only output
refactor(cmd): extract findModuleID helper
test(upgrade): cover addon not found case
```

Rules:
- keep commits small and meaningful
- write commit messages in English
- prefer one focused change per commit
- use a scope that matches the main area you changed

Common types:
- `feat`
- `fix`
- `docs`
- `refactor`
- `test`
- `chore`
