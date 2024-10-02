package rename

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SearchFilesInDirectory searches for files containing search keys.
func SearchFilesInDirectory(directory string, searchKeys []string) ([]string, error) {
	var matchingFiles []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			for _, key := range searchKeys {
				if strings.Contains(string(content), key) {
					matchingFiles = append(matchingFiles, path)
					break
				}
			}
		}
		return nil
	})

	return matchingFiles, err
}

// ProcessDirectory applies replacements recursively in a directory or to a single file.
func ProcessDirectory(directory string, replacements map[string]string) error {
	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			err := processFile(path, replacements)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// processFile applies replacements in a single file.
func processFile(filePath string, replacements map[string]string) error {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	content := string(input)
	for oldStr, newStr := range replacements {
		content = strings.ReplaceAll(content, oldStr, newStr)
	}

	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filePath, err)
	}

	return nil
}
