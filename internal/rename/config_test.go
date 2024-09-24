package rename

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadConfigWithFakeHome(t *testing.T) {
	//Create a temporary directory to mock the home directory
	fakeHomeDir, err := os.MkdirTemp("", "fakeHomeDir")
	if err != nil {
		t.Fatal(err)
	}
	// Clean up the temporary directory after the test
	defer os.RemoveAll(fakeHomeDir)

	//Set the $HOME environment variable to the fake home directory
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", fakeHomeDir)
	defer os.Setenv("HOME", originalHome) // Restore the original $HOME after the test

	//Create the .rename/conf.yaml file in the fake home directory
	configDir := filepath.Join(fakeHomeDir, ".rename")
	err = os.MkdirAll(configDir, 0755) // Create ~/.rename directory
	if err != nil {
		t.Fatalf("Failed to create config directory: %v", err)
	}

	configFilePath := filepath.Join(configDir, "conf.yaml")
	mockConfig := `
replacements:
  dockerhub.io: dockerhub.com
  quay.io: quay.x
`
	err = os.WriteFile(configFilePath, []byte(mockConfig), 0644) // Write conf.yaml
	if err != nil {
		t.Fatalf("Failed to write mock config file: %v", err)
	}

	//Test the LoadConfig function with the mocked $HOME and config file
	viper.Reset() // Ensure Viper doesn't have any old configuration loaded
	replacements := LoadConfig()

	// Verify the replacements were loaded correctly
	expected := map[string]string{
		"dockerhub.io": "dockerhub.com",
		"quay.io":      "quay.x",
	}

	for key, val := range expected {
		if replacements[key] != val {
			t.Errorf("Expected replacement for '%s' to be '%s', but got '%s'", key, val, replacements[key])
		}
	}
}
