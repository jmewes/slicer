# Slicer

**Slicer** is a commandline program that can generate [Gherkin feature-files](https://cucumber.io/docs/gherkin/reference/#feature) from existing test suites and specification documents from feature-files. This can be useful for extracting specifications from existing programs to re-create the code in other programming languages or other frameworks. It may also be useful to improve the quality of the test suite.

## Getting started / Installation

```sh
go install
```

## Development

### Run tests

```sh
go test ./...
```

### Regenerating the ANTLR spec parser

The `core/antlr` package contains an ANTLR-generated Go parser for a
minimal TypeScript spec-file island grammar (`SpecGrammar.g4`). The
generated `.go` files are committed, so building and testing the project
only requires the Go toolchain.

To regenerate them (only needed after editing the grammar), download the
ANTLR complete jar into `.tools/` and run `go generate`:

```sh
brew install antlr
```

```sh
go generate ./javascript_parser/antlr/...
```

This requires a JDK on the `PATH`. The jar is a dev-only dependency and
is not needed to build or run Slicer.

## Integration tests

```sh
TEMP_DIR=$(mktemp -d)
go run main.go rev --source "$SOURCE_DIR" --target "$TEMP_DIR" && code $TEMP_DIR
```

## Maintenance

### Open GitHub issue in the browser

If the GitHub CLI tool `gh` is installed, the description for a ticket can be opened like this:

```sh
ISSUE_NUMBER=10
gh issue view ${ISSUE_NUMBER} --json url
```

## Alternatives

**Test execution engines**

- [Cucumber](https://cucumber.io/)

**Development processes**

- [Structured-Prompt-Driven Development (SPDD) | Wei Zhang, Jessie Jie Xia | martinfowler.com](https://martinfowler.com/articles/structured-prompt-driven)
- [AI Unified Process (AIUP)](https://unifiedprocess.ai/)

**Annotated Textual Descriptions of Processes (ATDP)**

- [Process Extraction from Text | Patrizio Bellan et al. | arxiv.org](https://arxiv.org/pdf/2110.03754)

**Literate programming**

- [Knuth on Literate Programming | Turing Awardee Clips | youtube.com](https://www.youtube.com/watch?v=Mr3WTR0a5SM)

## Credits

- The Given/When/Then notation originates from the concept of Behavior-Driven-Development (BDD) invented by Daniel Terhorst-North and Chris Matts (see [martinfowler.com](https://martinfowler.com/bliki/GivenWhenThen.html)). 
- The Gherkin language is a formalization of the Given/When/Then notation invented by Aslak Hellesøy for the [Cucumber](https://cucumber.io) test execution engine (see [infoq.com](https://www.infoq.com/news/2018/04/cucumber-bdd-ten-years/)).
- A secondary goal of this project is to explore the benefits and limits of agentic coding. The original proof-of-concept has been generated with JetBrains Junie and the directly provided models. The ongoing development is done using JetBrains Junie and local models from the [adesso ai hub](https://www.adesso.de/en/technologies/adesso-business-cloud/ai-hub.jsp).

## References

- https://cucumber.io/docs/gherkin/reference/
- https://github.com/cucumber/gherkin
- https://marketplace.visualstudio.com/items?itemName=alexkrechik.cucumberautocomplete
