
Create a plan for the introduction of a new parser for "vitest"!

It should be structurally equivalent to the existing "javascript" parser (specialised on the Jasmine framework), but strictly separated.

## vitest grammar

(see https://vitest.dev/guide/learn/writing-tests.html and https://vitest.dev/guide/learn/writing-tests-with-ai.html for details)

For each of the vitest testsuite features listed below, create a separate test file and a separate test case.

See the test files in "vitest/testdata" for the vitest test suite structures that need to be supported.

## Test file naming conventions

The following test file pattern should be supported:

1. *.test.ts / *.test.js: e.g. utils.test.ts, utils.test.js
2. *.spec.ts / *.spec.js: Scaffold.spec.tsx, Scaffold.spec.ts

## Command-line flags

For the `slicer rev` command, add a flag `--parser` (short: `-p`) where one of the following options has to be provided `jasmine`, `vitest`.

## Documentation

### Slicer CLI usage

Add a basic user guide to the "usage" section of the README.

### Adding a new parser

In the file @docs/create-new-parser.md, add documentation for the steps that will need to be executed with later adding more parsers (e.g. Deno, JUnit).
