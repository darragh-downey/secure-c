package tests

import (
	"os"
	"path/filepath"
	"testing"
)

func loadTestCase(t *testing.T, secure bool, filename string) string {
	t.Helper()
	var path string
	
	switch secure {
	case false:
		path = filepath.Join("..", "test_cases", "insecure", filename)
	default:
		path = filepath.Join("..", "test_cases", "secure", filename)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read test case file %s: %v", filename, err)
	}
	return string(content)
}
