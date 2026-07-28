package manager

import (
	"github.com/live-by-unix/guac/pkg/platform"
)

type WingetManager struct{}

func (m *WingetManager) Name() Manager {
	return Winget
}

func (m *WingetManager) IsSupported() bool {
	return platform.IsWindows()
}

func (m *WingetManager) Install(pkg string) error {
	return runCommand("winget", "install", "--id", pkg, "--accept-package-agreements", "--accept-source-agreements")
}

func (m *WingetManager) Remove(pkg string) error {
	return runCommand("winget", "uninstall", "--id", pkg)
}

func (m *WingetManager) Upgrade(pkg string) error {
	return runCommand("winget", "upgrade", "--id", pkg)
}

func (m *WingetManager) List() ([]string, error) {
	output, err := runCommandWithOutput("winget", "list")
	if err != nil {
		return nil, err
	}
	return parseWingetList(output), nil
}

func (m *WingetManager) Search(pkg string) ([]string, error) {
	output, err := runCommandWithOutput("winget", "search", pkg)
	if err != nil {
		return nil, err
	}
	return parseWingetSearch(output), nil
}

func parseWingetList(output string) []string {
	lines := splitLines(output)
	var packages []string
	for _, line := range lines {
		fields := splitFields(line)
		if len(fields) > 0 {
			packages = append(packages, fields[0])
		}
	}
	return packages
}

func parseWingetSearch(output string) []string {
	lines := splitLines(output)
	var results []string
	for _, line := range lines {
		fields := splitFields(line)
		if len(fields) > 0 {
			results = append(results, fields[0])
		}
	}
	return results
}
