# Agentic coding with Docker Sandboxes

This document describes how to run a coding agent (Claude Code) against this repository inside a [Docker Sandbox](https://www.docker.com/products/docker-sandboxes/) — an isolated microVM with its own filesystem, network and Docker daemon. The agent can freely install packages, run `go generate`/`go test`, and use Git without touching the host machine. Only the project's working tree is shared with the sandbox; everything else the agent does (installed tools, processes, containers) disappears once the sandbox is removed.

## Prerequisites

https://www.docker.com/products/docker-sandboxes/

1. Install the `sbx` CLI (see the [official docs](https://docs.docker.com/ai/sandboxes/) for the latest instructions):

   ```sh
   # macOS
   brew trust docker/tap && brew install docker/tap/sbx
   ```

2. Sign in:

   ```sh
   sbx login
   ```

3. Choose a network policy. **Balanced** is recommended for this project — it allows common package registries and API endpoints while still blocking arbitrary outbound traffic:

   ```sh
   # Does not work as suggested:
   # sbx policy set balanced
   ```

   Running `go generate ./...` (which downloads ANTLR) and talking to the Anthropic/Junie API endpoints requires outbound network access. If a request gets blocked, allow-list the missing domain (see [Common workflow](#common-workflow) below) instead of switching to the "Open" policy.

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

## Troubleshooting

- **Junie fails to log in with "Secure connection to JetBrains services failed... the corporate CA certificate must be trusted by the Java runtime"**: Junie is a JVM-based CLI, but the JVM ships its own trust store that is separate from the sandbox's OS-level one. The sandbox's TLS-intercepting proxy certificate is trusted by the OS out of the box, but not by the JVM, so Junie's HTTPS calls to JetBrains fail. The [`docker-sandboxes/junie/spec.yaml`](../docker-sandboxes/junie/spec.yaml) kit installs `ca-certificates-java` and copies the resulting Java trust store over any other `cacerts` files found in the sandbox (e.g. a private JBR bundled by the Junie installer) so this happens automatically. If the error still occurs after re-creating the sandbox, run `sbx kit validate ./docker-sandboxes/junie/` to confirm the kit applied, or fall back to manually importing the CA into Junie's `cacerts` inside the sandbox (see Docker's [Install an internal CA certificate](https://docs.docker.com/ai/sandboxes/customize/kit-examples/) kit example for the general pattern).
