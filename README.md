# gadoo

A CLI tool for managing the lifecycle of Odoo addons.

## What it is

`gadoo` is a dev tool. It is built for local Odoo development instances
only - not for production.

The daily workflow is simple:

```sh
gadoo context create mydev     # one time
gadoo upgrade foo              # upgrade addon
gadoo install bar              # install addon
```

## Install

```sh
go install github.com/lxkrmr/gadoo@latest
```

Requires Go. The binary lands in `~/go/bin/gadoo`, which should already
be in your `$PATH` if you have used `go install` before.

If `@latest` resolves to an older version after a new release, bypass
the module proxy cache with:

```sh
GOPROXY=direct go install github.com/lxkrmr/gadoo@latest
```

## Setup

Before using `gadoo`, create a connection context:

```sh
gadoo context create mydev
```

This will prompt for:
- URL (e.g. http://localhost:8069)
- Database name
- Login user
- Password

The context is saved to `~/.config/gadoo/contexts.json` and can be
reused. If you have multiple Odoo instances:

```sh
gadoo context create mydev
gadoo context create staging
gadoo context list
gadoo context use staging   # switch between contexts
```

## Usage

### Manage contexts

```sh
gadoo context create <name>   # Create a new connection context
gadoo context list            # Show all contexts (current marked with *)
gadoo context use <name>      # Set as current context
gadoo context remove <name>   # Delete a context
```

### Manage addons

```sh
gadoo upgrade <addon>   # Upgrade an installed addon
gadoo install <addon>   # Install a new addon
```

All output is JSON.

Run `gadoo <command> --help` for command-specific usage.

## License

MIT
