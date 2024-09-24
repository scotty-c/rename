package rename

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ProcessFile processes a single file, performing the replacements
func ProcessFile(filePath string, replacements map[string]string) error {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	updatedContent := string(content)
	for old, new := range replacements {
		updatedContent = strings.ReplaceAll(updatedContent, old, new)
	}

	// Write the updated content back to the file
	err = os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

// ProcessDirectory recursively processes files in a directory
func ProcessDirectory(dirPath string, replacements map[string]string) error {
	// Walk through files in the directory
	err := filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process regular files
		if !info.IsDir() {
			log.Printf("Processing file: %s", filePath)
			return ProcessFile(filePath, replacements)
		}
		return nil
	})
	return err
}
