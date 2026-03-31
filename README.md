# gadoo

A CLI tool for managing the lifecycle of custom Odoo addons.

## What it is

`gadoo` is for developers who build custom Odoo addons. The daily
workflow is simple:

```sh
gadoo upgrade foo    # upgrade a custom addon
gadoo install foo    # install a custom addon
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

Set up a shell alias to avoid repeating them:

```sh
alias gadoo='gadoo --url http://localhost:8069 --db mydb --user admin --password secret'
gadoo upgrade foo
gadoo install foo
```

All output is JSON.

Run `gadoo <command> --help` for command-specific usage.

## License

MIT
