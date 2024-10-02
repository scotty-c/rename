package rename

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSearchFilesInDirectory(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "testDir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	files := map[string]string{
		"file1.txt": "dockerhub.io and quay.io",
		"file2.txt": "no matching strings here",
		"file3.txt": "dockerhub.io",
	}

	for fileName, content := range files {
		tmpFile := filepath.Join(tmpDir, fileName)
		err := ioutil.WriteFile(tmpFile, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file '%s': %v", tmpFile, err)
		}
	}

	searchKeys := []string{"dockerhub.io", "quay.io"}
	matchingFiles, err := SearchFilesInDirectory(tmpDir, searchKeys)
	if err != nil {
		t.Fatalf("Error searching directory: %v", err)
	}

	expectedFiles := []string{
		filepath.Join(tmpDir, "file1.txt"),
		filepath.Join(tmpDir, "file3.txt"),
	}

	if len(matchingFiles) != len(expectedFiles) {
		t.Fatalf("Expected %d matching files, got %d", len(expectedFiles), len(matchingFiles))
	}
}
