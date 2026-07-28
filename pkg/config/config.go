package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	PreferredManager string `json:"preferred_manager"`
	UseSudo          bool   `json:"use_sudo"`
}

var cfg Config
var cfgFile string

func DefaultConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".guac", "config.json")
}

func SetConfigFile(path string) {
	cfgFile = path
}

func InitConfig() error {
	if cfgFile == "" {
		cfgFile = DefaultConfigPath()
	}

	// Create config directory if it doesn't exist
	configDir := filepath.Dir(cfgFile)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	// Read config file if it exists
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Create default config
			cfg = Config{
				PreferredManager: getDefaultManager(),
				UseSudo:          true,
			}
			return SaveConfig()
		}
		return err
	}

	return json.Unmarshal(data, &cfg)
}

func getDefaultManager() string {
	switch runtime.GOOS {
	case "linux":
		return "apt"
	case "darwin":
		return "brew"
	case "windows":
		return "winget"
	default:
		return ""
	}
}

func SaveConfig() error {
	if cfgFile == "" {
		cfgFile = DefaultConfigPath()
	}

	configDir := filepath.Dir(cfgFile)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cfgFile, data, 0644)
}

func Get() Config {
	return cfg
}

func SetPreferredManager(manager string) error {
	cfg.PreferredManager = manager
	return SaveConfig()
}

func SetUseSudo(useSudo bool) error {
	cfg.UseSudo = useSudo
	return SaveConfig()
}
