| :warning: WARNING                                |
|:-------------------------------------------------|
| This file contains unverified AI-generated text. |


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