package cmd

import (
	"fmt"

	"github.com/live-by-unix/guac/pkg/manager"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed packages for each manager",
	Long:  `List all installed packages for APT, Homebrew, and Winget (platform-dependent).`,
	Run: func(cmd *cobra.Command, args []string) {
		managers := []manager.Manager{manager.APT, manager.Brew, manager.Winget}

		for _, m := range managers {
			// Check if manager is supported on this platform
			if err := manager.CheckPlatformSupport(m); err != nil {
				Logger.Debugf("Skipping %s: %v", m, err)
				continue
			}

			pm, err := manager.GetManager(m)
			if err != nil {
				Logger.Errorf("Failed to get %s manager: %v", m, err)
				continue
			}

			packages, err := pm.List()
			if err != nil {
				Logger.Errorf("Failed to list packages for %s: %v", m, err)
				fmt.Printf("Error listing %s packages: %v\n", m, err)
				continue
			}

			fmt.Printf("\n%s packages:\n", m)
			if len(packages) == 0 {
				fmt.Println("  No packages installed")
			} else {
				for _, pkg := range packages {
					fmt.Printf("  - %s\n", pkg)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
