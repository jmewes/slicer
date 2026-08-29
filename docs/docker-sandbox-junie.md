## Running Junie

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
