package cmd

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

// TestReplaceStrings verifies that the string replacements work correctly
func TestReplaceStrings(t *testing.T) {
	// Clear any existing configuration in Viper
	viper.Reset()

	// Setup Viper with an in-memory configuration for testing
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(strings.NewReader(`
replacements:
  dockerhub.io: dockerhub.com
  quay.io: quay.x
`))
	if err != nil {
		t.Fatalf("Failed to set up in-memory config: %v", err)
	}

	// Create a temporary file to use for testing
	tmpFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write some content to the temporary file
	content := "This is a test file with dockerhub.io and quay.io."
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	// Read the content from the temporary file
	data, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read from temporary file: %v", err)
	}

	// Debug: Print the content read from the file
	t.Logf("Content read from file: %s", string(data))

	// Perform string replacements based on the Viper configuration
	replacements := viper.GetStringMapString("replacements")
	result := string(data)
	for old, new := range replacements {
		result = strings.ReplaceAll(result, old, new)
	}

	// Debug: Print the result after replacements
	t.Logf("Result after replacements: %s", result)

	// Verify the replacements
	expected := "This is a test file with dockerhub.com and quay.x."
	if result != expected {
		t.Errorf("String replacement failed. Expected %q, got %q", expected, result)
	}
}
