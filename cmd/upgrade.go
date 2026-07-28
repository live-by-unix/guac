package cmd

import (
	"fmt"

	"github.com/live-by-unix/guac/pkg/manager"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade <package-manager> <package>",
	Short: "Upgrade a package using the specified package manager",
	Long:  `Upgrade a package using APT, Homebrew, or Winget. Platform-specific rules apply.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		managerName := manager.Manager(args[0])
		packageName := args[1]

		// Check platform support
		if err := manager.CheckPlatformSupport(managerName); err != nil {
			Logger.Errorf("Platform check failed: %v", err)
			fmt.Println("Not supported for your platform.")
			return
		}

		// Get package manager
		pm, err := manager.GetManager(managerName)
		if err != nil {
			Logger.Errorf("Failed to get package manager: %v", err)
			fmt.Printf("Error: %v\n", err)
			return
		}

		Logger.Infof("Upgrading %s using %s", packageName, managerName)
		if err := pm.Upgrade(packageName); err != nil {
			Logger.Errorf("Upgrade failed: %v", err)
			fmt.Printf("Failed to upgrade %s: %v\n", packageName, err)
			return
		}

		Logger.Infof("Successfully upgraded %s", packageName)
		fmt.Printf("Successfully upgraded %s\n", packageName)
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
