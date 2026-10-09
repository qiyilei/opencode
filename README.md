# Archived: Project has Moved

This repository is no longer maintained and has been archived for provenance.

The project has continued under the name [Crush][crush], developed by the original author and the Charm team.

Please follow [Crush][crush] for ongoing development.

[crush]: https://github.com/charmbracelet/crush


# ⌬ OpenCode

<p align="center"><img src="https://github.com/user-attachments/assets/9ae61ef6-70e5-4876-bc45-5bcb4e52c714" width="800"></p>

> **⚠️ Early Development Notice:** This project is in early development and is not yet ready for production use. Features may change, break, or be incomplete. Use at your own risk.

A powerful terminal-based AI assistant for developers, providing intelligent coding assistance directly in your terminal.

## Overview

OpenCode is a Go-based CLI application that brings AI assistance to your terminal. It provides a TUI (Terminal User Interface) for interacting with various AI models to help with coding tasks, debugging, and more.

<p>For a quick video overview, check out
<a href="https://www.youtube.com/watch?v=P8luPmEa1QI"><img width="25" src="https://upload.wikimedia.org/wikipedia/commons/0/09/YouTube_full-color_icon_%282017%29.svg"> OpenCode + Gemini 2.5 Pro: BYE Claude Code! I'm SWITCHING To the FASTEST AI Coder!</a></p>

<a href="https://www.youtube.com/watch?v=P8luPmEa1QI"><img width="550" src="https://i3.ytimg.com/vi/P8luPmEa1QI/maxresdefault.jpg"></a><p>

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
  - [Environment Variables](#environment-variables)
  - [Shell Configuration](#shell-configuration)
  - [Configuration File Structure](#configuration-file-structure)
  - [Context Files](#context-files)
  - [Themes](#themes)
  - [JSON Schema](#json-schema)
- [Supported AI Models](#supported-ai-models)
- [Usage](#usage)
- [Non-interactive Prompt Mode](#non-interactive-prompt-mode)
- [Command-line Flags](#command-line-flags)
- [Keyboard Shortcuts](#keyboard-shortcuts)
- [Image Attachments](#image-attachments)
- [AI Assistant Tools](#ai-assistant-tools)
- [Architecture](#architecture)
- [Custom Commands](#custom-commands)
- [MCP (Model Context Protocol)](#mcp-model-context-protocol)
- [LSP (Language Server Protocol)](#lsp-language-server-protocol)
- [Using Github Copilot](#using-github-copilot)
- [Using a self-hosted model provider](#using-a-self-hosted-model-provider)
- [Development](#development)
- [Acknowledgments](#acknowledgments)
- [License](#license)
- [Contributing](#contributing)

## Features

- **Interactive TUI**: Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) for a smooth terminal experience
- **Multiple AI Providers**: Support for OpenAI, Anthropic Claude, Google Gemini, AWS Bedrock, Groq, Azure OpenAI, OpenRouter, and xAI
- **Session Management**: Save and manage multiple conversation sessions
- **Tool Integration**: AI can execute commands, search files, and modify code
- **Vim-like Editor**: Integrated editor with text input capabilities
- **Image Attachments**: Attach images to your messages via the built-in file picker (with image preview)
- **Themes**: 9 built-in color themes (including Catppuccin, Dracula, Gruvbox, Tokyo Night) with light/dark terminal support
- **Persistent Storage**: SQLite database for storing conversations and sessions
- **LSP Integration**: Language Server Protocol support for code intelligence
- **File Change Tracking**: Track and visualize file changes during sessions
- **External Editor Support**: Open your preferred editor for composing messages
- **Named Arguments for Custom Commands**: Create powerful custom commands with multiple named placeholders
- **Auto Compact**: Automatically summarize the conversation when approaching the model's context window limit

## Installation

### Using the Install Script

```bash
# Install the latest version
curl -fsSL https://raw.githubusercontent.com/opencode-ai/opencode/refs/heads/main/install | bash

# Install a specific version
curl -fsSL https://raw.githubusercontent.com/opencode-ai/opencode/refs/heads/main/install | VERSION=0.1.0 bash
```

### Using Homebrew (macOS and Linux)

```bash
brew install opencode-ai/tap/opencode
```

### Using AUR (Arch Linux)

```bash
# Using yay
yay -S opencode-ai-bin

# Using paru
paru -S opencode-ai-bin
```

### Using Go

```bash
go install github.com/opencode-ai/opencode@latest
```

## Configuration

OpenCode looks for configuration in the following locations:

- `$HOME/.opencode.json`
- `$XDG_CONFIG_HOME/opencode/.opencode.json`
- `./.opencode.json` (local directory)

### Auto Compact Feature

OpenCode includes an auto compact feature that automatically summarizes your conversation when it approaches the model's context window limit. When enabled (default setting), this feature:

- Monitors token usage during your conversation
- Automatically triggers summarization when usage reaches 95% of the model's context window
- Creates a new session with the summary, allowing you to continue your work without losing context
- Helps prevent "out of context" errors that can occur with long conversations

You can enable or disable this feature in your configuration file:

```json
{
  "autoCompact": true // default is true
}
```

### Environment Variables

You can configure OpenCode using environment variables:

| Environment Variable       | Purpose                                                                          |
| -------------------------- | -------------------------------------------------------------------------------- |
| `ANTHROPIC_API_KEY`        | For Claude models                                                                |
| `OPENAI_API_KEY`           | For OpenAI models                                                                |
| `GEMINI_API_KEY`           | For Google Gemini models                                                         |
| `GITHUB_TOKEN`             | For Github Copilot models (see [Using Github Copilot](#using-github-copilot))    |
| `VERTEXAI_PROJECT`         | For Google Cloud VertexAI (Gemini)                                               |
| `VERTEXAI_LOCATION`        | For Google Cloud VertexAI (Gemini)                                               |
| `GROQ_API_KEY`             | For Groq models                                                                  |
| `OPENROUTER_API_KEY`       | For OpenRouter models                                                            |
| `XAI_API_KEY`              | For xAI (Grok) models                                                            |
| `AWS_ACCESS_KEY_ID`        | For AWS Bedrock (Claude)                                                         |
| `AWS_SECRET_ACCESS_KEY`    | For AWS Bedrock (Claude)                                                         |
| `AWS_REGION`               | For AWS Bedrock (Claude)                                                         |
| `AZURE_OPENAI_ENDPOINT`    | For Azure OpenAI models                                                          |
| `AZURE_OPENAI_API_KEY`     | For Azure OpenAI models (optional when using Entra ID)                           |
| `AZURE_OPENAI_API_VERSION` | For Azure OpenAI models                                                          |
| `LOCAL_ENDPOINT`           | For self-hosted models (see [Using a self-hosted model provider](#using-a-self-hosted-model-provider)) |
| `SHELL`                    | Default shell to use (if not specified in config)                                |

### Shell Configuration

OpenCode allows you to configure the shell used by the bash tool. By default, it uses the shell specified in the `SHELL` environment variable, or falls back to `/bin/bash` if not set.

You can override this in your configuration file:

```json
{
  "shell": {
    "path": "/bin/zsh",
    "args": ["-l"]
  }
}
```

This is useful if you want to use a different shell than your default system shell, or if you need to pass specific arguments to the shell.

### Configuration File Structure

```json
{
  "data": {
    "directory": ".opencode"
  },
  "providers": {
    "openai": {
      "apiKey": "your-api-key",
      "disabled": false
    },
    "anthropic": {
      "apiKey": "your-api-key",
      "disabled": false
    },
    "copilot": {
      "disabled": false
    },
    "groq": {
      "apiKey": "your-api-key",
      "disabled": false
    },
    "openrouter": {
      "apiKey": "your-api-key",
      "disabled": false
    }
  },
  "agents": {
    "coder": {
      "model": "claude-3.7-sonnet",
      "maxTokens": 5000,
      "reasoningEffort": "high"
    },
    "task": {
      "model": "claude-3.7-sonnet",
      "maxTokens": 5000
    },
    "title": {
      "model": "claude-3.7-sonnet",
      "maxTokens": 80
    }
  },
  "contextPaths": [
    "opencode.md",
    ".cursorrules"
  ],
  "shell": {
    "path": "/bin/bash",
    "args": ["-l"]
  },
  "mcpServers": {
    "example": {
      "type": "stdio",
      "command": "path/to/mcp-server",
      "env": [],
      "args": []
    }
  },
  "lsp": {
    "go": {
      "disabled": false,
      "command": "gopls"
    }
  },
  "tui": {
    "theme": "opencode"
  },
  "debug": false,
  "debugLSP": false,
  "autoCompact": true
}
```

#### Agent Configuration Options

Each agent (`coder`, `summarizer`, `task`, `title`) supports the following options:

| Option            | Description                                                                                             |
| ----------------- | ------------------------------------------------------------------------------------------------------- |
| `model`           | Model ID (e.g. `claude-3.7-sonnet`, `gpt-4.1`, `gemini-2.5`)                                             |
| `maxTokens`       | Maximum number of tokens the model can generate (capped at half the model's context window)               |
| `reasoningEffort` | Reasoning effort for models that support it: `low`, `medium`, or `high` (OpenAI and local models)        |

### Context Files

OpenCode automatically loads project context from files such as `opencode.md`, `CLAUDE.md`, `.cursorrules`, and `.github/copilot-instructions.md` at startup. These files' contents are included as system context for the AI assistant.

By default, the following paths are checked:

- `.github/copilot-instructions.md`
- `.cursorrules`
- `.cursor/rules/`
- `CLAUDE.md`
- `CLAUDE.local.md`
- `opencode.md`
- `opencode.local.md`
- `OpenCode.md`
- `OpenCode.local.md`
- `OPENCODE.md`
- `OPENCODE.local.md`

You can override this list with the `contextPaths` option in your configuration file to point at any custom context files:

```json
{
  "contextPaths": ["docs/conventions.md", ".cursorrules"]
}
```

### Themes

OpenCode includes 9 built-in color themes with adaptive colors that work on both light and dark terminal backgrounds:

| Theme         | Description                          |
| ------------ | ----------------------------------- |
| `opencode`   | The default OpenCode theme          |
| `catppuccin` | Soothing pastel theme               |
| `dracula`    | Dark purple theme                   |
| `flexoki`    | Ink-on-paper theme                  |
| `gruvbox`    | Retro groove color scheme           |
| `monokai`    | Monokai Pro inspired theme          |
| `onedark`    | Atom's One Dark theme               |
| `tokyonight` | Tokyo Night theme                   |
| `tron`       | Tron-inspired theme                 |

You can change themes at runtime with the `Ctrl+T` shortcut, or set one permanently in your configuration file:

```json
{
  "tui": {
    "theme": "catppuccin"
  }
}
```

### JSON Schema

A JSON Schema for the configuration file is available in [`opencode-schema.json`](opencode-schema.json). Reference it from your config file to get validation and autocompletion in editors that support JSON Schema:

```json
{
  "$schema": "https://raw.githubusercontent.com/opencode-ai/opencode/main/opencode-schema.json"
}
```

To regenerate the schema after changing the configuration code, run:

```bash
go run cmd/schema/main.go > opencode-schema.json
```

## Supported AI Models

OpenCode supports a variety of AI models from different providers:

### OpenAI

- GPT-4.1 family (gpt-4.1, gpt-4.1-mini, gpt-4.1-nano)
- GPT-4.5 Preview
- GPT-4o family (gpt-4o, gpt-4o-mini)
- O1 family (o1, o1-pro, o1-mini)
- O3 family (o3, o3-mini)
- O4 Mini

### Anthropic

- Claude 4 Sonnet
- Claude 4 Opus
- Claude 3.5 Sonnet
- Claude 3.5 Haiku
- Claude 3.7 Sonnet
- Claude 3 Haiku
- Claude 3 Opus

### GitHub Copilot

- GPT-3.5 Turbo
- GPT-4
- GPT-4o
- GPT-4o Mini
- GPT-4.1
- Claude 3.5 Sonnet
- Claude 3.7 Sonnet
- Claude 3.7 Sonnet Thinking
- Claude Sonnet 4
- O1
- O3 Mini
- O4 Mini
- Gemini 2.0 Flash
- Gemini 2.5 Pro

### Google

- Gemini 2.5
- Gemini 2.5 Flash
- Gemini 2.0 Flash
- Gemini 2.0 Flash Lite

### AWS Bedrock

- Claude 3.7 Sonnet

### Groq

- Llama 4 Maverick (17b-128e-instruct)
- Llama 4 Scout (17b-16e-instruct)
- QWEN QWQ-32b
- Deepseek R1 distill Llama 70b
- Llama 3.3 70b Versatile

### OpenRouter

- GPT-4.1 family (gpt-4.1, gpt-4.1-mini, gpt-4.1-nano)
- GPT-4.5 Preview
- GPT-4o family (gpt-4o, gpt-4o-mini)
- O1 family (o1, o1-pro, o1-mini)
- O3 family (o3, o3-mini)
- O4 Mini
- Claude 3.5 Sonnet / Claude 3.5 Haiku
- Claude 3.7 Sonnet
- Claude 3 Opus / Claude 3 Haiku
- Gemini 2.5 / Gemini 2.5 Flash
- DeepSeek R1 (free)

### xAI

- Grok 3 (beta)
- Grok 3 mini (beta)
- Grok 3 fast (beta)
- Grok 3 mini fast (beta)

### Azure OpenAI

- GPT-4.1 family (gpt-4.1, gpt-4.1-mini, gpt-4.1-nano)
- GPT-4.5 Preview
- GPT-4o family (gpt-4o, gpt-4o-mini)
- O1 family (o1, o1-mini)
- O3 family (o3, o3-mini)
- O4 Mini

### Google Cloud VertexAI

- Gemini 2.5
- Gemini 2.5 Flash

### Local (self-hosted)

- Any model served by an OpenAI-compatible local endpoint (LM Studio, llama.cpp, Ollama, etc.). See [Using a self-hosted model provider](#using-a-self-hosted-model-provider).

## Usage

```bash
# Start OpenCode
opencode

# Start with debug logging
opencode -d

# Start with a specific working directory
opencode -c /path/to/project

# Print version
opencode -v
```

## Non-interactive Prompt Mode

You can run OpenCode in non-interactive mode by passing a prompt directly as a command-line argument. This is useful for scripting, automation, or when you want a quick answer without launching the full TUI.

```bash
# Run a single prompt and print the AI's response to the terminal
opencode -p "Explain the use of context in Go"

# Get response in JSON format
opencode -p "Explain the use of context in Go" -f json

# Run without showing the spinner (useful for scripts)
opencode -p "Explain the use of context in Go" -q
```

In this mode, OpenCode will process your prompt, print the result to standard output, and then exit. All permissions are auto-approved for the session.

By default, a spinner animation is displayed while the model is processing your query. You can disable this spinner with the `-q` or `--quiet` flag, which is particularly useful when running OpenCode from scripts or automated workflows.

### Output Formats

OpenCode supports the following output formats in non-interactive mode:

| Format | Description                     |
| ------ | ------------------------------- |
| `text` | Plain text output (default)     |
| `json` | Output wrapped in a JSON object |

The output format is implemented as a strongly-typed `OutputFormat` in the codebase, ensuring type safety and validation when processing outputs.

## Command-line Flags

| Flag              | Short | Description                                         |
| ----------------- | ----- | --------------------------------------------------- |
| `--help`          | `-h`  | Display help information                            |
| `--version`       | `-v`  | Print version and exit                              |
| `--debug`         | `-d`  | Enable debug mode                                   |
| `--cwd`           | `-c`  | Set current working directory                       |
| `--prompt`        | `-p`  | Run a single prompt in non-interactive mode         |
| `--output-format` | `-f`  | Output format for non-interactive mode (text, json) |
| `--quiet`         | `-q`  | Hide spinner in non-interactive mode                |

## Keyboard Shortcuts

### Global Shortcuts

| Shortcut | Action                                                  |
| -------- | ------------------------------------------------------- |
| `Ctrl+C` | Quit application                                        |
| `Ctrl+?` | Toggle help dialog                                      |
| `?`      | Toggle help dialog (when not in editing mode)           |
| `Ctrl+L` | View logs                                               |
| `Ctrl+S` | Switch session (when editor is not focused)              |
| `Ctrl+K` | Command dialog                                          |
| `Ctrl+O` | Toggle model selection dialog                           |
| `Ctrl+T` | Switch theme                                            |
| `Ctrl+F` | Select files/images to attach to your message           |
| `Esc`    | Close current overlay/dialog or return to previous mode |

### Chat Page Shortcuts

| Shortcut | Action                                  |
| -------- | --------------------------------------- |
| `Ctrl+N` | Create new session                      |
| `Ctrl+X` | Cancel current operation/generation     |
| `i`      | Focus editor (when not in writing mode) |
| `Esc`    | Exit writing mode and focus messages    |

### Editor Shortcuts

| Shortcut            | Action                                    |
| ------------------- | ----------------------------------------- |
| `Enter` or `Ctrl+S` | Send message (when editor is focused)     |
| `\` + `Enter`       | Insert a newline instead of sending       |
| `Ctrl+E`            | Open external editor                      |
| `Ctrl+R` + index    | Delete the attachment at the given index  |
| `Ctrl+R` `R`       | Delete all attachments                    |
| `Esc`               | Blur editor and focus messages            |

> **Tip:** If the last character of your message is a backslash, pressing `Enter` inserts a newline instead of sending — useful for composing multi-line messages.

### Session Dialog Shortcuts

| Shortcut   | Action           |
| ---------- | ---------------- |
| `↑` or `k` | Previous session |
| `↓` or `j` | Next session     |
| `Enter`    | Select session   |
| `Esc`      | Close dialog     |

### Model Dialog Shortcuts

| Shortcut   | Action            |
| ---------- | ----------------- |
| `↑` or `k` | Move up           |
| `↓` or `j` | Move down         |
| `←` or `h` | Previous provider |
| `→` or `l` | Next provider     |
| `Esc`      | Close dialog      |

### Permission Dialog Shortcuts

| Shortcut                | Action                       |
| ----------------------- | ---------------------------- |
| `←` or `left`           | Switch options left          |
| `→` or `right` or `tab` | Switch options right         |
| `Enter` or `space`      | Confirm selection            |
| `a`                     | Allow permission             |
| `A`                     | Allow permission for session |
| `d`                     | Deny permission              |

### Logs Page Shortcuts

| Shortcut           | Action              |
| ------------------ | ------------------- |
| `Backspace` or `q` | Return to chat page |

## Image Attachments

OpenCode supports attaching images to your messages, which is useful for sharing screenshots, error messages, or designs with vision-capable models.

- Press `Ctrl+F` to open the file picker and select an image
- Supported formats: `.jpg`, `.jpeg`, `.png`, and `.webp`
- Maximum file size: 5 MB
- Maximum of 5 attachments per message
- Images are previewed directly in the file picker

The selected model must support attachments (most vision-capable models do). Use `Ctrl+R` followed by an index to remove an attachment, or `Ctrl+R` `R` to remove them all.

## AI Assistant Tools

OpenCode's AI assistant has access to various tools to help with coding tasks:

### File and Code Tools

| Tool          | Description                 | Parameters                                                                               |
| ------------- | --------------------------- | ---------------------------------------------------------------------------------------- |
| `glob`        | Find files by pattern       | `pattern` (required), `path` (optional)                                                  |
| `grep`        | Search file contents        | `pattern` (required), `path` (optional), `include` (optional), `literal_text` (optional) |
| `ls`          | List directory contents     | `path` (optional), `ignore` (optional array of patterns)                                 |
| `view`        | View file contents          | `file_path` (required), `offset` (optional), `limit` (optional)                          |
| `write`       | Write to files              | `file_path` (required), `content` (required)                                             |
| `edit`        | Edit files                  | Various parameters for file editing                                                      |
| `patch`       | Apply patches to files      | `file_path` (required), `diff` (required)                                                |
| `diagnostics` | Get diagnostics information | `file_path` (optional)                                                                   |

### Other Tools

| Tool          | Description                            | Parameters                                                                                |
| ------------- | -------------------------------------- | ----------------------------------------------------------------------------------------- |
| `bash`        | Execute shell commands                 | `command` (required), `timeout` (optional)                                                |
| `fetch`       | Fetch data from URLs                   | `url` (required), `format` (required), `timeout` (optional)                               |
| `sourcegraph` | Search code across public repositories | `query` (required), `count` (optional), `context_window` (optional), `timeout` (optional) |
| `agent`       | Run sub-tasks with the AI agent        | `prompt` (required)                                                                       |

## Architecture

OpenCode is built with a modular architecture:

- **cmd**: Command-line interface using Cobra
- **internal/app**: Core application services
- **internal/config**: Configuration management
- **internal/db**: Database operations and migrations
- **internal/llm**: LLM providers and tools integration
- **internal/tui**: Terminal UI components and layouts
- **internal/logging**: Logging infrastructure
- **internal/message**: Message handling
- **internal/session**: Session management
- **internal/lsp**: Language Server Protocol integration

## Custom Commands

OpenCode supports custom commands that can be created by users to quickly send predefined prompts to the AI assistant.

### Creating Custom Commands

Custom commands are predefined prompts stored as Markdown files in one of three locations:

1. **User Commands** (prefixed with `user:`):

   ```
   $XDG_CONFIG_HOME/opencode/commands/
   ```

   (typically `~/.config/opencode/commands/` on Linux/macOS)

   or

   ```
   $HOME/.opencode/commands/
   ```

2. **Project Commands** (prefixed with `project:`):

   ```
   <PROJECT DIR>/.opencode/commands/
   ```

Each `.md` file in these directories becomes a custom command. The file name (without extension) becomes the command ID.

For example, creating a file at `~/.config/opencode/commands/prime-context.md` with content:

```markdown
RUN git ls-files
READ README.md
```

This creates a command called `user:prime-context`.

### Command Arguments

OpenCode supports named arguments in custom commands using placeholders in the format `$NAME` (where NAME consists of uppercase letters, numbers, and underscores, and must start with a letter).

For example:

```markdown
# Fetch Context for Issue $ISSUE_NUMBER

RUN gh issue view $ISSUE_NUMBER --json title,body,comments
RUN git grep --author="$AUTHOR_NAME" -n .
RUN grep -R "$SEARCH_PATTERN" $DIRECTORY
```

When you run a command with arguments, OpenCode will prompt you to enter values for each unique placeholder. Named arguments provide several benefits:

- Clear identification of what each argument represents
- Ability to use the same argument multiple times
- Better organization for commands with multiple inputs

### Organizing Commands

You can organize commands in subdirectories:

```
~/.config/opencode/commands/git/commit.md
```

This creates a command with ID `user:git:commit`.

### Using Custom Commands

1. Press `Ctrl+K` to open the command dialog
2. Select your custom command (prefixed with either `user:` or `project:`)
3. Press Enter to execute the command

The content of the command file will be sent as a message to the AI assistant.

### Built-in Commands

OpenCode includes several built-in commands:

| Command            | Description                                                                                         |
| ------------------ | --------------------------------------------------------------------------------------------------- |
| Initialize Project | Creates or updates the OpenCode.md memory file with project-specific information                    |
| Compact Session    | Manually triggers the summarization of the current session, creating a new session with the summary |

## MCP (Model Context Protocol)

OpenCode implements the Model Context Protocol (MCP) to extend its capabilities through external tools. MCP provides a standardized way for the AI assistant to interact with external services and tools.

### MCP Features

- **External Tool Integration**: Connect to external tools and services via a standardized protocol
- **Tool Discovery**: Automatically discover available tools from MCP servers
- **Multiple Connection Types**:
  - **Stdio**: Communicate with tools via standard input/output
  - **SSE**: Communicate with tools via Server-Sent Events
- **Security**: Permission system for controlling access to MCP tools

### Configuring MCP Servers

MCP servers are defined in the configuration file under the `mcpServers` section:

```json
{
  "mcpServers": {
    "example": {
      "type": "stdio",
      "command": "path/to/mcp-server",
      "env": [],
      "args": []
    },
    "web-example": {
      "type": "sse",
      "url": "https://example.com/mcp",
      "headers": {
        "Authorization": "Bearer token"
      }
    }
  }
}
```

### MCP Tool Usage

Once configured, MCP tools are automatically available to the AI assistant alongside built-in tools. They follow the same permission model as other tools, requiring user approval before execution.

## LSP (Language Server Protocol)

OpenCode integrates with Language Server Protocol to provide code intelligence features across multiple programming languages.

### LSP Features

- **Multi-language Support**: Connect to language servers for different programming languages
- **Diagnostics**: Receive error checking and linting information
- **File Watching**: Automatically notify language servers of file changes

### Configuring LSP

Language servers are configured in the configuration file under the `lsp` section:

```json
{
  "lsp": {
    "go": {
      "disabled": false,
      "command": "gopls"
    },
    "typescript": {
      "disabled": false,
      "command": "typescript-language-server",
      "args": ["--stdio"]
    }
  }
}
```

### LSP Integration with AI

The AI assistant can access LSP features through the `diagnostics` tool, allowing it to:

- Check for errors in your code
- Suggest fixes based on diagnostics

While the LSP client implementation supports the full LSP protocol (including completions, hover, definition, etc.), currently only diagnostics are exposed to the AI assistant.

## Using Github Copilot

_Copilot support is currently experimental._

### Requirements
- [Copilot chat in the IDE](https://github.com/settings/copilot) enabled in GitHub settings
- One of:
  - VSCode Github Copilot chat extension
  - Github `gh` CLI
  - Neovim Github Copilot plugin (`copilot.vim` or `copilot.lua`)
  - Github token with copilot permissions

If using one of the above plugins or cli tools, make sure you use the authenticate
the tool with your github account. This should create a github token at one of the following locations:
- ~/.config/github-copilot/[hosts,apps].json
- $XDG_CONFIG_HOME/github-copilot/[hosts,apps].json

If using an explicit github token, you may either set the $GITHUB_TOKEN environment variable or add it to the opencode.json config file at `providers.copilot.apiKey`.

## Using a self-hosted model provider

OpenCode can also load and use models from a self-hosted (OpenAI-like) provider.
This is useful for developers who want to experiment with custom models.

### Configuring a self-hosted provider

You can use a self-hosted model by setting the `LOCAL_ENDPOINT` environment variable.
This will cause OpenCode to load and use the models from the specified endpoint.

```bash
LOCAL_ENDPOINT=http://localhost:1235/v1
```

### Configuring a self-hosted model

You can also configure a self-hosted model in the configuration file under the `agents` section:

```json
{
  "agents": {
    "coder": {
      "model": "local.granite-3.3-2b-instruct@q8_0",
      "reasoningEffort": "high"
    }
  }
}
```

## Development

### Prerequisites

- Go 1.24.0 or higher
- [sqlc](https://sqlc.dev/) (only needed when changing database queries or migrations)

### Building from Source

```bash
# Clone the repository
git clone https://github.com/opencode-ai/opencode.git
cd opencode

# Build
go build -o opencode

# Run
./opencode
```

### Running Tests

```bash
go test ./...
```

### Database Code Generation

The SQLite data layer is generated with [sqlc](https://sqlc.dev/). If you change the SQL queries in `internal/db/sql/` or the migrations in `internal/db/migrations/`, regenerate the Go code:

```bash
sqlc generate
```

The configuration lives in [`sqlc.yaml`](sqlc.yaml).

### Debug Logging

For development, set `OPENCODE_DEV_DEBUG=true` to write logs to `.opencode/debug.log` and persist message debugging information:

```bash
OPENCODE_DEV_DEBUG=true go run .
```

## Acknowledgments

OpenCode gratefully acknowledges the contributions and support from these key individuals:

- [@isaacphi](https://github.com/isaacphi) - For the [mcp-language-server](https://github.com/isaacphi/mcp-language-server) project which provided the foundation for our LSP client implementation
- [@adamdottv](https://github.com/adamdottv) - For the design direction and UI/UX architecture

Special thanks to the broader open source community whose tools and libraries have made this project possible.

## License

OpenCode is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Here's how you can contribute:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please make sure to update tests as appropriate and follow the existing code style.
