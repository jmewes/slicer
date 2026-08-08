# Slicer

**Slicer** is a command-line program that can generate [Gherkin feature files](https://cucumber.io/docs/gherkin/reference/#feature) from existing test suites and specification documents from feature files (coming soon). This can be useful for recreating existing programs in other programming languages or frameworks. It might also be useful to analyze a project’s test suite to improve its expressiveness.

## Usage

### Dependencies

The program is distributed as a binary file, so after downloading it, no further dependencies are required. The program can be executed via a terminal emulator like [Windows Terminal](https://github.com/microsoft/terminal), [iTerm2](https://iterm2.com/) or [GNOME Terminal](https://github.com/GNOME/gnome-terminal).

### Installation

#### Windows

Download the latest release `.exe` and its `.sha256` checksum file from the [releases page](https://github.com/jmewes/slicer/releases), then verify the checksum before running it:

```powershell
$version = "0.1.2-RC4"
$binary = "gherkin-$version-windows-amd64.exe"
$expected = "6b455b4051ff5f8149e21e82df6ff9e91c4f8556a9b8259296433b648a9719f7"

Invoke-WebRequest -Uri "https://github.com/jmewes/slicer/releases/download/$version/$binary" -OutFile $binary

$actual = (Get-FileHash -Path $binary -Algorithm SHA256).Hash.ToLower()

if ($expected -eq $actual) {
    Write-Output "Checksum OK"
} else {
    Write-Error "Checksum mismatch. Do not run this file!"
}
```

## Development

### Dependencies

For the development of the project, the following tools are needed:

- GitHub
- Git
- ANTRL

<details>
<summary>Setup macOS</summary>

<br />

```sh
brew install git
brew install antlr
```

</details>

### Run tests

```sh
go test ./...
```

## Integration tests

```sh
TEMP_DIR=$(mktemp -d)
go run main.go rev --source "$SOURCE_DIR" --target "$TEMP_DIR" && code $TEMP_DIR
```

### Installation from source code

```sh
go install
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

- [GitHub Spec Kit](https://github.github.com/spec-kit/)
- [OpenSpec](https://openspec.dev/)
- [BMAD Method](https://docs.bmad-method.org/)
- [Spec-Driven Development | addyosmani/agent-skills](https://github.com/addyosmani/agent-skills/blob/main/skills/spec-driven-development/SKILL.md)
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
