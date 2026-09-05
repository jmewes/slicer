# Docker Sandbox

## Login

In the Docker account, an access token can be created and used in place of a password to log in to `sbx` like this:

```sh
echo "<personal access token>" | sbx login --username $docker_account --password-stdin
```

## Claude

### Environment file

To use Claude via a custom LLM gateway, one may define a `.env.sbx` file like this:

```sh
cat > .env.sbx <<'EOF'
ANTHROPIC_BASE_URL=https://llm-gateway.example.com
ANTHROPIC_AUTH_TOKEN=your-token
EOF

chmod 600 .env.sbx
```

### Start sandbox

With the following command, a pre-defined Docker Sandbox for Claude can be started:

```sh
sbx run --env-file .env.sbx claude
```

## Junie

### Validate the kit

```sh
sbx kit validate ./.junie/docker-sandbox
```

### Launch the sandbox

```sh
sbx run --kit ./.junie/docker-sandbox junie
```

## Troubleshooting

These commands may be useful to troubleshoot the Docker Sandbox:

```sh
sbx diagnose
sbx prune
sbx reset
```

## References
 
- https://www.docker.com/products/docker-sandboxes/
