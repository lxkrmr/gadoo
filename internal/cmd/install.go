package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/lxkrmr/godoorpc"
)

const installHelp = `Install an Odoo addon.

Usage:
  gadoo [connection flags] install <addon>

Arguments:
  addon     Technical addon name (e.g. sale, my_custom_addon)

Examples:
  gadoo install my_custom_addon
  gadoo install sale`

// installInput holds the parsed data for an install command.
type installInput struct {
	addon string
}

// parseInstallArgs parses flags and positional args — calculation.
func parseInstallArgs(args []string) (installInput, error) {
	fs := flag.NewFlagSet("install", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.Usage = func() { fmt.Println(installHelp) }

	if err := fs.Parse(args); err != nil {
		return installInput{}, err
	}

	positional := fs.Args()
	if len(positional) == 0 {
		return installInput{}, fmt.Errorf("addon name is required — run 'gadoo install --help'")
	}
	if len(positional) > 1 {
		return installInput{}, fmt.Errorf(
			"unexpected argument %q — install takes exactly one addon name",
			positional[1],
		)
	}

	return installInput{addon: positional[0]}, nil
}

// buildInstallResult shapes the data for the JSON response — pure calculation.
func buildInstallResult(addon string) map[string]any {
	return map[string]any{
		"addon":  addon,
		"result": "installed",
	}
}

// RunInstall executes the install command: installs an Odoo addon.
func RunInstall(args []string, conn ConnFlags) {
	input, err := parseInstallArgs(args)
	if err == flag.ErrHelp {
		os.Exit(0)
	}
	if err != nil {
		write(errorPayload("install", err))
		os.Exit(1)
	}

	client, err := conn.Connect()
	if err != nil {
		write(errorPayload("install", fmt.Errorf("cannot connect to Odoo at %s — is Odoo running?", conn.URL)))
		os.Exit(1)
	}

	moduleID, err := findModuleID(client, input.addon)
	if err != nil {
		write(errorPayload("install", err))
		os.Exit(1)
	}

	_, err = client.ExecuteKW(
		"ir.module.module", "button_immediate_install",
		godoorpc.Args{[]int{moduleID}},
		godoorpc.KWArgs{},
	)
	if err != nil {
		write(errorPayload("install", fmt.Errorf("install failed for addon %q: %w", input.addon, err)))
		os.Exit(1)
	}

	write(successPayload("install", buildInstallResult(input.addon)))
}
