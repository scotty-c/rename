package rename

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig loads the replacement pairs from the configuration file.
func LoadConfig() (map[string]string, error) {
	var config map[string]string
	err := viper.UnmarshalKey("replacements", &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return config, nil
}

// LoadConfigKeys loads only the search keys from the configuration file.
func LoadConfigKeys() ([]string, error) {
	replacements, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	var keys []string
	for key := range replacements {
		keys = append(keys, key)
	}

	return keys, nil
}
