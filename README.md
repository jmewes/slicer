# Slicer

Slicer is a command-line program that can generate Gherkin feature files from existing test suites and, coming soon, specification documents from feature files. This might enable a software development approach where the source code gets auto-generated for the most part, but the specifications are still written by hand. It could also be useful for recreating existing programs in other programming languages or frameworks. And to look closely at a project's test suite to make it clearer.

## Usage

### Dependencies

The program is distributed as a binary file, so after downloading it, no further dependencies are required. The program can be executed via a terminal emulator like [Windows Terminal](https://github.com/microsoft/terminal), [iTerm2](https://iterm2.com/) or [GNOME Terminal](https://github.com/GNOME/gnome-terminal).

### Installation

#### Windows

Download the latest release `.exe` from the [releases page](https://github.com/jmewes/slicer/releases), then verify the checksum before running it:

```powershell
$version = "0.2.0"
$binary = "slicer-$version-windows-amd64.exe"
$expected = "e67311fe0c29f8eb7d0db31f56ecd2a98458f9973d55897d286a085d5f6a078b"

Invoke-WebRequest -Uri "https://github.com/jmewes/slicer/releases/download/$version/$binary" -OutFile $binary

$actual = (Get-FileHash -Path $binary -Algorithm SHA256).Hash.ToLower()

if ($expected -eq $actual) {
    Write-Output "Checksum OK"
} else {
    Write-Error "Checksum mismatch. Do not run this file!"
}
```

### Create feature files

`slicer rev` reverse-engineers Gherkin feature files from an existing test suite:

```sh
slicer rev --parser <jasmine|vitest> --source <path> --target <path> [--relaxed]
```

## Development

### Dependencies

For the development of the project, the following tools are needed:

- GitHub
- Git
- ANTLR

<details>
<summary>Setup Ubuntu</summary>

<br />

Download the ANTLR Jar file into the `.tools` directory:

```sh
mkdir -p .tools
curl -sSL -o .tools/antlr-4.13.2-complete.jar \
    https://www.antlr.org/download/antlr-4.13.2-complete.jar
```

</details>

<details>
<summary>Setup macOS</summary>

<br />

```sh
brew install git
brew install antlr
```

</details>

### Run tests

The unit tests of the project can be executed with the following command:

```sh
go test ./...
```

### Introducing a new parser

Each framework (e.g. Jasmine, vitest) has its own parser. So, if support for further frameworks should be added, a new parser should be created.

See [docs/create-new-parser.md](./docs/create-new-parser.md) for details.

### Regenerating the ANTLR spec parser

The `parsers` directory has a subdirectory for each supported source language (currently `javascript` for Jasmine and `vitest` for Vitest). There is a go-package for the respective language and another subdirectory for the ANTLR-generated parser.

To re-generate the ANTLR-generated parsers, execute the following command:

```sh
go generate ./parsers/...
```

## Integration tests

The **Slicer** program can be executed without full compilation by running its `main.go` file using the `go` program:

```sh
TEMP_DIR=$(mktemp -d)
go run main.go rev --parser jasmine --source "$SOURCE_DIR" --target "$TEMP_DIR" && code $TEMP_DIR
```

### Installation from source code

Instead of downloading the binary files from the GitHub releases, the `slicer` binary file can be compiled and installed from the source code, by running the following command:

```sh
go install
```

## Maintenance

### Create a new release

Creating new releases is done via the GitHub website:

https://github.com/jmewes/slicer/releases/new

Creating the release should automatically trigger the creation of the binary files via GitHub Actions:

https://github.com/jmewes/slicer/actions

As the final result, the GitHub Actions create a pull request for including the new release in the project's README file:

https://github.com/jmewes/slicer/pulls

### Open GitHub issue in the browser

If the GitHub CLI tool `gh` is installed, the description for a ticket can be opened like this:

```sh
ISSUE_NUMBER=10
gh issue view ${ISSUE_NUMBER} --json url
```

## Alternatives

**Test execution engines**

- [Cucumber](https://cucumber.io/)

**Structured-Prompt-Driven Developmnet (SPDD)**

- [Structured-Prompt-Driven Development (SPDD) | Wei Zhang, Jessie Jie Xia | martinfowler.com](https://martinfowler.com/articles/structured-prompt-driven)

**Specification-driven development (SDD)**

- [Understanding Spec-Driven-Development: Kiro, spec-kit, and Tessl | Birgitta Böckeler | martinfowler.com](https://martinfowler.com/articles/exploring-gen-ai/sdd-3-tools.html)
- [GitHub Spec Kit](https://github.github.com/spec-kit/)
- [OpenSpec](https://openspec.dev/)
- [BMAD Method](https://docs.bmad-method.org/)
- [Spec-Driven Development | addyosmani/agent-skills](https://github.com/addyosmani/agent-skills/blob/main/skills/spec-driven-development/SKILL.md)
- [dhinojosa/oreilly-spec-driven-development](https://github.com/dhinojosa/oreilly-spec-driven-development/blob/main/AGENTS.md) is an example for definiting Gherkin feature files in a separate cucumber acceptance test module. This is similar to the approach suggested to be used with `slicer`, but `slicer` assumes the Gherkin feature files embedded in the tests, with the encoding of the respective testing framework. Another difference is that with `slicer`, there is not so much of a need for a planning mode. The planning is done by the developer. The agent might just be requested to do the coding. Then the developer refactors the AI generated code. In so small cycles that instructions like the "five-minute limit" are not needed.

**Agentic software development frameworks**

- [adSCAILE](https://www.adesso.de/en/impulse/adscaile/index.jsp)
- [AI Unified Process (AIUP)](https://unifiedprocess.ai/)

**Annotated Textual Descriptions of Processes (ATDP)**

- [Process Extraction from Text | Patrizio Bellan et al. | arxiv.org](https://arxiv.org/pdf/2110.03754)

**Literate programming**

- [Knuth on Literate Programming | Turing Awardee Clips | youtube.com](https://www.youtube.com/watch?v=Mr3WTR0a5SM)

## Credits

- The Given/When/Then notation originates from the concept of Behavior-Driven-Development (BDD) invented by Daniel Terhorst-North and Chris Matts (see [martinfowler.com](https://martinfowler.com/bliki/GivenWhenThen.html)). 
- The Gherkin language is a formalization of the Given/When/Then notation invented by Aslak Hellesøy for the [Cucumber](https://cucumber.io) test execution engine (see [infoq.com](https://www.infoq.com/news/2018/04/cucumber-bdd-ten-years/)).
- A secondary goal of this project is to explore the benefits and limits of agentic coding. The original proof-of-concept has been generated with JetBrains Junie and the directly provided models. The ongoing development is done using JetBrains Junie and Claude Code, using LLMs via the local models from the [adesso ai hub](https://www.adesso.de/en/technologies/adesso-business-cloud/ai-hub.jsp) and JetBrains AI. The agents are requested to auto-commit their changes. So, the parts of the code that have been auto-generated by which agent using which LLM can be traced back in the commit history of the pull requests.

## References

- https://cucumber.io/docs/gherkin/reference/
- https://github.com/cucumber/gherkin
