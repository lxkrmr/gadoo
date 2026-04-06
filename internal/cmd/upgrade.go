package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/lxkrmr/godoorpc"
)

const upgradeHelp = `Upgrade an Odoo addon.

Usage:
  gadoo upgrade <addon>

Arguments:
  addon     Technical addon name (e.g. sale, my_custom_addon)

Examples:
  gadoo upgrade my_custom_addon
  gadoo upgrade sale

Uses the current context. Set it with: gadoo context use <name>`

// upgradeInput holds the parsed data for an upgrade command.
type upgradeInput struct {
	addon string
}

// parseUpgradeArgs parses flags and positional args — calculation.
func parseUpgradeArgs(args []string) (upgradeInput, error) {
	fs := flag.NewFlagSet("upgrade", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.Usage = func() { fmt.Println(upgradeHelp) }

	if err := fs.Parse(args); err != nil {
		return upgradeInput{}, err
	}

	positional := fs.Args()
	if len(positional) == 0 {
		return upgradeInput{}, fmt.Errorf("addon name is required — run 'gadoo upgrade --help'")
	}
	if len(positional) > 1 {
		return upgradeInput{}, fmt.Errorf(
			"unexpected argument %q — upgrade takes exactly one addon name",
			positional[1],
		)
	}

	return upgradeInput{addon: positional[0]}, nil
}

// buildUpgradeResult shapes the data for the JSON response — pure calculation.
func buildUpgradeResult(addon string) map[string]any {
	return map[string]any{
		"addon":  addon,
		"result": "upgraded",
	}
}

// RunUpgrade executes the upgrade command: upgrades an Odoo addon.
func RunUpgrade(args []string) {
	input, err := parseUpgradeArgs(args)
	if err == flag.ErrHelp {
		os.Exit(0)
	}
	if err != nil {
		write(errorPayload("upgrade", err))
		os.Exit(1)
	}

	_, ctx, err := GetCurrentContext()
	if err != nil {
		write(errorPayload("upgrade", err))
		os.Exit(1)
	}

	conn := ConvertContextToConnFlags(ctx)

	client, err := conn.Connect()
	if err != nil {
		write(errorPayload("upgrade", fmt.Errorf("cannot connect to Odoo at %s — is Odoo running?", conn.URL)))
		os.Exit(1)
	}

	moduleID, err := findModuleID(client, input.addon)
	if err != nil {
		write(errorPayload("upgrade", err))
		os.Exit(1)
	}

	_, err = client.ExecuteKW(
		"ir.module.module", "button_immediate_upgrade",
		godoorpc.Args{[]int{moduleID}},
		godoorpc.KWArgs{},
	)
	if err != nil {
		write(errorPayload("upgrade", fmt.Errorf("upgrade failed for addon %q: %w", input.addon, err)))
		os.Exit(1)
	}

	write(successPayload("upgrade", buildUpgradeResult(input.addon)))
}
