package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	messages "github.com/cucumber/messages/go/v28"
	"github.com/experimental-software/gherkin/core"
	javascript "github.com/experimental-software/gherkin/parsers/javascript"
	"github.com/spf13/cobra"
)

var RelaxedOption bool

var revCmd = &cobra.Command{
	Use:   "rev [flags]",
	Short: "Reverse engineer feature files from source code",

	Run: func(cmd *cobra.Command, args []string) {
		docs, err := parseSpecSources(SourceParameter, RelaxedOption)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		err = core.WriteFeatureFiles(docs, TargetParameter)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Wrote %d feature files to %s\n", len(docs), TargetParameter)
	},
}

func parseSpecSources(source string, relaxed bool) ([]*messages.GherkinDocument, error) {
	info, err := os.Stat(source)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return javascript.ParseSpecFile(source, relaxed)
	}

	var docs []*messages.GherkinDocument
	err = filepath.WalkDir(source, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".spec.ts") {
			return nil
		}

		parsedDocs, parseErr := javascript.ParseSpecFile(path, relaxed)
		// Use any documents the parser managed to produce even when it
		// reports an error, so partially-invalid spec files still yield
		// their recognizable features instead of aborting the walk.
		docs = append(docs, parsedDocs...)
		if parseErr != nil && len(parsedDocs) == 0 {
			return parseErr
		}
		if parseErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s: %v\n", path, parseErr)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func init() {
	revCmd.Flags().StringVarP(&SourceParameter, "source", "s", "", "Path to source file or directory")
	revCmd.Flags().StringVarP(&TargetParameter, "target", "t", "", "Path to target directory")
	_ = revCmd.MarkFlagRequired("source")
	_ = revCmd.MarkFlagRequired("target")

	revCmd.Flags().BoolVar(&RelaxedOption, "relaxed", false, "Relaxed mode")

	rootCmd.AddCommand(revCmd)
}
