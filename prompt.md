Create a step-by-step implementation plan for adding a new "vitest" parser package in `parsers/vitest`.
## Architectural Requirements
- Implement `parsers/vitest` structurally separated from `parsers/javascript` (Jasmine), using ANTLR grammar under `parsers/vitest/antlr`.
- Support Vitest test structures demonstrated in `parsers/vitest/testdata/` (e.g. `describe`, `it`/`test`, `test.skip`, `test.only`, `test.todo`, `test.each`, `beforeAll`/`afterEach`, `async/await`).
## Supported File Patterns
Configure the Vitest parser to handle:
- `*.test.ts`, `*.test.tsx`, `*.test.js`, `*.test.jsx`
- `*.spec.ts`, `*.spec.tsx`, `*.spec.js`, `*.spec.jsx`
## CLI Flag & Integration (`cmd/rev.go`)
- Add madatory `--parser` / `-p` flag to `slicer rev`. Accepted values: `jasmine` (pointing to `parsers/javascript`), `vitest` (pointing to `parsers/vitest`).
- Refactor `parseSpecSources` in `cmd/rev.go` to filter source files based on the selected parser's supported extensions.
## Documentation Updates
- `README.md`: Update usage section with `--parser` flag options.
- `docs/create-new-parser.md`: Document step-by-step instructions for adding future parsers.
## Verification & Constraints
- Plan must include running `go generate ./...` and `go test ./...`.
- Ensure AGPL headers are present on new Go files (except generated ANTLR files).