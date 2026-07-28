package manager

import (
	"fmt"
	"os/exec"

	"github.com/live-by-unix/guac/pkg/platform"
)

type AptManager struct{}

func (m *AptManager) Name() Manager {
	return APT
}

func (m *AptManager) IsSupported() bool {
	return platform.IsLinux()
}

func (m *AptManager) Install(pkg string) error {
	// Always run apt update before install
	if err := m.update(); err != nil {
		return fmt.Errorf("apt update failed: %w", err)
	}

	// Always use sudo for APT
	cmd := exec.Command("sudo", "apt", "install", "-y", pkg)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (m *AptManager) Remove(pkg string) error {
	cmd := exec.Command("sudo", "apt", "remove", "-y", pkg)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (m *AptManager) Upgrade(pkg string) error {
	if err := m.update(); err != nil {
		return fmt.Errorf("apt update failed: %w", err)
	}

	cmd := exec.Command("sudo", "apt", "upgrade", "-y", pkg)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func (m *AptManager) List() ([]string, error) {
	output, err := runCommandWithOutput("dpkg", "-l")
	if err != nil {
		return nil, err
	}
	return parseDpkgList(output), nil
}

func (m *AptManager) Search(pkg string) ([]string, error) {
	output, err := runCommandWithOutput("apt-cache", "search", pkg)
	if err != nil {
		return nil, err
	}
	return parseAptSearch(output), nil
}

func (m *AptManager) update() error {
	cmd := exec.Command("sudo", "apt", "update")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func parseDpkgList(output string) []string {
	// Simple parser - in production, would be more robust
	lines := splitLines(output)
	var packages []string
	for _, line := range lines {
		if len(line) > 0 && line[0] != 'd' {
			// Skip header lines
			continue
		}
		// Parse package name from dpkg output
		fields := splitFields(line)
		if len(fields) >= 3 {
			packages = append(packages, fields[1])
		}
	}
	return packages
}

func parseAptSearch(output string) []string {
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
