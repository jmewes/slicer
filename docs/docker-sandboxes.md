# Agentic coding with Docker Sandboxes

This document describes how to run a coding agent (Claude Code or JetBrains Junie) against this repository inside a [Docker Sandbox](https://www.docker.com/products/docker-sandboxes/) — an isolated microVM with its own filesystem, network and Docker daemon. The agent can freely install packages, run `go generate`/`go test`, and use Git without touching the host machine. Only the project's working tree is shared with the sandbox; everything else the agent does (installed tools, processes, containers) disappears once the sandbox is removed.

## Prerequisites

1. Install the `sbx` CLI (see the [official docs](https://docs.docker.com/ai/sandboxes/) for the latest instructions):

   ```sh
   # macOS
   brew install docker/sandboxes/sbx

   # Windows (PowerShell)
   winget install Docker.Sandboxes

   # Linux
   curl -fsSL https://sbx.docker.com/install.sh | bash
   ```

2. Sign in:

   ```sh
   sbx login
   ```

3. Choose a network policy. **Balanced** is recommended for this project — it allows common package registries and API endpoints while still blocking arbitrary outbound traffic:

   ```sh
   sbx policy set balanced
   ```

   Running `go generate ./...` (which downloads ANTLR) and talking to the Anthropic/Junie API endpoints requires outbound network access. If a request gets blocked, allow-list the missing domain (see [Common workflow](#common-workflow) below) instead of switching to the "Open" policy.

## Running Claude Code

Claude Code ships as a built-in sandbox template, so no custom setup is needed beyond providing a key.

1. Store a custom Anthropic-compatible key. This works for a direct Anthropic key as well as a custom/BYOK endpoint (e.g. a company-hosted proxy) via `ANTHROPIC_BASE_URL`:

   ```sh
   sbx secret set anthropic --env ANTHROPIC_API_KEY
   # or, for a custom base URL:
   sbx secret set anthropic --env ANTHROPIC_AUTH_TOKEN --extra ANTHROPIC_BASE_URL=https://my-proxy.example.com
   ```

   The key is stored on the host and injected into the sandbox at runtime — it is never written to the sandbox's filesystem.

2. Launch a sandbox for this repository:

   ```sh
   sbx run --name slicer-claude claude
   ```

   This mounts the current working tree read-write into the sandbox and starts Claude Code.

3. Project-specific setup: the sandbox needs a JRE and the ANTLR tool to run `go generate` (see [docs/create-new-parser.md](./create-new-parser.md)). Ask the agent (or run manually inside the sandbox shell) to install them once per sandbox session:

   ```sh
   sudo apt-get update && sudo apt-get install -y default-jre-headless
   sudo curl -fsSL -o /usr/local/lib/antlr.jar https://www.antlr.org/download/antlr-4.13.2-complete.jar
   sudo printf '#!/bin/sh\nexec java -jar /usr/local/lib/antlr.jar "$@"\n' | sudo tee /usr/local/bin/antlr > /dev/null
   sudo chmod +x /usr/local/bin/antlr
   ```

4. Run the project's checks inside the sandbox:

   ```sh
   go generate ./...
   go test ./...
   ```

5. Review the agent's changes from the host, using the shared working tree:

   ```sh
   git diff
   ```

## Running Junie

Junie is not one of Docker Sandboxes' built-in agent templates (only Claude Code, Codex, Copilot, Cursor, Docker Agent, Droid, Gemini, Kiro, OpenCode and Shell ship out of the box). This repository provides a small custom **kit** at [`docker-sandboxes/junie/spec.yaml`](../docker-sandboxes/junie/spec.yaml) that extends the generic `shell` template, installs the Junie CLI, and wires up a custom API key.

1. Validate the kit (the kit format is an early-access feature, so it's worth checking it against your installed `sbx` version before relying on it):

   ```sh
   sbx kit validate ./docker-sandboxes/junie/
   ```

2. Store the Junie API key. Junie supports both models from the JetBrains AI Hub and custom models behind an OpenAI-compatible API, so the key/endpoint pair is stored as a custom secret:

   ```sh
   sbx secret set-custom junie --env JUNIE_API_KEY --extra JUNIE_BASE_URL=https://your-openai-compatible-endpoint.example.com
   ```

   As with Claude Code, the real key stays on the host and is injected into the sandbox only at runtime.

3. Launch the sandbox using the local kit:

   ```sh
   sbx run --kit ./docker-sandboxes/junie/ junie
   ```

4. Follow the same project-specific setup (JRE + ANTLR, see step 3 above) and run `go generate ./...` / `go test ./...` as needed.

## Common workflow

- List running sandboxes:

  ```sh
  sbx ls
  ```

- Stop a sandbox (keeps it around for a later restart):

  ```sh
  sbx stop slicer-claude
  ```

- Remove a sandbox entirely once you're done:

  ```sh
  sbx rm slicer-claude
  ```

  Removing a sandbox discards everything created inside it (installed packages, containers, shell state). The only thing that survives is the host's working tree, so commit or otherwise persist anything you want to keep before removing the sandbox.

- If the network policy blocks a domain the agent needs (e.g. ANTLR's download server, or the Anthropic/Junie API endpoint), allow-list it explicitly instead of loosening the whole policy:

  ```sh
  sbx policy allow network antlr.org
  ```

- Agents running inside a sandbox must still follow the auto-commit rule from [`CLAUDE.md`](../CLAUDE.md) / [`.junie/AGENTS.md`](../.junie/AGENTS.md): every modified, added or deleted file must be staged and committed before the agent finishes its turn.
