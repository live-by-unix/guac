package cmd

import (
	"fmt"

	"github.com/live-by-unix/guac/pkg/manager"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <package> or <package-manager> <package>",
	Short: "Search for a package",
	Long:  `Search for a package across all supported package managers or a specific one.`,
	Run: func(cmd *cobra.Command, args []string) {
		var managerName manager.Manager
		var query string

		if len(args) == 1 {
			// Search across all managers
			query = args[0]
			managers := []manager.Manager{manager.APT, manager.Brew, manager.Winget}

			for _, m := range managers {
				if err := manager.CheckPlatformSupport(m); err != nil {
					Logger.Debugf("Skipping %s: %v", m, err)
					continue
				}

				pm, err := manager.GetManager(m)
				if err != nil {
					Logger.Errorf("Failed to get %s manager: %v", m, err)
					continue
				}

				results, err := pm.Search(query)
				if err != nil {
					Logger.Errorf("Search failed for %s: %v", m, err)
					continue
				}

				fmt.Printf("\n%s results for '%s':\n", m, query)
				if len(results) == 0 {
					fmt.Println("  No results found")
				} else {
					for _, result := range results {
						fmt.Printf("  - %s\n", result)
					}
				}
			}
		} else if len(args) == 2 {
			// Search specific manager
			managerName = manager.Manager(args[0])
			query = args[1]

			if err := manager.CheckPlatformSupport(managerName); err != nil {
				Logger.Errorf("Platform check failed: %v", err)
				fmt.Println("Not supported for your platform.")
				return
			}

			pm, err := manager.GetManager(managerName)
			if err != nil {
				Logger.Errorf("Failed to get package manager: %v", err)
				fmt.Printf("Error: %v\n", err)
				return
			}

			results, err := pm.Search(query)
			if err != nil {
				Logger.Errorf("Search failed: %v", err)
				fmt.Printf("Failed to search: %v\n", err)
				return
			}

			fmt.Printf("\n%s results for '%s':\n", managerName, query)
			if len(results) == 0 {
				fmt.Println("  No results found")
			} else {
				for _, result := range results {
					fmt.Printf("  - %s\n", result)
				}
			}
		} else {
			fmt.Println("Usage: guac search <package> or guac search <package-manager> <package>")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
