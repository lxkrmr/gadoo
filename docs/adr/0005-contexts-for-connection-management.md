# ADR 0005: Contexts for connection management

## Status

Accepted (supersedes ADR-0003)

## Context

Connection via flags (ADR-0003) is transparent and explicit, but creates
unnecessary noise. For developers working with multiple Odoo instances,
repeating `--url --db --user --password` on every command becomes tedious.

A context system allows connection details to be stored once and reused
without repeating flags on every call, similar to how `docker context` or
`kubectl context` work.

## Decision

Introduce a `context` subcommand to manage named connection profiles:

```sh
gadoo context create mydev          # Interactive wizard
gadoo context list                  # Show all contexts
gadoo context use mydev             # Set as default
gadoo context remove mydev          # Delete context
```

Contexts are stored in `~/.config/gadoo/contexts.json`:

```json
{
  "contexts": {
    "mydev": {
      "url": "http://localhost:8069",
      "db": "mydb",
      "user": "admin",
      "password": "secret"
    }
  },
  "current_context": "mydev"
}
```

Commands use the current context automatically:

```sh
gadoo upgrade my_addon
gadoo install my_addon
```

## Consequences

Positive:
- Reduces noise: no repeated flags
- Supports multiple Odoo instances: can switch contexts easily
- Agent-friendly: credentials stay on user's system
- Explicit: `gadoo context list` shows what's active
- Minimal dependencies: JSON only, no external packages

Negative:
- Password stored in plain text JSON file (acceptable for local dev)
- User must run `context create` before first use
- Changes from "pure flags" to "stateful contexts"

## Notes

Context is stored locally under `~/.config/gadoo/` with restricted
permissions (0600). File is not meant to be version controlled or shared.
For development purposes with default credentials, the risk is acceptable.
