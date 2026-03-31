# ADR 0004: findModuleID lives in gadoo, not in godoorpc

## Status

Accepted

## Context

Both `upgrade` and `install` need to resolve an addon name to an
`ir.module.module` record ID before calling the button method. This
lookup is shared between the two commands.

It could live in `godoorpc` since it might be needed by other tools.
However, `godoorpc` is a pure RPC transport library — it does not know
about Odoo's business concepts such as `ir.module.module`. Adding
`FindModuleID` would introduce Odoo domain knowledge into a layer that
should only handle JSON-RPC mechanics.

## Decision

`findModuleID` lives in `gadoo`'s internal package.

If a future tool needs the same logic, a shared `odoo-helpers` library
can be introduced at that point. Until then, YAGNI applies.

## Consequences

Positive:
- godoorpc stays a pure RPC transport layer
- no premature abstraction

Negative:
- if a second tool needs findModuleID, the logic must be duplicated
  or a new shared library must be created
