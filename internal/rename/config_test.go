package rename

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadConfig(t *testing.T) {
	// Dynamically create the test config file
	tmpDir := t.TempDir() // Use Go's test temporary directory, which is automatically cleaned up
	configFilePath := tmpDir + "/conf.yaml"

	// YAML content as a map of strings to strings (key-value pairs)
	configContent := `
replacements:
  "dockerhub.io": "dockerhub.com"
  "quay.io": "quay.x"
`

	// Write the config file
	err := os.WriteFile(configFilePath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create config file: %v", err)
	}

	// Set the config file path to the dynamically created file
	viper.SetConfigFile(configFilePath)
	err = viper.ReadInConfig()
	if err != nil {
		t.Fatalf("Error reading config: %v", err)
	}

	replacements, err := LoadConfig()
	if err != nil {
		t.Fatalf("Error loading config: %v", err)
	}

	// Verify the replacements are loaded correctly
	if replacements["dockerhub.io"] != "dockerhub.com" {
		t.Errorf("Expected 'dockerhub.com', got '%s'", replacements["dockerhub.io"])
	}

	if replacements["quay.io"] != "quay.x" {
		t.Errorf("Expected 'quay.x', got '%s'", replacements["quay.io"])
	}
}
