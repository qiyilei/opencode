# OpenCode

A terminal-based AI coding assistant written in Go. OpenCode gives you an interactive TUI for chatting with LLMs, reading and editing your code, running shell commands, and pulling in diagnostics from language servers — all without leaving the terminal.

- Language: Go 1.24+
- License: MIT
- Interface: Bubble Tea TUI + one-shot non-interactive mode

## Features

- **Interactive TUI** built with Bubble Tea, with a Vim-like embedded editor.
- **Multiple providers**: Anthropic, OpenAI, Google Gemini, GitHub Copilot, Groq, OpenRouter, xAI, AWS Bedrock, Azure OpenAI, Vertex AI, and any OpenAI-compatible local endpoint.
- **Session management** backed by SQLite, with per-session token and cost tracking.
- **Tool use**: the assistant can glob/grep/ls, view, write, edit and patch files, run shell commands, fetch URLs, search Sourcegraph, and query LSP diagnostics.
- **MCP support**: attach external tools over stdio or SSE.
- **LSP integration**: start language servers, watch files for changes, and surface diagnostics to the model.
- **Custom commands**: Markdown-defined prompts with named `$ARGUMENTS`.
- **Image attachments**: attach screenshots and designs to your messages.
- **9 built-in themes**, adaptive to light and dark terminals.
- **Auto compact**: summarize the conversation before the context window fills up.
- **External editor support** for composing long messages.

## Requirements

- Go 1.24.0 or later to build from source.
- macOS or Linux for prebuilt binaries (`amd64` / `arm64`).
- An API key for at least one supported provider.

## Installation

Install script:

```bash
curl -fsSL https://raw.githubusercontent.com/opencode-ai/opencode/refs/heads/main/install | bash
```

Homebrew:

```bash
brew install opencode-ai/tap/opencode
```

Arch Linux (AUR):

```bash
yay -S opencode-ai-bin
```

Go:

```bash
go install github.com/opencode-ai/opencode@latest
```

From source:

```bash
git clone https://github.com/opencode-ai/opencode.git
cd opencode
go build -o opencode
./opencode
```

## Quick start

Export a provider key and start the TUI:

```bash
export ANTHROPIC_API_KEY=sk-ant-...
opencode
```

Run a single prompt and exit:

```bash
opencode -p "Explain how context cancellation works in Go"
```

Point it at a specific project:

```bash
opencode -c /path/to/project
```

## Command-line flags

| Flag | Short | Description |
| --- | --- | --- |
| `--help` | `-h` | Show help |
| `--version` | `-v` | Print version and exit |
| `--debug` | `-d` | Enable debug logging |
| `--cwd` | `-c` | Working directory to use |
| `--prompt` | `-p` | Run a single prompt in non-interactive mode |
| `--output-format` | `-f` | `text` (default) or `json` |
| `--quiet` | `-q` | Hide the spinner in non-interactive mode |

Non-interactive mode prints the result to stdout and exits. Permissions are auto-approved for that session, which makes it suitable for scripts and CI.

```bash
opencode -p "Summarize the changes in this repo" -f json -q
```

## Configuration

OpenCode reads configuration from the first of these locations, then merges a project-local file on top:

1. `~/.opencode.json`
2. `$XDG_CONFIG_HOME/opencode/.opencode.json`
3. `~/.config/opencode/.opencode.json`
4. `<project>/.opencode.json` (merged over the global config)

Example:

```json
{
  "$schema": "https://raw.githubusercontent.com/opencode-ai/opencode/main/opencode-schema.json",
  "data": {
    "directory": ".opencode"
  },
  "providers": {
    "anthropic": { "apiKey": "sk-ant-..." },
    "openai": { "apiKey": "sk-...", "disabled": false }
  },
  "agents": {
    "coder": { "model": "claude-4-sonnet", "maxTokens": 5000, "reasoningEffort": "high" },
    "summarizer": { "model": "claude-4-sonnet", "maxTokens": 5000 },
    "task": { "model": "claude-4-sonnet", "maxTokens": 5000 },
    "title": { "model": "claude-4-sonnet", "maxTokens": 80 }
  },
  "contextPaths": ["CLAUDE.md", ".cursorrules", "docs/conventions.md"],
  "shell": { "path": "/bin/bash", "args": ["-l"] },
  "mcpServers": {
    "example": {
      "type": "stdio",
      "command": "path/to/mcp-server",
      "env": [],
      "args": []
    }
  },
  "lsp": {
    "go": { "disabled": false, "command": "gopls" }
  },
  "tui": { "theme": "opencode" },
  "autoCompact": true,
  "debug": false,
  "debugLSP": false
}
```

### Agents

The assistant is split into four roles. If a model is not configured, OpenCode picks sensible defaults based on the providers you have credentials for (Copilot → Anthropic → OpenAI → Gemini → Groq → OpenRouter → xAI → Bedrock → Azure → Vertex AI).

| Agent | Responsibility |
| --- | --- |
| `coder` | Main coding assistant |
| `summarizer` | Conversation compaction |
| `task` | Sub-tasks delegated by the `agent` tool |
| `title` | Short session titles |

Each agent accepts:

| Option | Description |
| --- | --- |
| `model` | Model ID, e.g. `claude-3.7-sonnet`, `gpt-4.1`, `gemini-2.5` |
| `maxTokens` | Max output tokens (capped at half the model's context window) |
| `reasoningEffort` | `low`, `medium`, or `high` for reasoning-capable OpenAI/local models |

### Context files

Project context is loaded at startup from the following paths by default:

- `.github/copilot-instructions.md`
- `.cursorrules`
- `.cursor/rules/`
- `CLAUDE.md`, `CLAUDE.local.md`
- `opencode.md`, `opencode.local.md`
- `OpenCode.md`, `OpenCode.local.md`
- `OPENCODE.md`, `OPENCODE.local.md`

Override the list with `contextPaths` in your config.

### Auto compact

When `autoCompact` is enabled (default), OpenCode summarizes the session once token usage reaches about 95% of the model's context window, then continues in a new session with the summary. It can also be triggered manually with the **Compact Session** command.

### Themes

Built-in: `opencode` (default), `catppuccin`, `dracula`, `flexoki`, `gruvbox`, `monokai`, `onedark`, `tokyonight`, `tron`. Switch with `Ctrl+T` or set `tui.theme`.

### JSON Schema

[`opencode-schema.json`](opencode-schema.json) provides validation and autocompletion in editors that support JSON Schema. Regenerate it with:

```bash
go run cmd/schema/main.go > opencode-schema.json
```

## Environment variables

| Variable | Purpose |
| --- | --- |
| `ANTHROPIC_API_KEY` | Anthropic Claude |
| `OPENAI_API_KEY` | OpenAI |
| `GEMINI_API_KEY` | Google Gemini |
| `GITHUB_TOKEN` | GitHub Copilot |
| `GROQ_API_KEY` | Groq |
| `OPENROUTER_API_KEY` | OpenRouter |
| `XAI_API_KEY` | xAI (Grok) |
| `AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / `AWS_REGION` | AWS Bedrock |
| `AZURE_OPENAI_ENDPOINT` / `AZURE_OPENAI_API_KEY` / `AZURE_OPENAI_API_VERSION` | Azure OpenAI |
| `VERTEXAI_PROJECT` / `VERTEXAI_LOCATION` | Google Cloud Vertex AI |
| `LOCAL_ENDPOINT` | OpenAI-compatible self-hosted endpoint |
| `SHELL` | Shell used by the `bash` tool (default `/bin/bash`) |
| `OPENCODE_DEV_DEBUG` | Set to `true` to write `.opencode/debug.log` |

Copilot also discovers OAuth tokens from `~/.config/github-copilot/hosts.json` and `apps.json`.

## Supported providers

| Provider | Models (summary) |
| --- | --- |
| Anthropic | Claude 3 Haiku/Opus, 3.5 Sonnet/Haiku, 3.7 Sonnet, 4 Sonnet, 4 Opus |
| OpenAI | GPT-4.1 family, GPT-4.5 Preview, GPT-4o family, o1/o3/o4 families |
| GitHub Copilot | GPT-4o/4.1, Claude 3.5/3.7/4, o1/o3-mini/o4-mini, Gemini 2.0 Flash / 2.5 Pro |
| Google Gemini | 2.5 Pro, 2.5 Flash, 2.0 Flash, 2.0 Flash Lite |
| Groq | Llama 4 Maverick/Scout, Llama 3.3 70B, Qwen QwQ-32B, DeepSeek R1 distill |
| OpenRouter | OpenAI, Anthropic and Gemini model IDs plus DeepSeek R1 (free) |
| xAI | Grok 3 beta, mini, fast, mini-fast |
| AWS Bedrock | Claude 3.7 Sonnet |
| Azure OpenAI | GPT-4.1 / 4.5 / 4o / o1 / o3 / o4 Mini |
| Vertex AI | Gemini 2.5 / 2.5 Flash |
| Local | Any model served by LM Studio, llama.cpp, Ollama, etc. |

## Keyboard shortcuts

### Global

| Shortcut | Action |
| --- | --- |
| `Ctrl+C` | Quit |
| `Ctrl+?` / `?` | Toggle help |
| `Ctrl+L` | Logs page |
| `Ctrl+K` | Command dialog |
| `Ctrl+O` | Model selection |
| `Ctrl+T` | Switch theme |
| `Ctrl+F` | Attach files/images |
| `Ctrl+S` | Switch session (editor not focused) |
| `Esc` | Close overlay / go back |

### Chat and editor

| Shortcut | Action |
| --- | --- |
| `Ctrl+N` | New session |
| `Ctrl+X` | Cancel generation |
| `i` | Focus editor |
| `Enter` / `Ctrl+S` | Send message |
| `\` + `Enter` | Insert newline |
| `Ctrl+E` | Open external editor |
| `Ctrl+R` + index | Remove an attachment |
| `Ctrl+R` `R` | Remove all attachments |

### Dialogs

| Shortcut | Action |
| --- | --- |
| `↑`/`k`, `↓`/`j` | Move through lists |
| `←`/`h`, `→`/`l` | Previous / next provider (model dialog) |
| `Enter` / `Space` | Confirm |
| `a` / `A` / `d` | Allow / allow for session / deny permission |

## Image attachments

Press `Ctrl+F` to pick an image. Supported formats: `.jpg`, `.jpeg`, `.png`, `.webp`; max 5 MB each and 5 attachments per message. The selected model must support vision input.

## Tools

| Tool | Description |
| --- | --- |
| `glob` | Find files by pattern |
| `grep` | Search file contents |
| `ls` | List directory contents |
| `view` | Read a file |
| `write` | Create or overwrite a file |
| `edit` | Modify a file |
| `patch` | Apply a unified diff |
| `bash` | Run shell commands |
| `fetch` | Fetch and convert a URL to markdown |
| `sourcegraph` | Search public code repositories |
| `diagnostics` | Read LSP diagnostics |
| `agent` | Delegate a read-only sub-task |

Read-only bash commands run without prompting; destructive commands and network tools are gated by the permission dialog. MCP tools follow the same permission model.

## MCP

External tools are configured under `mcpServers` and discovered at startup. Tools are exposed to the model as `<server>_<tool>`.

```json
{
  "mcpServers": {
    "example": {
      "type": "stdio",
      "command": "path/to/mcp-server",
      "env": ["KEY=value"],
      "args": []
    },
    "web-example": {
      "type": "sse",
      "url": "https://example.com/mcp",
      "headers": { "Authorization": "Bearer token" }
    }
  }
}
```

`type` defaults to `stdio`; `sse` is also supported.

## LSP

```json
{
  "lsp": {
    "go": { "command": "gopls" },
    "typescript": { "command": "typescript-language-server", "args": ["--stdio"] }
  }
}
```

OpenCode starts each configured server, watches the workspace for changes, and restarts a server if it exits unexpectedly. The assistant currently consumes LSP through the `diagnostics` tool, though the client also implements completion, hover, definition, references and inline completion.

## Custom commands

Create Markdown files under:

- `$XDG_CONFIG_HOME/opencode/commands/` or `~/.opencode/commands/` → `user:<name>`
- `<project>/.opencode/commands/` → `project:<name>`

Subdirectories become namespaces: `commands/git/commit.md` → `user:git:commit`. Placeholders such as `$ISSUE_NUMBER` prompt you for values when the command runs.

```markdown
RUN gh issue view $ISSUE_NUMBER --json title,body,comments
RUN git grep --author="$AUTHOR_NAME" -n .
```

Open the command palette with `Ctrl+K` to run custom or built-in commands (**Initialize Project**, **Compact Session**).

## Data and storage

Sessions, messages, token counts and file metadata are stored in SQLite at `<data.directory>/opencode.db` (`.opencode/opencode.db` by default). Migrations run automatically on startup.

## Architecture

```
cmd/                  CLI entrypoint (Cobra)
internal/app/         Application services and non-interactive runner
internal/config/      Config loading, defaults and validation
internal/db/          SQLite layer, migrations, sqlc-generated queries
internal/llm/         Providers, models, prompts, agents, tools, MCP
internal/lsp/         LSP client, protocol types, file watcher
internal/tui/         Bubble Tea UI: pages, components, themes, layout
internal/session/     Session service
internal/message/     Message and attachment models
internal/permission/  Tool permission handling
internal/logging/     Structured logging
```

## Development

```bash
go build -o opencode       # build
go test ./...              # tests
sqlc generate              # regenerate DB code after SQL/migration changes
OPENCODE_DEV_DEBUG=true go run .   # debug logs in .opencode/debug.log
```

`CGO_ENABLED=0` is used for releases, so the binaries are static and portable.

## License

MIT — see [LICENSE](LICENSE).
