package cmd

import (
	"fmt"

	"github.com/live-by-unix/guac/pkg/manager"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <package-manager> <package>",
	Short: "Remove a package using the specified package manager",
	Long:  `Remove a package using APT, Homebrew, or Winget. Platform-specific rules apply.`,
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

		Logger.Infof("Removing %s using %s", packageName, managerName)
		if err := pm.Remove(packageName); err != nil {
			Logger.Errorf("Remove failed: %v", err)
			fmt.Printf("Failed to remove %s: %v\n", packageName, err)
			return
		}

		Logger.Infof("Successfully removed %s", packageName)
		fmt.Printf("Successfully removed %s\n", packageName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
