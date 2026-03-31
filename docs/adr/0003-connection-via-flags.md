# ADR 0003: Connection via flags

## Status

Accepted

## Context

gadoo needs Odoo connection details: URL, database, user, password.
Environment variables for secrets are a security risk. A config file
adds complexity that is not justified for a tool used against a single
local Odoo instance.

## Decision

Connection credentials are passed as CLI flags:
`--url`, `--db`, `--user`, `--password`.

These flags must come before the command name. A shell alias is the
recommended way to avoid repeating them:

    alias gadoo='gadoo --url http://localhost:8069 --db mydb --user admin --password secret'
    gadoo upgrade foo

## Consequences

Positive:
- no config file to manage
- transparent: credentials are visible at the call site
- secure: no env var leakage

Negative:
- flags must be repeated on every call without an alias
