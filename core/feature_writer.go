package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	messages "github.com/cucumber/messages/go/v28"
)

func CalculateFeaturePath(originalUri string) string {
	parts := strings.Split(originalUri, "/")
	var beforeParts []string
	var afterParts []string
	foundSpec := false
	for _, part := range parts {
		if strings.HasSuffix(part, ".spec.ts") {
			foundSpec = true
			continue
		}
		if foundSpec {
			afterParts = append(afterParts, part)
		} else {
			beforeParts = append(beforeParts, part)
		}
	}
	result := strings.Join(beforeParts, "/")
	if len(afterParts) > 0 {
		if result != "" {
			result += "/"
		}
		result += strings.Join(afterParts, ".")
	}
	return result
}

func WriteFeatureFiles(docs []*messages.GherkinDocument, targetDir string) error {
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

		uri := CalculateFeaturePath(doc.Uri)
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

// sanitizeFeatureFileBaseName returns a filesystem-safe base name for a feature
// file by replacing whitespace characters with underscores.
func sanitizeFeatureFileBaseName(name string) string {
	return strings.ReplaceAll(name, " ", "_")
}

// featureFilePathPartsFromURI splits a slash-separated feature URI into the
// relative directory and the base name (the last segment).
func featureFilePathPartsFromURI(uri string) (string, string) {
	idx := strings.LastIndex(uri, "/")
	if idx < 0 {
		return "", uri
	}
	return uri[:idx], uri[idx+1:]
}

func renderFeatureDocument(doc *messages.GherkinDocument) string {
	if doc == nil || doc.Feature == nil {
		return ""
	}

	var b strings.Builder
	b.WriteString("Feature: ")
	b.WriteString(doc.Feature.Name)
	b.WriteString("\n")

	for _, child := range doc.Feature.Children {
		if child == nil {
			continue
		}

		if child.Background != nil {
			b.WriteString("\n  Background:\n")
			for _, step := range child.Background.Steps {
				if step == nil {
					continue
				}
				b.WriteString("    ")
				b.WriteString(strings.TrimSpace(step.Keyword))
				b.WriteString(" ")
				b.WriteString(step.Text)
				b.WriteString("\n")
			}
		}

		if child.Scenario != nil {
			b.WriteString("\n  Scenario: ")
			b.WriteString(child.Scenario.Name)
			b.WriteString("\n")
			for _, step := range child.Scenario.Steps {
				if step == nil {
					continue
				}
				b.WriteString("    ")
				b.WriteString(strings.TrimSpace(step.Keyword))
				b.WriteString(" ")
				b.WriteString(step.Text)
				b.WriteString("\n")
			}
		}
	}

	b.WriteString("\n")

	return b.String()
}
