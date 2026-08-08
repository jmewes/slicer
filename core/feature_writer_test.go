package core

import (
	"os"
	"path/filepath"
	"testing"

	messages "github.com/cucumber/messages/go/v28"
)

func TestCalculateFeaturePath(t *testing.T) {
	t.Run("nested describe section", func(t *testing.T) {
		// Given a URI for a nested describe section in a test file
		originalUri := "components/input/add-item-button/add-item-button.component.spec.ts/AddItemButtonComponent/NestedSection"

		// When the feature file path is calculated
		actual := CalculateFeaturePath(originalUri)

		// Then for everything after the test file, the "/" is replaced with "."
		expected := "components/input/add-item-button/AddItemButtonComponent.NestedSection"
		if actual != expected {
			t.Errorf("Expected %s, but got %s", expected, actual)
		}
	})

	t.Run("single describe section", func(t *testing.T) {
		// Given a URI for a single describe section in a test file
		originalUri := "components/input/add-item-button/add-item-button.component.spec.ts/AddItemButtonComponent"

		// When the feature file path is calculated
		actual := CalculateFeaturePath(originalUri)

		// Then the test file name gets discarded
		expected := "components/input/add-item-button/AddItemButtonComponent"
		if actual != expected {
			t.Errorf("Expected %s, but got %s", expected, actual)
		}
	})
}

func TestWriteFeatureFiles_nested_sub_directories(t *testing.T) {
	// Given a path with multiple slashes
	uri := "components/input/add-item-button/AddItemButtonComponent"
	docs := []*messages.GherkinDocument{
		{
			Uri: uri,
			Feature: &messages.Feature{
				Name: "Add item button",
			},
		},
	}

	targetDir := t.TempDir()

	// When the feature files are written into the target directory
	if err := WriteFeatureFiles(docs, targetDir); err != nil {
		t.Fatalf("WriteFeatureFiles returned error: %v", err)
	}

	// Then for each path element (separated by the slashes), a sub-directory gets created
	pathElements := []string{"components", "input", "add-item-button"}
	currentDir := targetDir
	for _, element := range pathElements {
		currentDir = filepath.Join(currentDir, element)
		info, err := os.Stat(currentDir)
		if err != nil {
			t.Fatalf("expected sub-directory %q to exist: %v", currentDir, err)
		}
		if !info.IsDir() {
			t.Fatalf("expected %q to be a directory", currentDir)
		}
	}

	expectedFile := filepath.Join(targetDir, "components", "input", "add-item-button", "AddItemButtonComponent.feature")
	if _, err := os.Stat(expectedFile); err != nil {
		t.Fatalf("expected feature file %q to exist: %v", expectedFile, err)
	}
}
