package manager

import (
	"github.com/live-by-unix/guac/pkg/platform"
)

type BrewManager struct{}

func (m *BrewManager) Name() Manager {
	return Brew
}

func (m *BrewManager) IsSupported() bool {
	return platform.IsDarwin() || platform.IsLinux()
}

func (m *BrewManager) Install(pkg string) error {
	return runCommand("brew", "install", pkg)
}

func (m *BrewManager) Remove(pkg string) error {
	return runCommand("brew", "uninstall", pkg)
}

func (m *BrewManager) Upgrade(pkg string) error {
	return runCommand("brew", "upgrade", pkg)
}

func (m *BrewManager) List() ([]string, error) {
	output, err := runCommandWithOutput("brew", "list")
	if err != nil {
		return nil, err
	}
	return splitLines(output), nil
}

func (m *BrewManager) Search(pkg string) ([]string, error) {
	output, err := runCommandWithOutput("brew", "search", pkg)
	if err != nil {
		return nil, err
	}
	return splitLines(output), nil
}
