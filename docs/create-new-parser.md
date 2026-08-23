# Create new parser

This document describes the steps that are needed to introduce a new parser.

Each supported test framework (e.g. Jasmine, Vitest) has its own parser package under `parsers/<name>`, structurally separate from every other parser. `parsers/vitest` is the reference implementation to look at when adding another one.

## Procedure

1. **Create the ANTLR grammar.** Add `parsers/<name>/antlr/SpecGrammar.g4`, describing a tolerant "island" grammar that recognizes only the constructs you care about (e.g. the framework's suite/test keywords) plus enough filler rules to keep parentheses and braces balanced. `parsers/vitest/antlr/SpecGrammar.g4` (derived from `parsers/javascript/antlr/SpecGrammar.g4`) is a good template.
2. **Add the generator directive.** Add `parsers/<name>/antlr/generate.go` with the AGPL header, `package specantlr`, and the following directive:

   ```go
   //go:generate antlr -Dlanguage=Go -no-listener -visitor -package specantlr -o . SpecGrammar.g4
   ```
3. **Generate and commit the parser.** Run `go generate ./...` from the repository root and commit the generated `specgrammar_*.go`, `.interp` and `.tokens` files. These generated files must **not** be hand-edited and do **not** get an AGPL header — to change them, edit `SpecGrammar.g4` and re-run `go generate`.
4. **Implement the parser package.** In `parsers/<name>`, add:
   - `parser.go` with `func ParseSpecFile(path string, relaxed bool) ([]*messages.GherkinDocument, error)`, mirroring the existing parsers: build the ANTLR lexer/parser with a `collectingErrorListener` so partially-valid files still return whatever was parsed, alongside a non-nil error.
   - `parser.go` also exposes `func SupportedExtensions() []string`, the list of file suffixes (e.g. `.spec.ts`) this parser should be applied to when walking a directory.
   - `visitor.go`, a small ANTLR tree visitor that walks the parse tree into a `*frame` tree (or an equivalent intermediate representation), extracting Gherkin steps from `// Given` / `// When` / `// Then` / `// And` / `// But` line comments.
   - `emitter.go`, converting the frame tree into `[]*messages.GherkinDocument`, including URI normalization (kebab-casing segments, stripping the framework's file suffix).
   - Every hand-written Go file carries the AGPL header (`Copyright (C) 2026 <author>`), matching the existing files.
5. **Add test fixtures and tests.** Add representative spec files under `parsers/<name>/testdata/`, and a `parser_test.go` covering every fixture, following the `TestXxx_lowercase_suffix` naming convention and putting each Given/When/Then step as a comment directly above the corresponding Go code. Do not add tests for the `cmd` package.
6. **Register the parser.** In `cmd/rev.go`, add a case to the `lookupParser(name string)` switch mapping the new `--parser` flag value to a `specParser{name, extensions, parse}` built from your package's `ParseSpecFile` and `SupportedExtensions`.
7. **Update the documentation.** Mention the new `--parser` value in the `slicer rev` usage section of [../README.md](../README.md).
8. **Verify.** Run `go generate ./...` and `go test ./...` from the repository root to confirm the tree is clean and every parser (old and new) still passes.
