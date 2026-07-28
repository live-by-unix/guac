package manager

import (
	"fmt"
	"os/exec"

	"github.com/live-by-unix/guac/pkg/platform"
)

type Manager string

const (
	APT    Manager = "apt"
	Brew   Manager = "brew"
	Winget Manager = "winget"
)

type PackageManager interface {
	Name() Manager
	Install(pkg string) error
	Remove(pkg string) error
	Upgrade(pkg string) error
	List() ([]string, error)
	Search(pkg string) ([]string, error)
	IsSupported() bool
}

func GetManager(m Manager) (PackageManager, error) {
	switch m {
	case APT:
		return &AptManager{}, nil
	case Brew:
		return &BrewManager{}, nil
	case Winget:
		return &WingetManager{}, nil
	default:
		return nil, fmt.Errorf("unknown package manager: %s", m)
	}
}

func CheckPlatformSupport(m Manager) error {
	switch m {
	case APT:
		if !platform.IsLinux() {
			return fmt.Errorf("APT is not supported on %s", platform.String())
		}
	case Brew:
		if !platform.IsDarwin() && !platform.IsLinux() {
			return fmt.Errorf("Brew is not supported on %s", platform.String())
		}
	case Winget:
		if !platform.IsWindows() {
			return fmt.Errorf("Winget is not supported on %s", platform.String())
		}
	}
	return nil
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func runCommandWithOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
