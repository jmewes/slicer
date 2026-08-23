package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	messages "github.com/cucumber/messages/go/v28"
	"github.com/jmewes/slicer/core"
	javascript "github.com/jmewes/slicer/parsers/javascript"
	"github.com/jmewes/slicer/parsers/vitest"
	"github.com/spf13/cobra"
)

var RelaxedOption bool
var ParserParameter string

// specParser bundles a parser's name, the file extensions it recognizes
// when walking a directory, and its ParseSpecFile entry point.
type specParser struct {
	name       string
	extensions []string
	parse      func(path string, relaxed bool) ([]*messages.GherkinDocument, error)
}

// lookupParser resolves the --parser flag value to a specParser. Unknown
// values return an error listing the accepted values.
func lookupParser(name string) (specParser, error) {
	switch name {
	case "jasmine":
		return specParser{
			name:       "jasmine",
			extensions: []string{".spec.ts"},
			parse:      javascript.ParseSpecFile,
		}, nil
	case "vitest":
		return specParser{
			name:       "vitest",
			extensions: vitest.SupportedExtensions(),
			parse:      vitest.ParseSpecFile,
		}, nil
	default:
		return specParser{}, fmt.Errorf("unknown parser %q, accepted values are: jasmine, vitest", name)
	}
}

var revCmd = &cobra.Command{
	Use:   "rev [flags]",
	Short: "Reverse engineer feature files from source code",

	Run: func(cmd *cobra.Command, args []string) {
		parser, err := lookupParser(ParserParameter)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		docs, err := parseSpecSources(SourceParameter, RelaxedOption, parser)
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

// hasSupportedExtension reports whether name ends with one of extensions,
// matched case-insensitively.
func hasSupportedExtension(name string, extensions []string) bool {
	lower := strings.ToLower(name)
	for _, ext := range extensions {
		if strings.HasSuffix(lower, ext) {
			return true
		}
	}
	return false
}

func parseSpecSources(source string, relaxed bool, p specParser) ([]*messages.GherkinDocument, error) {
	info, err := os.Stat(source)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return p.parse(source, relaxed)
	}

	var docs []*messages.GherkinDocument
	err = filepath.WalkDir(source, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !hasSupportedExtension(d.Name(), p.extensions) {
			return nil
		}

		parsedDocs, parseErr := p.parse(path, relaxed)
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

	revCmd.Flags().StringVarP(&ParserParameter, "parser", "p", "", "Spec parser to use (jasmine, vitest)")
	_ = revCmd.MarkFlagRequired("parser")

	revCmd.Flags().BoolVar(&RelaxedOption, "relaxed", false, "Relaxed mode")

	rootCmd.AddCommand(revCmd)
}
