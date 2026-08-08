package core

import (
	"os"
	"path/filepath"
	"testing"
	"github.com/experimental-software/gherkin/parsers/javascript"
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

func TestWriteFeatureFiles_NestedSubDirectories(t *testing.T) {
	// Scenario: create nested sub-directories
	
	// Given a spec file "src/app/shared/utils.spec.ts" (relative to the "testdata" source directory)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	err = os.Chdir("testdata")
	if err != nil {
		t.Fatalf("Failed to chdir to testdata: %v", err)
	}
	defer os.Chdir(cwd)
	
	specPath := filepath.Join("src", "app", "shared", "utils.spec.ts")
	
	// When the spec file gets parsed into Gherkin documents
	docs, err := javascript.ParseSpecFile(specPath, true)
	if err != nil {
		t.Fatalf("Failed to parse spec file: %v", err)
	}
	
	// And the feature files created 
	targetDir := t.TempDir()
	err = WriteFeatureFiles(docs, targetDir)
	if err != nil {
		t.Fatalf("Failed to write feature files: %v", err)
	}
	
	// Then the directory "src" exists (relative to the tempory target directory)
	if _, err := os.Stat(filepath.Join(targetDir, "src")); os.IsNotExist(err) {
		t.Errorf("Directory 'src' does not exist")
	}
	
	// And the directory "src/app" exists
	if _, err := os.Stat(filepath.Join(targetDir, "src", "app")); os.IsNotExist(err) {
		t.Errorf("Directory 'src/app' does not exist")
	}
	
	// And the directory "src/app/shared" exists
	if _, err := os.Stat(filepath.Join(targetDir, "src", "app", "shared")); os.IsNotExist(err) {
		t.Errorf("Directory 'src/app/shared' does not exist")
	}
	
	// And the directory "src/app/shared/utils" exists
	if _, err := os.Stat(filepath.Join(targetDir, "src", "app", "shared", "utils")); os.IsNotExist(err) {
		t.Errorf("Directory 'src/app/shared/utils' does not exist")
	}
	
	// And the file "src/app/shared/utils/foo.feature" exists
	if _, err := os.Stat(filepath.Join(targetDir, "src", "app", "shared", "utils", "foo.feature")); os.IsNotExist(err) {
		t.Errorf("File 'src/app/shared/utils/foo.feature' does not exist")
	}
}
