package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	messages "github.com/cucumber/messages/go/v28"
	"github.com/experimental-software/gherkin/core"
)

func writeFeatureFiles(docs []*messages.GherkinDocument, targetDir string) error {
	if targetDir == "" {
		return fmt.Errorf("target parameter is required")
	}

	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}

	for _, doc := range docs {
		if doc == nil {
			continue
		}

		uri := core.CalculateFeaturePath(doc.Uri)
		relativeDir, baseName := featureFilePathPartsFromURI(uri)
		baseName = sanitizeFeatureFileBaseName(baseName)

		featurePath := filepath.Join(targetDir, filepath.FromSlash(relativeDir), baseName+".feature")

		if err := os.MkdirAll(filepath.Dir(featurePath), 0o755); err != nil {
			return err
		}

		if err := os.WriteFile(featurePath, []byte(renderFeatureDocument(doc)), 0o644); err != nil {
			return err
		}
	}

	return nil
}
