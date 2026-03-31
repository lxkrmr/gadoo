# gadoo

A CLI tool for managing the lifecycle of Odoo addons.

## What it is

`gadoo` is for developers who work with Odoo addons. The daily
workflow is simple:

```sh
gadoo upgrade foo    # upgrade an addon
gadoo install foo    # install an addon
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

## Usage

Connection flags are required for every command and must come before
the command name:

```sh
gadoo --url <url> --db <db> --user <user> --password <password> <command> [args]
```

If you are a human typing commands, set up a shell alias to avoid
repeating the connection flags:

```sh
alias gadoo='gadoo --url http://localhost:8069 --db mydb --user admin --password secret'
gadoo upgrade foo
gadoo install foo
```

If you are a coding assistant, construct the full command with flags
directly — no alias needed.

All output is JSON.

Run `gadoo <command> --help` for command-specific usage.

## License

MIT
