package cmd

import (
	"os"

	"github.com/live-by-unix/guac/pkg/config"
	"github.com/live-by-unix/guac/pkg/platform"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	version = "dev"
)

var Logger = logrus.New()

var rootCmd = &cobra.Command{
	Use:   "guac",
	Short: "Guac - Cross-platform package manager wrapper",
	Long: `Guac is a cross-platform CLI package manager wrapper that unifies 
APT, Homebrew, and winget under a single, playful interface. It provides 
consistent commands across Linux, macOS, and Windows.`,
	Version: version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		Logger.Errorf("Error executing command: %v", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.guac/config.json)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}

func initConfig() {
	if cfgFile != "" {
		config.SetConfigFile(cfgFile)
	} else {
		config.SetConfigFile(config.DefaultConfigPath())
	}

	if err := config.InitConfig(); err != nil {
		Logger.Warnf("Could not load config: %v", err)
	}

	// Set up logging
	if rootCmd.PersistentFlags().Changed("verbose") {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.InfoLevel)
	}

	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Detect platform
	platform.Detect()
}
