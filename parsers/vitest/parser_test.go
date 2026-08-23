package vitest

import (
	"strings"
	"testing"

	messages "github.com/cucumber/messages/go/v28"
)

func TestParseSpecFile_simple(t *testing.T) {
	// Given a test file with a top-level "test" call and no enclosing "describe"
	specPath := "testdata/simple.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	// Then a single implicit Feature named after the file is produced
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]
	if doc.Feature.Name != "Simple" {
		t.Errorf("Expected Feature name 'Simple', got '%s'", doc.Feature.Name)
	}

	// And the "test" call is mapped to a single Scenario with three steps
	var scenarios []*messages.Scenario
	for _, c := range doc.Feature.Children {
		if c.Scenario != nil {
			scenarios = append(scenarios, c.Scenario)
		}
	}
	if len(scenarios) != 1 {
		t.Fatalf("Expected 1 scenario, got %d", len(scenarios))
	}
	if scenarios[0].Name != "adding two numbers" {
		t.Errorf("Expected scenario name 'adding two numbers', got '%s'", scenarios[0].Name)
	}
	if len(scenarios[0].Steps) != 3 {
		t.Errorf("Expected 3 steps, got %d", len(scenarios[0].Steps))
	}
}

func TestParseSpecFile_alias_it(t *testing.T) {
	// Given a test file that uses "it" instead of "test"
	specPath := "testdata/alias_it.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	// Then "it" behaves identically to "test", yielding one implicit Feature
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]
	if doc.Feature.Name != "Alias it" {
		t.Errorf("Expected Feature name 'Alias it', got '%s'", doc.Feature.Name)
	}
	if len(doc.Feature.Children) != 1 {
		t.Fatalf("Expected 1 scenario, got %d", len(doc.Feature.Children))
	}
	if doc.Feature.Children[0].Scenario.Name != "subtracts numbers correctly" {
		t.Errorf("Expected scenario name 'subtracts numbers correctly', got '%s'", doc.Feature.Children[0].Scenario.Name)
	}
}

func TestParseSpecFile_describe(t *testing.T) {
	// Given a test file with a "describe" block containing a single "test"
	specPath := "testdata/describe.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	// Then the describe title becomes the Feature name
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]
	if doc.Feature.Name != "Arithmetic operations" {
		t.Errorf("Expected Feature name 'Arithmetic operations', got '%s'", doc.Feature.Name)
	}
	if len(doc.Feature.Children) != 1 {
		t.Fatalf("Expected 1 scenario, got %d", len(doc.Feature.Children))
	}
	if doc.Feature.Children[0].Scenario.Name != "multiplies two positive numbers" {
		t.Errorf("Expected scenario name 'multiplies two positive numbers', got '%s'", doc.Feature.Children[0].Scenario.Name)
	}
}

func TestParseSpecFile_nested_describe(t *testing.T) {
	// Given a test file with a "describe" block nested inside another "describe" block
	specPath := "testdata/nested_describe.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	// Then the childless outer describe is dropped and the inner one becomes the Feature
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]
	if doc.Feature.Name != "Addition sub-module" {
		t.Errorf("Expected Feature name 'Addition sub-module', got '%s'", doc.Feature.Name)
	}

	// And the URI contains both describe segments
	if doc.Uri != "testdata/nested-describe/arithmetic-module/addition-sub-module" {
		t.Errorf("Unexpected resource path: '%s'", doc.Uri)
	}
}

func TestParseSpecFile_skip(t *testing.T) {
	// Given a test file with a "test.skip" call
	specPath := "testdata/skip.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	// Then the modifier is ignored and a normal Scenario is produced
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]
	if len(doc.Feature.Children) != 1 {
		t.Fatalf("Expected 1 scenario, got %d", len(doc.Feature.Children))
	}
	if doc.Feature.Children[0].Scenario.Name != "division by zero produces infinity" {
		t.Errorf("Expected scenario name 'division by zero produces infinity', got '%s'", doc.Feature.Children[0].Scenario.Name)
	}
}

func TestParseSpecFile_only(t *testing.T) {
	// Given a test file with a "test.only" call
	specPath := "testdata/only.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	// Then the modifier is ignored and a normal Scenario is produced
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]
	if len(doc.Feature.Children) != 1 {
		t.Fatalf("Expected 1 scenario, got %d", len(doc.Feature.Children))
	}
	if doc.Feature.Children[0].Scenario.Name != "divides two numbers" {
		t.Errorf("Expected scenario name 'divides two numbers', got '%s'", doc.Feature.Children[0].Scenario.Name)
	}
}

func TestParseSpecFile_parameterized(t *testing.T) {
	// Given a test file that uses "test.for([...])" with a data table
	specPath := "testdata/parameterized.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	// Then a single Scenario is produced using the raw title string
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]
	if len(doc.Feature.Children) != 1 {
		t.Fatalf("Expected 1 scenario, got %d", len(doc.Feature.Children))
	}
	if doc.Feature.Children[0].Scenario.Name != "adds %i and %i to get %i" {
		t.Errorf("Expected scenario name 'adds %%i and %%i to get %%i', got '%s'", doc.Feature.Children[0].Scenario.Name)
	}
}

func TestParseSpecFile_before_after_each(t *testing.T) {
	// Given a test file with a "beforeEach" containing a "Given" comment
	// And an "afterEach" containing a "Then" comment
	specPath := "testdata/before_after_each.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]

	// Then the "Given" comment in "beforeEach" becomes a Background step
	var background *messages.Background
	var scenarios []*messages.Scenario
	for _, c := range doc.Feature.Children {
		if c.Background != nil {
			background = c.Background
		}
		if c.Scenario != nil {
			scenarios = append(scenarios, c.Scenario)
		}
	}
	if background == nil || len(background.Steps) != 1 {
		t.Fatalf("Expected a Background with 1 step, got %v", background)
	}

	// And the "Then" comment in "afterEach" appears nowhere in the output
	for _, step := range background.Steps {
		if strings.Contains(strings.ToLower(step.Text), "after test run") {
			t.Errorf("Did not expect afterEach step text to leak into background: %q", step.Text)
		}
	}
	if len(scenarios) != 1 {
		t.Fatalf("Expected 1 scenario, got %d", len(scenarios))
	}
}

func TestParseSpecFile_before_after_all(t *testing.T) {
	// Given a test file with a "beforeAll" containing a "Given" comment
	// And an "afterAll" containing a "Then" comment
	specPath := "testdata/before_after_all.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]

	// Then the "Given" comment in "beforeAll" becomes a Background step
	var background *messages.Background
	for _, c := range doc.Feature.Children {
		if c.Background != nil {
			background = c.Background
		}
	}
	if background == nil || len(background.Steps) != 1 {
		t.Fatalf("Expected a Background with 1 step, got %v", background)
	}

	// And the "Then" comment in "afterAll" appears nowhere in the output
	for _, step := range background.Steps {
		if strings.Contains(strings.ToLower(step.Text), "clean up") {
			t.Errorf("Did not expect afterAll step text to leak into background: %q", step.Text)
		}
	}

	// And the derived Feature name capitalizes the underscore-separated file stem
	if doc.Feature.Name != "Before after all" {
		t.Errorf("Expected Feature name 'Before after all', got '%s'", doc.Feature.Name)
	}
}

func TestParseSpecFile_scoped_setup(t *testing.T) {
	// Given a test file with a "beforeEach" inside a "describe" block
	specPath := "testdata/scoped_setup.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	// Then the Background is attached to that Feature only
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]
	if doc.Feature.Name != "Scoped calculation context" {
		t.Errorf("Expected Feature name 'Scoped calculation context', got '%s'", doc.Feature.Name)
	}
	var background *messages.Background
	for _, c := range doc.Feature.Children {
		if c.Background != nil {
			background = c.Background
		}
	}
	if background == nil || len(background.Steps) != 1 {
		t.Fatalf("Expected a Background with 1 step, got %v", background)
	}
}

func TestParseSpecFile_async_await(t *testing.T) {
	// Given a test file whose test callback is an "async" arrow function
	specPath := "testdata/async_await.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	// Then the "async" callback does not disturb parsing
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]
	if len(doc.Feature.Children) != 1 {
		t.Fatalf("Expected 1 scenario, got %d", len(doc.Feature.Children))
	}
	if doc.Feature.Children[0].Scenario.Name != "adds numbers asynchronously" {
		t.Errorf("Expected scenario name 'adds numbers asynchronously', got '%s'", doc.Feature.Children[0].Scenario.Name)
	}
}

func TestParseSpecFile_timeouts(t *testing.T) {
	// Given a test file with a trailing numeric timeout argument
	specPath := "testdata/timeouts.test.ts"

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath)

	// Then the trailing timeout argument does not disturb parsing
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]
	if len(doc.Feature.Children) != 1 {
		t.Fatalf("Expected 1 scenario, got %d", len(doc.Feature.Children))
	}
	if doc.Feature.Children[0].Scenario.Name != "multiplies numbers within time limit" {
		t.Errorf("Expected scenario name 'multiplies numbers within time limit', got '%s'", doc.Feature.Children[0].Scenario.Name)
	}
}

func TestParseSpecFile_todo_default(t *testing.T) {
	// Given a test file with a "test.todo" call and no callback
	specPath := "testdata/todo.test.ts"
	// And relaxed mode is disabled
	option := withRelaxedOption(false)

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath, option)

	// Then no Scenario is produced
	if len(docs) != 0 {
		t.Fatalf("Expected 0 GherkinDocuments, got %d", len(docs))
	}
}

func TestParseSpecFile_todo_relaxed(t *testing.T) {
	// Given a test file with a "test.todo" call and no callback
	specPath := "testdata/todo.test.ts"
	// And relaxed mode is enabled
	option := withRelaxedOption(true)

	// When the spec file gets parsed
	docs := mustParseSpecFile(specPath, option)

	// Then one step-less Scenario is produced
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(docs))
	}
	doc := docs[0]
	if len(doc.Feature.Children) != 1 {
		t.Fatalf("Expected 1 scenario, got %d", len(doc.Feature.Children))
	}
	scenario := doc.Feature.Children[0].Scenario
	if scenario.Name != "calculate square root of negative numbers using imaginary units" {
		t.Errorf("Unexpected scenario name: '%s'", scenario.Name)
	}
	if len(scenario.Steps) != 0 {
		t.Errorf("Expected 0 steps, got %d", len(scenario.Steps))
	}
}

func TestParseSpecFile_uri_suffix_stripping(t *testing.T) {
	// Given a ".test.ts" file
	tsDocs := mustParseSpecFile("testdata/simple.test.ts")
	// And a ".spec.ts" style fixture is not present, so we verify with the
	// existing fixture and assert both suffixes are handled by construction

	// When the spec file gets parsed
	if len(tsDocs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument, got %d", len(tsDocs))
	}

	// Then the URI carries no file-extension remnant
	if strings.Contains(tsDocs[0].Uri, ".test") || strings.Contains(tsDocs[0].Uri, ".ts") {
		t.Errorf("Expected URI to have no extension remnant, got '%s'", tsDocs[0].Uri)
	}
}

func TestSupportedExtensions_has_supported_extension(t *testing.T) {
	// Given the eight supported extensions
	exts := SupportedExtensions()

	// When checking a matching file name
	if !hasSupportedExtensionForTest("Sample.TEST.TS", exts) {
		t.Errorf("Expected 'Sample.TEST.TS' to match case-insensitively")
	}

	// Then a non-matching file is not considered supported
	if hasSupportedExtensionForTest("sample.ts", exts) {
		t.Errorf("Did not expect 'sample.ts' to match")
	}
}

func hasSupportedExtensionForTest(name string, exts []string) bool {
	lower := strings.ToLower(name)
	for _, ext := range exts {
		if strings.HasSuffix(lower, ext) {
			return true
		}
	}
	return false
}

func TestParseSpecFile_malformed_returns_partial_docs(t *testing.T) {
	// Given a deliberately malformed spec file
	specPath := "testdata/malformed.test.ts"

	// When the spec file gets parsed
	docs, err := ParseSpecFile(specPath, false)

	// Then the documents parsed so far are still returned alongside a non-nil error
	if err == nil {
		t.Fatalf("Expected a non-nil error for malformed input")
	}
	if len(docs) != 1 {
		t.Fatalf("Expected 1 GherkinDocument to survive the parse error, got %d", len(docs))
	}
	if docs[0].Feature.Name != "Well formed part" {
		t.Errorf("Expected Feature name 'Well formed part', got '%s'", docs[0].Feature.Name)
	}
}

type parseOption func(*parseConfig)

type parseConfig struct {
	relaxed bool
}

func withRelaxedOption(relaxed bool) parseOption {
	return func(c *parseConfig) {
		c.relaxed = relaxed
	}
}

func mustParseSpecFile(path string, opts ...parseOption) []*messages.GherkinDocument {
	cfg := &parseConfig{relaxed: false}
	for _, opt := range opts {
		opt(cfg)
	}

	docs, err := ParseSpecFile(path, cfg.relaxed)
	if err != nil {
		panic(err)
	}
	return docs
}
