# ADR 0002: JSON-only output

## Status

Accepted

## Context

gadoo is used by developers and coding assistants. Both benefit from
structured, machine-readable output. A text mode adds maintenance
overhead without adding real value.

## Decision

`gadoo` always outputs JSON. There is no text mode and no `--output`
flag.

Every response follows this structure:

```json
{
  "ok": true,
  "command": "upgrade",
  "data": { "addon": "foo", "result": "upgraded" }
}
```

On error:

```json
{
  "ok": false,
  "command": "upgrade",
  "error": "addon 'foo' not found in Odoo — check the name and make sure Odoo has loaded it"
}
```

Output is always pretty-printed. Error messages are specific and
actionable — closer to Rust errors than terse Unix tool output.

## Consequences

Positive:
- consistent, machine-readable output
- pipeable and agent-friendly
- one format to maintain

Negative:
- requires `jq` or similar for comfortable human reading of large outputs
