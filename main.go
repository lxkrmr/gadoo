package main

import (
	"fmt"
	"os"

	"github.com/lxkrmr/gadoo/internal/cmd"
)

const help = `gadoo — Odoo addon lifecycle CLI

Usage:
  gadoo <command> [args]

Commands:
  context   Manage connection contexts
  upgrade   Upgrade an Odoo addon
  install   Install an Odoo addon

Examples:
  gadoo context create mydev
  gadoo context list
  gadoo context use mydev
  gadoo upgrade my_addon
  gadoo install my_addon

Run 'gadoo <command> --help' for command-specific usage.`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(help)
		os.Exit(0)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "context":
		cmd.RunContext(args)
	case "upgrade":
		cmd.RunUpgrade(args)
	case "install":
		cmd.RunInstall(args)
	case "help":
		fmt.Println(help)
	default:
		cmd.WriteError("", fmt.Errorf("unknown command %q — run gadoo --help", command))
		os.Exit(1)
	}
}
