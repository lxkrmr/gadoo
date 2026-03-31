# ADR 0001: upgrade and install only

## Status

Accepted

## Context

The predecessor tools octo and otto grew over time to include many
commands: uninstall, list, translate, add-known-translations,
load-language-terms, doctor, about, profile management.

Most of these commands were built during development ("vibed in") but
are rarely used in practice. The daily workflow for a custom addon
developer is almost always:

    gadoo upgrade foo
    gadoo install foo

Translate and load-language-terms are translation concerns that belong
in a separate, dedicated tool. Uninstall is a rare edge case. List,
doctor, and about added complexity without adding daily value.

## Decision

`gadoo` implements exactly two commands: `upgrade` and `install`.

Nothing else is added without a concrete, present use case.

## Consequences

Positive:
- minimal surface area to maintain
- the tool does one thing: custom addon lifecycle
- fast to learn, fast to use

Negative:
- users who need translate or uninstall must use a different tool
