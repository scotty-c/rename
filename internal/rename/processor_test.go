package rename

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// TestProcessFile verifies that the string replacements work correctly in a single file
func TestProcessFile(t *testing.T) {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up the temp file after test

	// Write sample content to the file
	originalContent := "dockerhub.io and quay.io"
	if _, err := tmpfile.WriteString(originalContent); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// Define the replacements
	replacements := map[string]string{
		"dockerhub.io": "dockerhub.com",
		"quay.io":      "quay.x",
	}

	// Process the file
	err = ProcessFile(tmpfile.Name(), replacements)
	if err != nil {
		t.Fatalf("Failed to process file: %v", err)
	}

	// Read the updated content from the file
	updatedContent, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}

	// Verify the content was updated correctly
	expectedContent := "dockerhub.com and quay.x"
	if string(updatedContent) != expectedContent {
		t.Errorf("Expected content '%s', but got '%s'", expectedContent, updatedContent)
	}
}

// TestProcessDirectory verifies that the string replacements work correctly across multiple files in a directory
func TestProcessDirectory(t *testing.T) {
	// Create a temporary directory
	tmpdir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir) // Clean up the temp directory after test

	// Create a couple of temporary files in the directory
	files := []string{"file1.txt", "file2.txt"}
	for _, file := range files {
		tmpfile := filepath.Join(tmpdir, file)
		if err := ioutil.WriteFile(tmpfile, []byte("dockerhub.io and quay.io"), 0644); err != nil {
			t.Fatalf("Failed to create test file '%s': %v", tmpfile, err)
		}
	}

	// Define the replacements
	replacements := map[string]string{
		"dockerhub.io": "dockerhub.com",
		"quay.io":      "quay.x",
	}

	// Process the directory
	err = ProcessDirectory(tmpdir, replacements)
	if err != nil {
		t.Fatalf("Failed to process directory: %v", err)
	}

	// Verify each file was updated correctly
	expectedContent := "dockerhub.com and quay.x"
	for _, file := range files {
		filePath := filepath.Join(tmpdir, file)
		updatedContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			t.Fatalf("Failed to read updated file '%s': %v", filePath, err)
		}

		if string(updatedContent) != expectedContent {
			t.Errorf("Expected content in '%s' to be '%s', but got '%s'", filePath, expectedContent, updatedContent)
		}
	}
}
