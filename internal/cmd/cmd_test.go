package cmd

import (
	"testing"
)

func TestParseUpgradeArgs(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		expectAddon string
		wantErr    bool
	}{
		{
			name:        "valid addon",
			args:        []string{"my_custom_addon"},
			expectAddon: "my_custom_addon",
		},
		{
			name:    "missing addon",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "too many args",
			args:    []string{"foo", "bar"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseUpgradeArgs(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.addon != tt.expectAddon {
				t.Errorf("addon: expected %q, got %q", tt.expectAddon, got.addon)
			}
		})
	}
}

func TestParseInstallArgs(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectAddon string
		wantErr     bool
	}{
		{
			name:        "valid addon",
			args:        []string{"my_custom_addon"},
			expectAddon: "my_custom_addon",
		},
		{
			name:    "missing addon",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "too many args",
			args:    []string{"foo", "bar"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInstallArgs(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.addon != tt.expectAddon {
				t.Errorf("addon: expected %q, got %q", tt.expectAddon, got.addon)
			}
		})
	}
}

func TestBuildUpgradeResult(t *testing.T) {
	result := buildUpgradeResult("foo")
	if result["addon"] != "foo" {
		t.Errorf("addon: expected %q, got %v", "foo", result["addon"])
	}
	if result["result"] != "upgraded" {
		t.Errorf("result: expected %q, got %v", "upgraded", result["result"])
	}
}

func TestBuildInstallResult(t *testing.T) {
	result := buildInstallResult("foo")
	if result["addon"] != "foo" {
		t.Errorf("addon: expected %q, got %v", "foo", result["addon"])
	}
	if result["result"] != "installed" {
		t.Errorf("result: expected %q, got %v", "installed", result["result"])
	}
}
