package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lxkrmr/gadoo/internal/cmd"
)

const help = `gadoo — Odoo addon lifecycle CLI

Usage:
  gadoo --url <url> --db <db> --user <user> --password <password> <command> [args]

Commands:
  upgrade   Upgrade an Odoo addon
  install   Install an Odoo addon

Connection flags (required, must come before the command):
  --url       Odoo base URL (e.g. http://localhost:8069)
  --db        Database name
  --user      Login user
  --password  Login password

Examples:
  gadoo --url http://localhost:8069 --db mydb --user admin --password secret upgrade my_custom_addon
  gadoo --url http://localhost:8069 --db mydb --user admin --password secret install my_custom_addon

Tip: use a shell alias to avoid repeating connection flags:
  alias gadoo='gadoo --url http://localhost:8069 --db mydb --user admin --password secret'
  gadoo upgrade my_custom_addon

Run 'gadoo <command> --help' for command-specific usage.`

func main() {
	fs := flag.NewFlagSet("gadoo", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.Usage = func() { fmt.Println(help) }

	var conn cmd.ConnFlags
	cmd.RegisterConnFlags(fs, &conn)

	if err := fs.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	remaining := fs.Args()
	if len(remaining) == 0 {
		fmt.Println(help)
		os.Exit(0)
	}

	switch remaining[0] {
	case "upgrade":
		cmd.RunUpgrade(remaining[1:], conn)
	case "install":
		cmd.RunInstall(remaining[1:], conn)
	case "help":
		fmt.Println(help)
	default:
		cmd.WriteError("", fmt.Errorf("unknown command %q — run gadoo --help", remaining[0]))
		os.Exit(1)
	}
}
