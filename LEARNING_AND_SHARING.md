# Learning & Sharing

> "We're Starfleet officers. We figure it out."
> — Ensign Tendi, Star Trek: Lower Decks

This is the agent collaboration log for `gadoo`.
Entries are written by the coding agent, newest first.

---

<!-- INSERT NEW ENTRIES BELOW THIS LINE -->

## Agent's Log - Terminal Time: 2026.03.31 | claude-sonnet-4-6

### Two Tools, One Workflow

First real test of gadoo against a local Odoo dev instance. We used gindoo
to find an uninstalled addon, then gadoo to install it, then the
captain confirmed it appeared in the UI. Three steps, three tools,
one clean workflow.

That moment when the captain said "I like gindoo" - that was the
whole point landing in real time. Not the code, not the architecture
discussions, not the ADRs. Just: ask a question, get an answer, do
something with it.

The error case worked well too. Upgrading a nonexistent addon 'foo'
returned exactly the kind of message we designed for - specific, calm,
tells you what to check. No stack trace, no cryptic code. The Rust
inspiration paid off.

One small surprise: installing an already-installed addon just succeeds
silently. Odoo's button_immediate_install is idempotent. We considered
checking state first and returning 'already_installed', decided against
it. The addon is in the desired state. That's enough.

Standing order: tools that complement each other are more valuable
than tools that try to do everything.

## Agent's Log — Terminal Time: 2026.03.31 | claude-sonnet-4-6

### The Tool That Outgrew Its Name Before It Was Built

First shift on gadoo. We had an interesting moment before a single line
of code was written: the captain pointed out that gadoo doesn't actually
need to be limited to custom addons. `ir.module.module` knows about
everything — official modules, community modules, custom ones. The
"custom" framing came from how we'd use it, not from what the tool
could do. So gadoo is just addon lifecycle. Period.

Small insight, real consequence: the README and the ADR say "Odoo addon
lifecycle" now, not "custom addon lifecycle". One word removed, scope
correctly set.

The implementation itself was straightforward. Two commands, one shared
helper. `findModuleID` is the interesting bit — it's just a search on
`ir.module.module` filtered by name, but it needed a good error message
for when the addon doesn't exist. "addon 'foo' not found in Odoo —
check the name and make sure Odoo has scanned it" is the kind of error
that saves five minutes of head-scratching.

The data/calculation/side effects pattern held up well for two simple
commands. Parse args (calculation), connect (side effect), find module
(side effect), call button (side effect), write result (side effect).
Clean lines.

Standing order: the name of a tool shapes how you think about its
scope. Get the name right before the first commit.
