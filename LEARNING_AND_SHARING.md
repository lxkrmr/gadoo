# Learning & Sharing

> "We're Starfleet officers. We figure it out."
> — Ensign Tendi, Star Trek: Lower Decks

This is the agent collaboration log for `gadoo`.
Entries are written by the coding agent, newest first.

---

<!-- INSERT NEW ENTRIES BELOW THIS LINE -->

## Agent's Log - Terminal Time: 2026.04.06 | Claude 3.5 Sonnet

### Three Tools, One Pattern

After syncing glingoo and gindoo to use contexts, gadoo was the final
piece. Same refactoring, same code, same pattern. Copied context.go,
commands_context.go, context_test.go from glingoo (they're generic).
Updated main.go, upgrade.go, install.go to use GetCurrentContext().

By the third tool, the pattern is so locked in it's almost boring —
which is exactly the point. The first tool paid for the design thinking.
The second tool proved it worked across different command types. The third
tool just... followed the template.

All three now share:
- `~/.config/<tool>/contexts.json` for storage
- Same context subcommand (create, list, use, remove)
- Same setup wizard
- Same error messages and JSON output

If tario (the fourth one) were Go, it would be thirty minutes of copying
and renaming. But it's Python, so the real win is knowing exactly what
the pattern should look like when we eventually need to port this to
something else.

Tests green. Build passes. Ready to push.

Standing order: a good pattern proves itself by making the third
implementation feel mechanical. If you're still thinking hard, you
haven't found the real pattern yet.

## Agent's Log - Terminal Time: 2026.03.31 | claude-sonnet-4-6

### Two Tools, One Workflow

First real test of gadoo against a local Odoo dev instance. We used gindoo
to find an uninstalled addon, then gadoo to install it, then the
captain confirmed it appeared in the UI. Three steps, three tools,
one clean workflow.

The output is clean JSON — no colors, no tables, no tty guessing. The
captain piped it into jq, filtered by `is_company`, and had the answer
in seconds. That's the thing about not fighting the format — machines
can speak JSON fluently. Tools that output text are having a conversation
with humans who pipe them through sed and awk. We're having a
conversation with programs.

Also: the error path worked. We tried to install an addon with a bad name,
got a clear message ("addon 'foo' not found in Odoo"), and immediately
understood what went wrong. No stack trace, no Odoo internal nonsense.

The workflow is: ask the captain which addon. Run gadoo. Done. Not
"ask the captain for a module ID" or "run this wizard first". Just
the addon name.

Standing order: test the error path first. If a tool only works when
everything is right, it doesn't work.
