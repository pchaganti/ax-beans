package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

const ConfigFile = "beans.toml"

// ValidStatuses defines the fixed set of allowed status values.
var ValidStatuses = []string{"open", "in-progress", "done"}

// DefaultStatus is the status assigned to new beans when not specified.
const DefaultStatus = "open"

// Config holds the beans configuration.
type Config struct {
	Beans BeansConfig `toml:"beans"`
}

// BeansConfig defines settings for bean creation.
type BeansConfig struct {
	Prefix   string `toml:"prefix"`
	IDLength int    `toml:"id_length"`
}

// Default returns a Config with default values.
func Default() *Config {
	return &Config{
		Beans: BeansConfig{
			Prefix:   "",
			IDLength: 4,
		},
	}
}

// DefaultWithPrefix returns a Config with the given prefix.
func DefaultWithPrefix(prefix string) *Config {
	cfg := Default()
	cfg.Beans.Prefix = prefix
	return cfg
}

// Load reads configuration from the given .beans directory.
// Returns default config if the file doesn't exist.
func Load(root string) (*Config, error) {
	path := filepath.Join(root, ConfigFile)

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Default(), nil
		}
		return nil, err
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// Apply defaults for missing values
	if cfg.Beans.IDLength == 0 {
		cfg.Beans.IDLength = 4
	}

	return &cfg, nil
}

// Save writes the configuration to the given .beans directory.
func (c *Config) Save(root string) error {
	path := filepath.Join(root, ConfigFile)

	data, err := toml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// IsValidStatus returns true if the status is in the allowed list.
func IsValidStatus(status string) bool {
	for _, s := range ValidStatuses {
		if s == status {
			return true
		}
	}
	return false
}

// StatusList returns a comma-separated list of valid statuses.
func StatusList() string {
	return strings.Join(ValidStatuses, ", ")
}
