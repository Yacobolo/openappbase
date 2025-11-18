### Opencode Config API Examples

Source: https://opencode.ai/docs/sdk

Examples for using the opencode Config API to fetch configuration details. This includes getting general config information and listing available providers with their default models.

```typescript
const config = await client.config.get()


const { providers, default: defaults } = await client.config.providers()
```

--------------------------------

### Install Opencode SDK

Source: https://opencode.ai/docs/sdk

Install the opencode SDK using npm. This is the first step to integrating opencode into your project.

```bash
npm install @opencode-ai/sdk
```

--------------------------------

### OpenCode Configuration - Basic LSP Setup

Source: https://opencode.ai/docs/lsp

This is the basic structure for configuring LSP servers in OpenCode. The 'lsp' object is where all LSP server configurations will reside. It's a starting point for customizing LSP behavior.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "lsp": {}
}
```

--------------------------------

### Opencode Project API Examples

Source: https://opencode.ai/docs/sdk

Examples for interacting with the opencode Project API. This covers listing all available projects and retrieving the currently active project.

```typescript
// List all projects
const projects = await client.project.list()


// Get current project
const currentProject = await client.project.current()
```

--------------------------------

### Opencode Path API Example

Source: https://opencode.ai/docs/sdk

An example demonstrating how to use the opencode Path API to retrieve information about the current path.

```typescript
// Get current path information
const pathInfo = await client.path.get()
```

--------------------------------

### Start OpenCode TUI - Shell Command

Source: https://opencode.ai/docs/tui

Run the opencode command to launch the Terminal User Interface for the current directory or a specific path. This initializes the interactive session with an LLM for project work. No dependencies required beyond OpenCode installation.

```shell
opencode
opencode /path/to/project
```

--------------------------------

### Opencode App API Examples

Source: https://opencode.ai/docs/sdk

Examples demonstrating the usage of the opencode App API. This includes writing a log entry with service, level, and message details, and listing available agents.

```typescript
// Write a log entry
await client.app.log({
  body: {
    service: "my-app",
    level: "info",
    message: "Operation completed",
  },
})


// List available agents
const agents = await client.app.agents()
```

--------------------------------

### Install clipboard utilities for Linux (X11)

Source: https://opencode.ai/docs/troubleshooting

Commands to install clipboard utilities required for copy/paste functionality on X11-based Linux systems.

```Command Line
apt install -y xclip
# or
apt install -y xsel
```

--------------------------------

### Install clipboard utilities for Linux (Wayland)

Source: https://opencode.ai/docs/troubleshooting

Command to install clipboard utilities required for copy/paste functionality on Wayland-based Linux systems.

```Command Line
apt install -y wl-clipboard
```

--------------------------------

### Set up headless environment for Linux

Source: https://opencode.ai/docs/troubleshooting

Commands to set up a virtual framebuffer for clipboard functionality in headless Linux environments.

```Command Line
apt install -y xvfb
# and run:
Xvfb :99 -screen 0 1024x768x24 > /dev/null 2>&1 &
export DISPLAY=:99.0
```

--------------------------------

### GitLab CI/CD YAML Configuration for Opencode Setup

Source: https://opencode.ai/docs/gitlab

This YAML defines a GitLab CI/CD job for installing and configuring opencode, authenticating with AI providers and GitLab, and running tasks based on inputs. It requires environment variables like ANTHROPIC_API_KEY and GITLAB_TOKEN_OPENCODE. The job installs dependencies, sets up authentication, configures git, and executes opencode with context from issues or merge requests. Limitations include dependency on GitLab runners and specific API keys; outputs include potential git commits and pushes if changes are made.

```yaml
image: node:22-slim
commands:
  - echo "Installing opencode"
  - npm install --global opencode-ai
  - echo "Installing glab"
  - export GITLAB_TOKEN=$GITLAB_TOKEN_OPENCODE
  - apt-get update --quiet && apt-get install --yes curl wget gpg git && rm --recursive --force /var/lib/apt/lists/*
  - curl --silent --show-error --location "https://raw.githubusercontent.com/upciti/wakemeops/main/assets/install_repository" | bash
  - apt-get install --yes glab
  - echo "Configuring glab"
  - echo $GITLAB_HOST
  - echo "Creating opencode auth configuration"
  - mkdir --parents ~/.local/share/opencode
  - |
    cat > ~/.local/share/opencode/auth.json << EOF
    {
      "anthropic": {
        "type": "api",
        "key": "$ANTHROPIC_API_KEY"
      }
    }
    EOF
  - echo "Configuring git"
  - git config --global user.email "opencode@gitlab.com"
  - git config --global user.name "Opencode"
  - echo "Testing glab"
  - glab issue list
  - echo "Running Opencode"
  - |
    opencode run "
    You are an AI assistant helping with GitLab operations.


    Context: $AI_FLOW_CONTEXT
    Task: $AI_FLOW_INPUT
    Event: $AI_FLOW_EVENT


    Please execute the requested task using the available GitLab tools.
    Be thorough in your analysis and provide clear explanations.


    <important>
    Please use the glab CLI to access data from GitLab. The glab CLI has already been authenticated. You can run the corresponding commands.


    If you are asked to summarize an MR or issue or asked to provide more information then please post back a note to the MR/Issue so that the user can see it.
    You don't need to commit or push up changes, those will be done automatically based on the file changes you make.
    </important>
    "
  - git checkout --branch $CI_WORKLOAD_REF origin/$CI_WORKLOAD_REF
  - echo "Checking for git changes and pushing if any exist"
  - |
    if ! git diff --quiet || ! git diff --cached --quiet || [ --not --zero "$(git ls-files --others --exclude-standard)" ]; then
      echo "Git changes detected, adding and pushing..."
      git add .
      if git diff --cached --quiet; then
        echo "No staged changes to commit"
      else
        echo "Committing changes to branch: $CI_WORKLOAD_REF"
        git commit --message "Codex changes"
        echo "Pushing changes up to $CI_WORKLOAD_REF"
        git push https://gitlab-ci-token:$GITLAB_TOKEN@$GITLAB_HOST/gl-demo-ultimate-dev-ai-epic-17570/test-java-project.git $CI_WORKLOAD_REF
        echo "Changes successfully pushed"
      fi
    else
      echo "No git changes detected, skipping push"
    fi
variables:
  - ANTHROPIC_API_KEY
  - GITLAB_TOKEN_OPENCODE
  - GITLAB_HOST

```

--------------------------------

### Basic Formatter Configuration

Source: https://opencode.ai/docs/formatters

Empty formatter configuration structure for OpenCode. This sets up the basic schema reference and formatter section without enabling any specific formatters. Use this as a starting point for formatter configuration.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "formatter": {}
}
```

--------------------------------

### Create Opencode Client Instance

Source: https://opencode.ai/docs/sdk

Instantiate the opencode client. This method starts both a server and a client, allowing immediate interaction. Options can be provided to customize server URL, fetch implementation, response parsing, and error handling.

```typescript
import { createOpencode } from "@opencode-ai/sdk"


const { client } = await createOpencode()
```

--------------------------------

### opencode Explain Issue Example

Source: https://opencode.ai/docs/github

An example comment for requesting opencode to explain an existing GitHub issue. It triggers opencode to analyze the entire issue thread and provide a clear explanation.

```text
/opencode explain this issue
```

--------------------------------

### Load Specific Model via Command Line in JSON

Source: https://opencode.ai/docs/models

This JSON config example shows how to specify a model for loading priority, often used with command line flags. It uses the provider/model format and $schema for validation. It requires the model to be configured and prioritizes it over defaults when starting OpenCode.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "model": "anthropic/claude-sonnet-4-20250514"
}
```

--------------------------------

### Use MCP tool in prompt

Source: https://opencode.ai/docs/mcp-servers

Shows how to reference and use an MCP tool by name in a prompt. This example demonstrates calling the mcp_everything tool to perform a simple addition.

```text
use the mcp_everything tool to add the number 3 and 4
```

--------------------------------

### opencode GitHub Install Command

Source: https://opencode.ai/docs/github

Command to install opencode in a GitHub repository

```text
opencode github install
```

--------------------------------

### Set Authentication Credentials (JavaScript)

Source: https://opencode.ai/docs/sdk

Demonstrates setting authentication credentials using `client.auth.set`. This example uses Anthropic API keys.

```javascript
await client.auth.set({
  path: { id: "anthropic" },
  body: { type: "api", key: "your-api-key" },
});
```

--------------------------------

### Custom Provider Setup in opencode.json

Source: https://opencode.ai/docs/providers

This JSON configuration defines a custom OpenAI-compatible provider for OpenCode projects, specifying the package, display name, API endpoint, and available models. It requires matching the provider ID from auth login and uses environment variables for keys. Limitations include ensuring the correct npm package and baseURL for compatibility; standard providers handle limits automatically.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "myprovider": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "My AI ProviderDisplay Name",
      "options": {
        "baseURL": "https://api.myprovider.com/v1"
      },
      "models": {
        "my-model-name": {
          "name": "My Model Display Name"
        }
      }
    }
  }
}
```

```json
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "myprovider": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "My AI ProviderDisplay Name",
      "options": {
        "baseURL": "https://api.myprovider.com/v1",
        "apiKey": "{env:ANTHROPIC_API_KEY}",
        "headers": {
          "Authorization": "Bearer custom-token"
        }
      },
      "models": {
        "my-model-name": {
          "name": "My Model Display Name",
          "limit": {
            "context": 200000,
            "output": 65536
          }
        }
      }
    }
  }
}
```

--------------------------------

### Create basic OpenCode plugin structure

Source: https://opencode.ai/docs/plugins

Basic JavaScript plugin template that shows the required export structure and context object parameters. This template serves as the foundation for all OpenCode plugins and demonstrates the async function signature with context destructuring.

```javascript
export const MyPlugin = async ({ project, client, $, directory, worktree }) => {
  console.log("Plugin initialized!")

  return {
    // Hook implementations go here
  }
}
```

--------------------------------

### OpenCode Configuration - Adding a Custom LSP Server

Source: https://opencode.ai/docs/lsp

This example demonstrates how to add a custom LSP server to OpenCode. You specify the command to run the server and the file extensions it should handle. This allows integration with language servers not included by default.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "lsp": {
    "custom-lsp": {
      "command": ["custom-lsp-server", "--stdio"],
      "extensions": [".custom"]
    }
  }
}
```

--------------------------------

### Implement TypeScript plugin with type safety

Source: https://opencode.ai/docs/plugins

TypeScript plugin example using the Plugin type from @opencode-ai/plugin package for compile-time type checking. This ensures type safety for context parameters and return values, providing better development experience and error prevention.

```typescript
import type { Plugin } from "@opencode-ai/plugin"

export const MyPlugin: Plugin = async ({ project, client, $, directory, worktree }) => {
  return {
    // Type-safe hook implementations
  }
}
```

--------------------------------

### Authenticate with Private NPM Registry (Shell)

Source: https://opencode.ai/docs/enterprise

This shell command logs into a private NPM registry like JFrog Artifactory to enable package installation. It requires NPM installed and the registry URL; it outputs authentication confirmation and creates ~/.npmrc. Ensure developers authenticate before running OpenCode to avoid installation failures.

```bash
npm login --registry=https://your-company.jfrog.io/api/npm/npm-virtual/
```

--------------------------------

### Configure OpenCode Providers in JSON

Source: https://opencode.ai/docs/providers

These JSON snippets define custom provider configurations in the OpenCode config.json file for integrating local and remote AI models. They utilize the ai-sdk compatible libraries for OpenAI-like APIs. Inputs include provider IDs, NPM packages, base URLs, and model details; outputs enable accessing specific models via the OpenCode interface. Limitations: Requires valid API keys, local server setup for Ollama, and compatibility with the ai-sdk package versions.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "ollama": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "Ollama (local)",
      "options": {
        "baseURL": "http://localhost:11434/v1"
      },
      "models": {
        "llama2": {
          "name": "Llama 2"
        }
      }
    }
  }
}
```

```json
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "openrouter": {
      "models": {
        "somecoolnewmodel": {}
      }
    }
  }
}
```

```json
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "openrouter": {
      "models": {
        "moonshotai/kimi-k2": {
          "options": {
            "provider": {
              "order": ["baseten"],
              "allow_fallbacks": false
            }
          }
        }
      }
    }
  }
}
```

--------------------------------

### Documentation agent configuration in YAML

Source: https://opencode.ai/docs/agents

Example configuration for a documentation agent that writes and maintains project documentation. Uses subagent mode and disables bash tools. Includes a system prompt focused on technical writing with clear explanations and code examples.

```yaml
---
description: Writes and maintains project documentation
mode: subagent
tools:
  bash: false
---


You are a technical writer. Create clear, comprehensive documentation.


Focus on:


- Clear explanations
- Proper structure
- Code examples
- User-friendly language
```

--------------------------------

### Set provider credentials via environment variables

Source: https://opencode.ai/docs/providers

Many providers require credentials to be supplied as environment variables. You can pass them inline when invoking `opencode` or export them in your shell profile for persistent use.

```Bash
AWS_ACCESS_KEY_ID=XXX opencode
```

```Bash
export AWS_ACCESS_KEY_ID=XXX
```

```Bash
AZURE_RESOURCE_NAME=XXX opencode
```

```Bash
export AZURE_RESOURCE_NAME=XXX
```

--------------------------------

### Default Keybinds Configuration in JSON

Source: https://opencode.ai/docs/keybinds

Shows the complete default keybind configuration for OpenCode in JSON format. Includes leader key setup and all application shortcuts. Requires OpenCode config schema validation.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "keybinds": {
    "leader": "ctrl+x",
    "app_exit": "ctrl+c,ctrl+d,<leader>q",
    "editor_open": "<leader>e",
    "theme_list": "<leader>t",
    "sidebar_toggle": "<leader>b",
    "status_view": "<leader>s",
    "session_export": "<leader>x",
    "session_new": "<leader>n",
    "session_list": "<leader>l",
    "session_timeline": "<leader>g",
    "session_share": "none",
    "session_unshare": "none",
    "session_interrupt": "escape",
    "session_compact": "<leader>c",
    "session_child_cycle": "ctrl+right",
    "session_child_cycle_reverse": "ctrl+left",
    "messages_page_up": "pageup",
    "messages_page_down": "pagedown",
    "messages_half_page_up": "ctrl+alt+u",
    "messages_half_page_down": "ctrl+alt+d",
    "messages_first": "ctrl+g,home",
    "messages_last": "ctrl+alt+g,end",
    "messages_copy": "<leader>y",
    "messages_undo": "<leader>u",
    "messages_redo": "<leader>r",
    "messages_toggle_conceal": "<leader>h",
    "model_list": "<leader>m",
    "model_cycle_recent": "f2",
    "model_cycle_recent_reverse": "shift+f2",
    "command_list": "ctrl+p",
    "agent_list": "<leader>a",
    "agent_cycle": "tab",
    "agent_cycle_reverse": "shift+tab",
    "input_clear": "ctrl+c",
    "input_forward_delete": "ctrl+d",
    "input_paste": "ctrl+v",
    "input_submit": "enter",
    "input_newline": "shift+enter,ctrl+j",
    "history_previous": "up",
    "history_next": "down"
  }
}
```

--------------------------------

### Add test MCP server locally

Source: https://opencode.ai/docs/mcp-servers

Example of adding the @modelcontextprotocol/server-everything test MCP server as a local server. Shows the minimal configuration needed to run a local MCP server using npx.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "mcp_everything": {
      "type": "local",
      "command": ["npx", "-y", "@modelcontextprotocol/server-everything"],
    },
  },
}
```

--------------------------------

### Authenticate with a provider using opencode auth login

Source: https://opencode.ai/docs/providers

Runs the OpenCode CLI authentication flow for any supported provider. The command opens a credential selector UI where you choose the provider and authentication method. After successful login, the credentials are stored securely.

```Bash
$ opencode auth login
```

--------------------------------

### Disable Specific Formatter

Source: https://opencode.ai/docs/formatters

Configuration example showing how to disable a specific formatter (prettier) while keeping the formatter system enabled. Set the disabled property to true to prevent automatic formatting for that tool.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "formatter": {
    "prettier": {
      "disabled": true
    }
  }
}
```

--------------------------------

### GET /zen/v1/models

Source: https://opencode.ai/docs/zen

Retrieve the full list of available models and their metadata from OpenCode Zen. This endpoint provides information about all supported models.

```APIDOC
## GET /zen/v1/models

### Description
Retrieve the full list of available models and their metadata from OpenCode Zen. This endpoint provides information about all supported models.

### Method
GET

### Endpoint
https://opencode.ai/zen/v1/models

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
None

### Request Example
```
GET https://opencode.ai/zen/v1/models
```

### Response
#### Success Response (200)
- **data** (array) - Array of model objects containing model information

#### Response Example
{
  "data": [
    {
      "id": "gpt-5",
      "name": "GPT 5",
      "endpoint": "https://opencode.ai/zen/v1/responses",
      "sdk": "@ai-sdk/openai"
    },
    {
      "id": "claude-sonnet-4-5",
      "name": "Claude Sonnet 4.5",
      "endpoint": "https://opencode.ai/zen/v1/messages",
      "sdk": "@ai-sdk/anthropic"
    }
  ]
}
```

--------------------------------

### Custom Formatter Configuration

Source: https://opencode.ai/docs/formatters

Advanced configuration with multiple custom formatters including custom command execution, environment variables, and file extensions. Shows how to override built-in formatters and add new ones with custom commands and runtime environments.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "formatter": {
    "prettier": {
      "command": ["npx", "prettier", "--write", "$FILE"],
      "environment": {
        "NODE_ENV": "development"
      },
      "extensions": [".js", ".ts", ".jsx", ".tsx"]
    },
    "custom-markdown-formatter": {
      "command": ["deno", "fmt", "$FILE"],
      "extensions": [".md"]
    }
  }
}
```

--------------------------------

### Re-authenticate with provider

Source: https://opencode.ai/docs/troubleshooting

Command to re-authenticate with a specified provider after clearing configuration or resolving authentication issues.

```Command Line
opencode auth login <provider>
```

--------------------------------

### Add custom tools to OpenCode via plugins

Source: https://opencode.ai/docs/plugins

Plugin that demonstrates creating custom tools using the tool helper from @opencode-ai/plugin package. Shows Zod schema validation, tool description, and async execution handler. Custom tools extend OpenCode's built-in functionality.

```typescript
import { type Plugin, tool } from "@opencode-ai/plugin"

export const CustomToolsPlugin: Plugin = async (ctx) => {
  return {
    tool: {
      mytool: tool({
        description: "This is a custom tool",
        args: {
          foo: tool.schema.string(),
        },
        async execute(args, ctx) {
          return `Hello ${args.foo}!`
        },
      }),
    },
  }
}
```

--------------------------------

### opencode Review and Modify PR Example

Source: https://opencode.ai/docs/github

An example comment to request opencode to review and make changes to a GitHub pull request. The comment provides a specific instruction to delete an attachment from S3.

```text
Delete the attachment from S3 when the note is removed /oc
```

--------------------------------

### opencode Fix Issue Example

Source: https://opencode.ai/docs/github

An example comment for requesting opencode to fix an existing GitHub issue. It triggers opencode to create a new branch, implement the changes, and open a pull request.

```text
/opencode fix this
```

--------------------------------

### Configure Auto-Update - JSON

Source: https://opencode.ai/docs/config

Controls whether OpenCode AI automatically downloads and installs updates on startup. Setting autoupdate to false disables automatic updates, requiring manual update management.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "autoupdate": false
}
```

--------------------------------

### Disable Keybind in JSON Configuration

Source: https://opencode.ai/docs/keybinds

Example showing how to disable a specific keybind by setting its value to "none" in the OpenCode config file. This example disables the session_compact shortcut.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "keybinds": {
    "session_compact": "none"
  }
}
```

--------------------------------

### Create new agent via CLI command

Source: https://opencode.ai/docs/agents

Interactive command to create a new agent configuration. The command guides users through specifying save location, agent description, system prompt generation, tool selection, and creates a markdown configuration file.

```bash
opencode agent create
```

--------------------------------

### Configure Zed Editor for OpenCode ACP

Source: https://opencode.ai/docs/acp

This JSON configuration snippet allows Zed editor to connect to OpenCode as an ACP agent. It specifies the command to run and its arguments. Ensure OpenCode is installed and accessible in your PATH.

```json
{
  "agent_servers": {
    "OpenCode": {
      "command": "opencode",
      "args": ["acp"]
    }
  }
}
```

--------------------------------

### Configure agents via JSON (opencode.json)

Source: https://opencode.ai/docs/agents

Configure built-in and custom agents using JSON format in the opencode.json config file. This example shows how to set up primary agents (Build, Plan) and a custom subagent (code-reviewer) with model selection, tool permissions, and custom prompts.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "agent": {
    "build": {
      "mode": "primary",
      "model": "anthropic/claude-sonnet-4-20250514",
      "prompt": "{file:./prompts/build.txt}",
      "tools": {
        "write": true,
        "edit": true,
        "bash": true
      }
    },
    "plan": {
      "mode": "primary",
      "model": "anthropic/claude-haiku-4-20250514",
      "tools": {
        "write": false,
        "edit": false,
        "bash": false
      }
    },
    "code-reviewer": {
      "description": "Reviews code for best practices and potential issues",
      "mode": "subagent",
      "model": "anthropic/claude-sonnet-4-20250514",
      "prompt": "You are a code reviewer. Focus on security, performance, and maintainability.",
      "tools": {
        "write": false,
        "edit": false
      }
    }
  }
}
```

--------------------------------

### Send system notifications on plugin events

Source: https://opencode.ai/docs/plugins

Notification plugin that listens for session completion events and triggers macOS system notifications using osascript. Demonstrates event hook implementation and external command execution through Bun's shell API.

```javascript
export const NotificationPlugin = async ({ project, client, $, directory, worktree }) => {
  return {
    event: async ({ event }) => {
      // Send notification on session completion
      if (event.type === "session.idle") {
        await $`osascript -e 'display notification "Session completed!" with title "opencode"'`
      }
    },
  }
}
```

--------------------------------

### Configure agents via Markdown files

Source: https://opencode.ai/docs/agents

Define agents using Markdown files with YAML front matter placed in global (.config/opencode/agent/) or per-project (.opencode/agent/) directories. The filename becomes the agent name. This example shows a code reviewer agent with temperature settings and custom instructions.

```yaml
---
description: Reviews code for quality and best practices
mode: subagent
model: anthropic/claude-sonnet-4-20250514
temperature: 0.1
tools:
  write: false
  edit: false
  bash: false
---


You are in code review mode. Focus on:


- Code quality and best practices
- Potential bugs and edge cases
- Performance implications
- Security considerations


Provide constructive feedback without making direct changes.
```

--------------------------------

### Create Custom Tool with Helper

Source: https://opencode.ai/docs/custom-tools

Defines a custom tool using the tool() helper function which provides type-safety and validation. The filename becomes the tool name and arguments are validated using Zod schema. Shows database query tool example.

```typescript
import { tool } from "@opencode-ai/plugin"

export default tool({
  description: "Query the project database",
  args: {
    query: tool.schema.string().describe("SQL query to execute"),
  },
  async execute(args) {
    // Your database logic here
    return `Executed query: ${args.query}`
  },
})
```

```javascript
import { tool } from "@opencode-ai/plugin"

export default tool({
  description: "Query the project database",
  args: {
    query: tool.schema.string().describe("SQL query to execute"),
  },
  async execute(args) {
    // Your database logic here
    return `Executed query: ${args.query}`
  },
})
```

--------------------------------

### Clear provider package cache

Source: https://opencode.ai/docs/troubleshooting

Command to clear cached provider packages, forcing OpenCode to download the latest versions. Resolves API call and compatibility issues.

```Command Line
rm -rf ~/.cache/opencode
```

--------------------------------

### Configure provider-specific options in JSON

Source: https://opencode.ai/docs/agents

Passes additional provider-specific options directly to the model. This example shows OpenAI-specific parameters for reasoning effort and text verbosity. These options vary by model and provider, so check the provider's documentation for available parameters.

```json
{
  "agent": {
    "deep-thinker": {
      "description": "Agent that uses high reasoning effort for complex problems",
      "model": "openai/gpt-5",
      "reasoningEffort": "high",
      "textVerbosity": "low"
    }
  }
}
```

--------------------------------

### Clear OpenCode configuration

Source: https://opencode.ai/docs/troubleshooting

Command to clear corrupted or invalid OpenCode configuration by removing the local storage directory. Requires re-authentication afterward.

```Command Line
rm -rf ~/.local/share/opencode
```

--------------------------------

### Set log level for debugging

Source: https://opencode.ai/docs/troubleshooting

Command to set the log level for detailed debug information in OpenCode AI. Useful for troubleshooting issues by increasing verbosity.

```Command Line
opencode --log-level DEBUG
```

--------------------------------

### Set Google Vertex AI environment variables for opencode

Source: https://opencode.ai/docs/providers

Define the required Google Vertex AI credentials and project information as environment variables before invoking opencode. You can provide them inline for a single command or export them in your shell profile for reuse. Ensure the service account JSON path is accessible.

```bash
GOOGLE_APPLICATION_CREDENTIALS=/path/to/service-account.json GOOGLE_VERTEX_PROJECT=your-project-id opencode
```

```bash
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service-account.json
export GOOGLE_VERTEX_PROJECT=your-project-id
export GOOGLE_VERTEX_REGION=us-central1
```

--------------------------------

### Configure Context7 MCP server

Source: https://opencode.ai/docs/mcp-servers

Configures the Context7 MCP server for document search functionality. The basic setup connects to the remote server, while the advanced configuration includes API key authentication using environment variables for higher rate limits.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "context7": {
      "type": "remote",
      "url": "https://mcp.context7.com/mcp"
    }
  }
}
```

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "context7": {
      "type": "remote",
      "url": "https://mcp.context7.com/mcp",
      "headers": {
        "CONTEXT7_API_KEY": "{env:CONTEXT7_API_KEY}"
      }
    }
  }
}
```

--------------------------------

### Configure LM Studio provider in opencode JSON config

Source: https://opencode.ai/docs/providers

Add a custom provider entry for LM Studio in the opencode configuration file. The JSON specifies the npm package, display name, endpoint URL, and model mappings. Place this file at the project root or reference it via the `--config` flag.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "lmstudio": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "LM Studio (local)",
      "options": {
        "baseURL": "http://127.0.0.1:1234/v1"
      },
      "models": {
        "google/gemma-3n-e4b": {
          "name": "Gemma 3n-e4b (local)"
        }
      }
    }
  }
}
```

--------------------------------

### Configure provider baseURL in OpenCode config

Source: https://opencode.ai/docs/providers

Defines a custom base URL for a provider in the OpenCode configuration file. The JSON snippet follows the OpenCode config schema and can be placed in `~/.config/opencode/config.json`. Adjust the provider name and URL as needed.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "anthropic": {
      "options": {
        "baseURL": "https://api.anthropic.com/v1"
      }
    }
  }
}
```

--------------------------------

### Control Web Fetch Permissions in JSON

Source: https://opencode.ai/docs/permissions

This JSON configuration manages permissions for the webfetch tool in OpenCode, setting it to 'ask' for approval on web page fetches. It directly controls LLM web access, requiring no other inputs. The setup limits web interaction to approved requests only.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "webfetch": "ask"
  }
}
```

--------------------------------

### Override Permissions for Specific Agents in JSON

Source: https://opencode.ai/docs/permissions

This JSON example overrides global permissions for the 'build' agent, allowing 'git push' while keeping it 'ask' globally. It uses an 'agent' object for per-agent configs, enabling customized access. Dependencies include matching agent names, with outputs tailored to agent-specific needs.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "bash": {
      "git push": "ask"
    }
  },
  "agent": {
    "build": {
      "permission": {
        "bash": {
          "git push": "allow"
        }
      }
    }
  }
}
```

--------------------------------

### Override Agent Model in JSON

Source: https://opencode.ai/docs/agents

This option overrides the default model for the agent, allowing selection of task-optimized models. For example, faster models for planning or capable ones for implementation. Specify model identifiers directly.

```json
{
  "agent": {
    "plan": {
      "model": "anthropic/claude-haiku-4-20250514"
    }
  }
}
```

--------------------------------

### Create Opencode Client (Existing Server)

Source: https://opencode.ai/docs/sdk

Create a client instance to connect to an already running opencode server. Specify the server's base URL. Options for server hostname, port, abort signal, timeout, and configuration can also be provided.

```typescript
import { createOpencodeClient } from "@opencode-ai/sdk"


const client = createOpencodeClient({
  baseUrl: "http://localhost:4096",
})
```

--------------------------------

### Deny All Bash Commands Except Specific Ones in JSON

Source: https://opencode.ai/docs/permissions

This setup denies all bash commands by default with '*' as 'deny', then overrides for 'pwd' as 'allow' and 'git status' as 'ask'. It uses wildcard precedence, allowing exceptions without dependencies. Inputs are the wildcard and override rules to manage comprehensive access.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "bash": {
      "*": "deny",
      "pwd": "allow",
      "git status": "ask"
    }
  }
}
```

--------------------------------

### Create and Manage Sessions (JavaScript)

Source: https://opencode.ai/docs/sdk

Demonstrates how to create, list, and interact with sessions using the `client.session` methods. It showcases creating a session, listing sessions, sending prompt messages, and injecting context without AI response.

```javascript
const session = await client.session.create({
  body: { title: "My session" },
});

const sessions = await client.session.list();

// Send a prompt message
const result = await client.session.prompt({
  path: { id: session.id },
  body: { model: { providerID: "anthropic", modelID: "claude-3-5-sonnet-20241022" }, parts: [{ type: "text", text: "Hello!" }], },
});

// Inject context without triggering AI response (useful for plugins)
await client.session.prompt({
  path: { id: session.id },
  body: { noReply: true, parts: [{ type: "text", text: "You are a helpful assistant." }], },
});
```

--------------------------------

### Protect .env files from being read

Source: https://opencode.ai/docs/plugins

Security plugin that prevents OpenCode from accessing .env files by implementing a tool execution hook. Uses conditional logic to detect .env file access attempts and throw errors, protecting sensitive environment variables.

```javascript
export const EnvProtection = async ({ project, client, $, directory, worktree }) => {
  return {
    "tool.execute.before": async (input, output) => {
      if (input.tool === "read" && output.args.filePath.includes(".env")) {
        throw new Error("Do not read .env files")
      },
    },
  }
}
```

--------------------------------

### Security auditor agent configuration in YAML

Source: https://opencode.ai/docs/agents

Example configuration for a security auditor agent that identifies vulnerabilities. Uses subagent mode and disables write and edit tools. Includes a system prompt focused on security expertise with specific areas of focus for vulnerability identification.

```yaml
---
description: Performs security audits and identifies vulnerabilities
mode: subagent
tools:
  write: false
  edit: false
---


You are a security expert. Focus on identifying potential security issues.


Look for:


- Input validation vulnerabilities
- Authentication and authorization flaws
- Data exposure risks
- Dependency vulnerabilities
- Configuration security issues
```

--------------------------------

### Configure List Tool

Source: https://opencode.ai/docs/tools

Shows how to enable the `list` tool for listing files and directories. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "list": true
  }
}
```

--------------------------------

### Force Subagent Invocation with Subtask in JSON

Source: https://opencode.ai/docs/commands

This boolean config forces the command to invoke a subagent, preventing context pollution. It's useful for isolating tasks and overrides default modes. Applicable in multi-agent setups like OpenCode AI.

```json
{
  "command": {
    "analyze": {
      "subtask": true
    }
  }
}
```

--------------------------------

### Configure Webfetch Tool

Source: https://opencode.ai/docs/tools

Shows how to enable the `webfetch` tool for fetching web content. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "webfetch": true
  }
}
```

--------------------------------

### Disable Share Feature (JSON)

Source: https://opencode.ai/docs/enterprise

This JSON configuration disables the optional /share feature in OpenCode to prevent data from leaving the organization. It uses a schema reference for validation and requires no external dependencies. The output ensures conversations remain local; it has no limitations on enterprise setups.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "share": "disabled"
}
```

--------------------------------

### Configure and Create Opencode Instance

Source: https://opencode.ai/docs/sdk

Create a customized opencode instance by passing a configuration object. This allows overriding default settings like hostname, port, and model. The instance can then be used to interact with the server, which is then closed.

```typescript
import { createOpencode } from "@opencode-ai/sdk"


const opencode = await createOpencode({
  hostname: "127.0.0.1",
  port: 4096,
  config: {
    model: "anthropic/claude-3-5-sonnet-20241022",
  },
})


console.log(`Server running at ${opencode.server.url}`)


opencode.server.close()
```

--------------------------------

### Configure providers and models (JSON)

Source: https://opencode.ai/docs/config

Illustrates configuring the provider block and selecting models including main and small_model. Useful for routing to a specific model and model variant across tasks.

```json
{\n  \"$schema\": \"https://opencode.ai/config.json\",\n  \"provider\": {},\n  \"model\": \"anthropic/claude-sonnet-4-20250514\",\n  \"small_model\": \"anthropic/claude-3-5-haiku-20241022\"\n}
```

--------------------------------

### Run OpenCode with custom config path or directory (bash)

Source: https://opencode.ai/docs/config

Demonstrates using environment variables to point OpenCode at a custom config location or directory before running. Inputs are the environment variable assignments and the command invocation; outputs are the program run with the specified config.

```bash
export OPENCODE_CONFIG=/path/to/my/custom-config.json\nopencode run \"Hello world\"
```

```bash
export OPENCODE_CONFIG_DIR=/path/to/my/config-directory\nopencode run \"Hello world\"
```

--------------------------------

### Listen to Real-Time Events (JavaScript)

Source: https://opencode.ai/docs/sdk

Shows how to subscribe to server-sent events using `client.event.subscribe` and iterate through the incoming events.

```javascript
// Listen to real-time events
const events = await client.event.subscribe();
for await (const event of events.stream) {
  console.log("Event:", event.type, event.properties);
}
```

--------------------------------

### Configure Todowrite Tool

Source: https://opencode.ai/docs/tools

Shows how to enable the `todowrite` tool for managing todo lists. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "todowrite": true
  }
}
```

--------------------------------

### Configure Instructions - JSON

Source: https://opencode.ai/docs/config

Specifies instruction files and glob patterns for AI behavior guidelines. Accepts an array of file paths and patterns to load custom instructions from markdown files, enabling personalized AI guidance.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "instructions": ["CONTRIBUTING.md", "docs/guidelines.md", ".cursor/rules/*.md"]
}
```

--------------------------------

### Sessions API

Source: https://opencode.ai/docs/sdk

Provides methods for creating, managing, and interacting with AI sessions including prompt handling, command execution, and session lifecycle management.

```APIDOC
## Sessions API

### Description
The Sessions API provides comprehensive session management capabilities for creating, retrieving, updating, and interacting with AI sessions. It supports prompt handling, command execution, shell operations, and message management.

### Methods
- **POST** `session.create({ body })` - Create a new session
- **GET** `session.list()` - List all sessions
- **GET** `session.get({ path })` - Get session by ID
- **GET** `session.children({ path })` - List child sessions
- **DELETE** `session.delete({ path })` - Delete a session
- **PUT** `session.update({ path, body })` - Update session properties
- **POST** `session.init({ path, body })` - Analyze app and create AGENTS.md
- **POST** `session.abort({ path })` - Abort a running session
- **POST** `session.share({ path })` - Share session
- **POST** `session.unshare({ path })` - Unshare session
- **POST** `session.summarize({ path, body })` - Summarize session
- **GET** `session.messages({ path })` - List messages in a session
- **GET** `session.message({ path })` - Get message details
- **POST** `session.prompt({ path, body })` - Send prompt message
- **POST** `session.command({ path, body })` - Send command to session
- **POST** `session.shell({ path, body })` - Run a shell command
- **POST** `session.revert({ path, body })` - Revert a message
- **POST** `session.unrevert({ path })` - Restore reverted messages
- **POST** `postSessionByIdPermissionsByPermissionId({ path, body })` - Respond to permission request

### Parameters
#### Path Parameters
- **id** (string) - Required - Session identifier
- **permissionId** (string) - Required - Permission request identifier

#### Request Body
- **title** (string) - Optional - Session title for creation
- **model** (object) - Optional - Model configuration with providerID and modelID
- **parts** (array) - Required - Message parts array with type and text
- **noReply** (boolean) - Optional - If true, returns UserMessage without AI response
- **text** (string) - Optional - Text content for TUI operations
- **command** (string) - Optional - Command to execute

### Request Example
{
  "body": {
    "title": "My session",
    "model": {
      "providerID": "anthropic",
      "modelID": "claude-3-5-sonnet-20241022"
    },
    "parts": [
      {
        "type": "text",
        "text": "Hello!"
      }
    ]
  }
}

### Response
#### Success Response (200)
- **Session** (object) - Session object with id, title, and properties
- **Message** (object) - Message object with info and parts
- **AssistantMessage** (object) - AI response message
- **boolean** - Operation success status

#### Response Example
{
  "id": "session-123",
  "title": "My session",
  "status": "active",
  "info": {
    "id": "msg-456",
    "type": "assistant",
    "timestamp": "2024-01-01T12:00:00Z"
  },
  "parts": [
    {
      "type": "text",
      "text": "Hello! How can I help you today?"
    }
  ]
}
```

--------------------------------

### Configure Todoread Tool

Source: https://opencode.ai/docs/tools

Demonstrates enabling the `todoread` tool for reading existing todo lists. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "todoread": true
  }
}
```

--------------------------------

### POST /zen/v1/chat/completions

Source: https://opencode.ai/docs/zen

Unified OpenAI-compatible endpoint for various models including GLM 4.6, Kimi K2, Qwen3 Coder 480B, and Grok Code Fast 1. This endpoint works with OpenAI-compatible SDKs.

```APIDOC
## POST /zen/v1/chat/completions

### Description
Unified OpenAI-compatible endpoint for various models including GLM 4.6, Kimi K2, Qwen3 Coder 480B, and Grok Code Fast 1. This endpoint works with OpenAI-compatible SDKs.

### Method
POST

### Endpoint
https://opencode.ai/zen/v1/chat/completions

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
- **model** (string) - Required - The model identifier (e.g., glm-4.6, kimi-k2)
- **messages** (array) - Required - Array of message objects
- **stream** (boolean) - Optional - Whether to stream the response
- **temperature** (number) - Optional - Sampling temperature

### Request Example
{
  "model": "qwen3-coder",
  "messages": [
    {
      "role": "user",
      "content": "Write a Python function to sort a list"
    }
  ],
  "temperature": 0.7,
  "stream": false
}

### Response
#### Success Response (200)
- **id** (string) - Unique identifier for the response
- **choices** (array) - Array of completion choices
- **usage** (object) - Token usage statistics

#### Response Example
{
  "id": "chat_123",
  "choices": [
    {
      "message": {
        "role": "assistant",
        "content": "def sort_list(arr):\n    return sorted(arr)"
      }
    }
  ],
  "usage": {
    "prompt_tokens": 20,
    "completion_tokens": 35,
    "total_tokens": 55
  }
}
```

--------------------------------

### Create custom command with markdown

Source: https://opencode.ai/docs/commands

Defines a custom command using a markdown file with frontmatter configuration. The file name determines the command name and the content becomes the prompt template. Supports description, agent, and model configuration options.

```markdown
---
description: Run tests with coverage
agent: build
model: anthropic/claude-3-5-sonnet-20241022
---


Run the full test suite with coverage report and show any failures.
Focus on the failing tests and suggest fixes.

```

--------------------------------

### Create project-specific custom theme directory

Source: https://opencode.ai/docs/themes

Commands to create a directory for project-specific custom themes in the project's .opencode directory.

```shell
mkdir -p .opencode/themes
vim .opencode/themes/my-theme.json
```

--------------------------------

### Control TUI Interface (JavaScript)

Source: https://opencode.ai/docs/sdk

Shows how to utilize the `client.tui` methods to append text to prompts, display toast notifications, and interact with the TUI.

```javascript
await client.tui.appendPrompt({
  body: { text: "Add this to prompt" },
});

await client.tui.showToast({
  body: { message: "Task completed", variant: "success" },
});
```

--------------------------------

### Create user-wide custom theme directory

Source: https://opencode.ai/docs/themes

Commands to create a directory for user-wide custom themes in the OpenCode configuration directory.

```shell
mkdir -p ~/.config/opencode/themes
vim ~/.config/opencode/themes/my-theme.json
```

--------------------------------

### Configure Glob Tool

Source: https://opencode.ai/docs/tools

Demonstrates enabling the `glob` tool for finding files by pattern matching. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "glob": true
  }
}
```

--------------------------------

### Configure Write Tool

Source: https://opencode.ai/docs/tools

Shows how to enable the `write` tool for creating or overwriting files. This allows the LLM to create new files. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "write": true
  }
}
```

--------------------------------

### Apply global OpenCode config (JSONC)

Source: https://opencode.ai/docs/config

Shows a global OpenCode configuration with theme, model, and autoupdate using JSON with Comments. The snippet includes a theme designation and a sample model and autoupdate flag to illustrate how these settings affect the global config.

```jsonc
{\n  \"$schema\": \"https://opencode.ai/config.json\",\n  // Theme configuration\n  \"theme\": \"opencode\",\n  \"model\": \"anthropic/claude-sonnet-4-20250514\",\n  \"autoupdate\": true,\n}
```

--------------------------------

### Auth API

Source: https://opencode.ai/docs/sdk

Provides authentication credential management for setting up API keys and authentication providers.

```APIDOC
## Auth API

### Description
The Auth API manages authentication credentials for various AI service providers, allowing you to set up and configure API keys and authentication settings.

### Methods
- **POST** `auth.set({ path, body })` - Set authentication credentials

### Parameters
#### Path Parameters
- **id** (string) - Required - Provider identifier (e.g., "anthropic")

#### Request Body
- **type** (string) - Required - Authentication type (e.g., "api")
- **key** (string) - Required - API key or authentication token

### Request Example
{
  "path": {
    "id": "anthropic"
  },
  "body": {
    "type": "api",
    "key": "your-api-key"
  }
}

### Response
#### Success Response (200)
- **boolean** - Operation success status

#### Response Example
{
  "success": true
}
```

--------------------------------

### App APIs

Source: https://opencode.ai/docs/server

Endpoints for retrieving app information and initializing the application. These APIs provide basic interaction with the opencode app state.

```APIDOC
## GET /app

### Description
Get app info.

### Method
GET

### Endpoint
/app

### Parameters
No parameters.

### Response
#### Success Response (200)
- App object.

#### Response Example
{
  "type": "App"
}

## POST /app/init

### Description
Initialize the app.

### Method
POST

### Endpoint
/app/init

### Parameters
No parameters.

### Request Body
No body required.

### Response
#### Success Response (200)
- boolean value indicating success.
```

--------------------------------

### Configure Patch Tool

Source: https://opencode.ai/docs/tools

Demonstrates enabling the `patch` tool for applying patches to files. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "patch": true
  }
}
```

--------------------------------

### Config API

Source: https://opencode.ai/docs/sdk

The Config API allows retrieving configuration information and listing providers and default models.

```APIDOC
## Config

### Description
The Config API allows retrieving configuration information and listing providers and default models.

### Method
APIDOC

### Endpoint
N/A - These are client-side methods.

### Parameters
#### Path Parameters
- N/A

#### Query Parameters
- N/A

#### Request Body
- N/A

### Get Config Info
#### Description
Retrieves the configuration information.
#### Response
#### Success Response (200)
- **config** (Config) - Configuration information.
#### Response Example
{
  "hostname": "127.0.0.1",
  "port": 4096
}

### List Providers and Defaults
#### Description
Lists available providers and their default models.
#### Response
#### Success Response (200)
- **providers** (Provider[]) - An array of provider objects.
- **default** ({ [key: string]: string }) - An object containing default models for each provider.
#### Response Example
{
  "providers": [
    {
      "name": "anthropic",
      "defaultModel": "claude-3-5-sonnet"
    }
  ],
  "default": {
    "anthropic": "claude-3-5-sonnet"
  }
}
```

--------------------------------

### Configure Custom Commands - JSON

Source: https://opencode.ai/docs/config

Defines custom commands for repetitive tasks with templates, descriptions, agents, and models. Each command includes a template for the action, a human-readable description, the agent to use, and the specific model to invoke.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "command": {
    "test": {
      "template": "Run the full test suite with coverage report and show any failures.\nFocus on the failing tests and suggest fixes.",
      "description": "Run tests with coverage",
      "agent": "build",
      "model": "anthropic/claude-3-5-sonnet-20241022",
    },
    "component": {
      "template": "Create a new React component named $ARGUMENTS with TypeScript support.\nInclude proper typing and basic structure.",
      "description": "Create a new component",
    },
  },
}
```

--------------------------------

### Search and Read Files (JavaScript)

Source: https://opencode.ai/docs/sdk

Illustrates searching for text within files, finding files by name, and reading the content of an existing file using `client.find` and `client.file`.

```javascript
const textResults = await client.find.text({
  query: { pattern: "function.*opencode" },
});

const files = await client.find.files({
  query: { query: "*.ts" },
});

const content = await client.file.read({
  query: { path: "src/index.ts" },
});
```

--------------------------------

### Configure Tools in Markdown

Source: https://opencode.ai/docs/tools

Demonstrates configuring tools within Markdown for specific agents, allowing for more readable format. Configuration are set in the format of key-value pairs.

```Markdown
---
description: Read-only analysis agent
mode: subagent
tools:
  write: false
  edit: false
  bash: false---

Analyze code without making any modifications.

```

--------------------------------

### File Content Substitution - JSON

Source: https://opencode.ai/docs/config

Shows how to substitute file contents using {file:path/to/file} syntax. Supports both relative and absolute paths, enabling secure management of sensitive data like API keys and inclusion of large instruction files.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "instructions": ["./custom-instructions.md"],
  "provider": {
    "openai": {
      "options": {
        "apiKey": "{file:~/.secrets/openai-key}"
      }
    }
  }
}
```

--------------------------------

### Configure Grep Tool

Source: https://opencode.ai/docs/tools

Shows how to enable the `grep` tool for searching file contents using regular expressions. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "grep": true
  }
}
```

--------------------------------

### Configure Tools with Wildcards

Source: https://opencode.ai/docs/tools

Shows how to use wildcards (e.g., `mymcp_*`) to configure multiple tools at once. This is useful for disabling related tools efficiently. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "mymcp_*": false
  }
}
```

--------------------------------

### Environment Variable Substitution - JSON

Source: https://opencode.ai/docs/config

Demonstrates using environment variables in configuration files through {env:VARIABLE_NAME} syntax. Environment variables are replaced with their values at runtime, with unset variables becoming empty strings.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "model": "{env:OPENCODE_MODEL}",
  "provider": {
    "anthropic": {
      "models": {},
      "options": {
        "apiKey": "{env:ANTHROPIC_API_KEY}"
      }
    }
  }
}
```

--------------------------------

### Configure Read Tool

Source: https://opencode.ai/docs/tools

Demonstrates enabling the `read` tool for reading file contents. Supports reading line ranges. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "read": true
  }
}
```

--------------------------------

### Project API

Source: https://opencode.ai/docs/sdk

The Project API allows listing and retrieving information about projects within the Opencode system.

```APIDOC
## Project

### Description
The Project API allows listing and retrieving information about projects within the Opencode system.

### Method
APIDOC

### Endpoint
N/A - These are client-side methods.

### Parameters
#### Path Parameters
- N/A

#### Query Parameters
- N/A

#### Request Body
- N/A

### List Projects
#### Description
Lists all projects.
#### Response
#### Success Response (200)
- **projects** (Project[]) - An array of projects.
#### Response Example
[
  {
    "id": "project_1",
    "name": "Project 1"
  },
  {
    "id": "project_2",
    "name": "Project 2"
  }
]

### Get Current Project
#### Description
Retrieves the current project.
#### Response
#### Success Response (200)
- **project** (Project) - The current project details.
#### Response Example
{
  "id": "current_project_id",
  "name": "Current Project"
}
```

--------------------------------

### Use command arguments in prompt

Source: https://opencode.ai/docs/commands

Shows how to pass arguments to custom commands using the $ARGUMENTS placeholder or positional parameters like $1, $2, etc. Arguments are specified when running the command and automatically substituted in the prompt template.

```markdown
---
description: Create a new file with content
---


Create a file named $1 in the directory $2
with the following content: $3

```

--------------------------------

### Configure Avante.nvim with Environment Variables for OpenCode ACP

Source: https://opencode.ai/docs/acp

This Lua configuration snippet for Avante.nvim includes setting environment variables, such as OPENCODE_API_KEY, for the OpenCode ACP provider. This is useful for authenticating with the OpenCode API or configuring its behavior. The environment variable is fetched using `os.getenv`.

```lua
{
  acp_providers = {
    ["opencode"] = {
      command = "opencode",
      args = { "acp" },
      env = {
        OPENCODE_API_KEY = os.getenv("OPENCODE_API_KEY")
      }
    }
  }
}
```

--------------------------------

### Configure Keybindings - JSON

Source: https://opencode.ai/docs/config

Sets up custom keybindings for OpenCode AI functionality. The keybinds option allows users to customize keyboard shortcuts. Currently shows an empty configuration object that can be populated with key-value pairs.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "keybinds": {}
}
```

--------------------------------

### POST /zen/v1/responses

Source: https://opencode.ai/docs/zen

Endpoint for GPT models including GPT-5 and GPT-5 Codex. This endpoint handles completion requests for these specific models.

```APIDOC
## POST /zen/v1/responses

### Description
Endpoint for GPT models including GPT-5 and GPT-5 Codex. This endpoint handles completion requests for these specific models.

### Method
POST

### Endpoint
https://opencode.ai/zen/v1/responses

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
- **model** (string) - Required - The model identifier (e.g., gpt-5, gpt-5-codex)
- **messages** (array) - Required - Array of message objects
- **stream** (boolean) - Optional - Whether to stream the response

### Request Example
{
  "model": "gpt-5-codex",
  "messages": [
    {
      "role": "user",
      "content": "Write a function to calculate fibonacci sequence"
    }
  ],
  "stream": false
}

### Response
#### Success Response (200)
- **id** (string) - Unique identifier for the response
- **choices** (array) - Array of completion choices
- **usage** (object) - Token usage statistics

#### Response Example
{
  "id": "response_123",
  "choices": [
    {
      "message": {
        "role": "assistant",
        "content": "Here's a function to calculate fibonacci sequence..."
      }
    }
  ],
  "usage": {
    "prompt_tokens": 25,
    "completion_tokens": 120,
    "total_tokens": 145
  }
}
```

--------------------------------

### Export Multiple Tools from File

Source: https://opencode.ai/docs/custom-tools

Shows how to export multiple tools from a single file where each export becomes a separate tool with filename_exportname naming convention. Demonstrates math operations with separate add and multiply tools.

```typescript
import { tool } from "@opencode-ai/plugin"

export const add = tool({
  description: "Add two numbers",
  args: {
    a: tool.schema.number().describe("First number"),
    b: tool.schema.number().describe("Second number"),
  },
  async execute(args) {
    return args.a + args.b
  },
})

export const multiply = tool({
  description: "Multiply two numbers",
  args: {
    a: tool.schema.number().describe("First number"),
    b: tool.schema.number().describe("Second number"),
  },
  async execute(args) {
    return args.a * args.b
  },
})
```

```javascript
import { tool } from "@opencode-ai/plugin"

export const add = tool({
  description: "Add two numbers",
  args: {
    a: tool.schema.number().describe("First number"),
    b: tool.schema.number().describe("Second number"),
  },
  async execute(args) {
    return args.a + args.b
  },
})

export const multiply = tool({
  description: "Multiply two numbers",
  args: {
    a: tool.schema.number().describe("First number"),
    b: tool.schema.number().describe("Second number"),
  },
  async execute(args) {
    return args.a * args.b
  },
})
```

--------------------------------

### Inject shell command output

Source: https://opencode.ai/docs/commands

Demonstrates how to inject bash command output into custom command prompts using the !`command` syntax. The command runs in the project root and its output becomes part of the prompt sent to the LLM.

```markdown
---
description: Analyze test coverage
---


Here are the current test results:
!`npm test`


Based on these results, suggest improvements to increase coverage.

```

--------------------------------

### App API

Source: https://opencode.ai/docs/sdk

The App API provides methods for logging events and listing available agents within the Opencode system.

```APIDOC
## App

### Description
The App API provides methods for logging events and listing available agents within the Opencode system.

### Method
APIDOC

### Endpoint
N/A - These are client-side methods.

### Parameters
#### Path Parameters
- N/A

#### Query Parameters
- N/A

#### Request Body
- **log.service** (string) - Required - Service name for the log entry.
- **log.level** (string) - Required - Log level (e.g., info, warning, error).
- **log.message** (string) - Required - The log message.

### Request Example
{
  "service": "my-app",
  "level": "info",
  "message": "Operation completed"
}

### Response
#### Success Response (200)
- **log** (boolean) - Indicates success.

#### Response Example
{
  "log": true
}

### Agents API
#### Description
Lists all available agents.
#### Method
APIDOC
#### Endpoint
N/A - client-side method
#### Parameters
#### Path Parameters
- N/A
#### Query Parameters
- N/A
#### Request Body
- N/A
#### Response
#### Success Response (200)
- **agents** (Agent[]) - An array of agents.
#### Response Example
[
  {
    "name": "agent_1",
    "description": "Agent 1 description"
  },
  {
    "name": "agent_2",
    "description": "Agent 2 description"
  }
]

```

--------------------------------

### Configure MCP Servers - JSON

Source: https://opencode.ai/docs/config

Configures Model Context Protocol (MCP) servers for extended functionality. The mcp option accepts server configurations that enable additional capabilities and integrations.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {}
}
```

--------------------------------

### Configure Bash Tool

Source: https://opencode.ai/docs/tools

Shows how to enable the `bash` tool for executing shell commands.  This allows control over shell execution within the project environment. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "bash": true
  }
}
```

--------------------------------

### Configure Grep by Vercel MCP server

Source: https://opencode.ai/docs/mcp-servers

Sets up the Grep by Vercel MCP server for searching code snippets on GitHub. The configuration connects to the remote grep.app service. Users can invoke the tool by referencing 'gh_grep' in their prompts.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "gh_grep": {
      "type": "remote",
      "url": "https://mcp.grep.app"
    }
  }
}
```

--------------------------------

### Configure custom command with JSON

Source: https://opencode.ai/docs/commands

Defines a custom command using JSON configuration in the OpenCode config file. The object key becomes the command name and requires a template property for the prompt. Optional description, agent, and model properties can be specified.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "command": {
    "test": {
      "template": "Run the full test suite with coverage report and show any failures.\nFocus on the failing tests and suggest fixes.",
      "description": "Run tests with coverage",
      "agent": "build",
      "model": "anthropic/claude-3-5-sonnet-20241022"
    }
  }
}
```

--------------------------------

### Configure Avante.nvim for OpenCode ACP

Source: https://opencode.ai/docs/acp

This Lua configuration snippet sets up Avante.nvim to use OpenCode as an ACP provider. It defines the command and arguments for the 'opencode' agent. This allows Avante.nvim to leverage OpenCode's AI capabilities.

```lua
{
  acp_providers = {
    ["opencode"] = {
      command = "opencode",
      args = { "acp" }
    }
  }
}
```

--------------------------------

### Configure Permissions - JSON

Source: https://opencode.ai/docs/config

Defines permission requirements for various operations like editing and bash commands. Setting operations to 'ask' requires user approval before execution, enhancing security and control over AI actions.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "edit": "ask",
    "bash": "ask"
  }
}
```

--------------------------------

### Config APIs

Source: https://opencode.ai/docs/server

Endpoints for accessing configuration information, including providers and default models. Useful for understanding available AI providers.

```APIDOC
## GET /config

### Description
Get config info.

### Method
GET

### Endpoint
/config

### Parameters
No parameters.

### Response
#### Success Response (200)
- Config object.

#### Response Example
{
  "type": "Config"
}

## GET /config/providers

### Description
List providers and default models.

### Method
GET

### Endpoint
/config/providers

### Parameters
No parameters.

### Response
#### Success Response (200)
- Object with providers array and default models.
```

--------------------------------

### Files API

Source: https://opencode.ai/docs/sdk

Provides file search, reading, and status operations for working with workspace files and content.

```APIDOC
## Files API

### Description
The Files API provides comprehensive file operations including text search, file discovery, symbol search, file reading, and status tracking for workspace files.

### Methods
- **POST** `find.text({ query })` - Search for text in files
- **POST** `find.files({ query })` - Find files by name
- **POST** `find.symbols({ query })` - Find workspace symbols
- **POST** `file.read({ query })` - Read a file
- **GET** `file.status({ query? })` - Get status for tracked files

### Parameters
#### Path Parameters
- **path** (string) - Optional - File path for read operations

#### Query Parameters
- **pattern** (string) - Required - Search pattern for text search
- **query** (string) - Required - Query string for file search
- **type** (string) - Optional - File type filter

### Request Example
{
  "query": {
    "pattern": "function.*opencode"
  }
}

### Response
#### Success Response (200)
- **Array<match>** (array) - Text search results with path, lines, line_number, absolute_offset, submatches
- **string[]** (array) - Array of file paths
- **Symbol[]** (array) - Array of symbol objects
- **File** (object) - File object with type, content
- **File[]** (array) - Array of tracked files

#### Response Example
{
  "matches": [
    {
      "path": "src/index.ts",
      "lines": "export function opencode() {",
      "line_number": 42,
      "absolute_offset": 1024,
      "submatches": [
        {
          "match": "opencode",
          "start": 14,
          "end": 22
        }
      ]
    }
  ]
}
```

--------------------------------

### Configure local MCP server

Source: https://opencode.ai/docs/mcp-servers

Shows how to configure a local MCP server in OpenCode with command execution and environment variables. Includes required fields like type and command, and optional fields like environment variables and timeout.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "my-local-mcp-server": {
      "type": "local",
      // Or ["bun", "x", "my-mcp-command"]
      "command": ["npx", "-y", "my-mcp-command"],
      "enabled": true,
      "environment": {
        "MY_ENV_VAR": "my_env_var_value",
      },
    },
  },
}
```

--------------------------------

### OpenCode TUI Configuration - JSON

Source: https://opencode.ai/docs/tui

JSON configuration file to customize TUI behavior, such as scroll speed. Use the specified schema. Inputs are configuration options. Limitations include available options only.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "tui": {
    "scroll_speed": 3
  }
}
```

--------------------------------

### Configure MCP servers in OpenCode

Source: https://opencode.ai/docs/mcp-servers

Defines the basic structure for configuring MCP servers in the OpenCode configuration file. Shows how to define multiple MCP servers with unique names and enable/disable them.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "name-of-mcp-server": {
      // ...
      "enabled": true,
    },
    "name-of-other-mcp-server": {
      // ...
    },
  },
}
```

--------------------------------

### Nord Theme Configuration in JSON

Source: https://opencode.ai/docs/themes

This JSON file defines a custom theme for OpenCode AI using the Nord color scheme. It references the schema at https://opencode.ai/theme.json and includes color definitions in the 'defs' section, followed by theme properties for UI elements like primary, secondary, text, backgrounds, and syntax highlighting in both dark and light modes. The structure allows for easy customization of colors without hardcoding hex values everywhere, though it requires the referenced schema for validation.

```json
{
  "$schema": "https://opencode.ai/theme.json",
  "defs": {
    "nord0": "#2E3440",
    "nord1": "#3B4252",
    "nord2": "#434C5E",
    "nord3": "#4C566A",
    "nord4": "#D8DEE9",
    "nord5": "#E5E9F0",
    "nord6": "#ECEFF4",
    "nord7": "#8FBCBB",
    "nord8": "#88C0D0",
    "nord9": "#81A1C1",
    "nord10": "#5E81AC",
    "nord11": "#BF616A",
    "nord12": "#D08770",
    "nord13": "#EBCB8B",
    "nord14": "#A3BE8C",
    "nord15": "#B48EAD"
  },
  "theme": {
    "primary": {
      "dark": "nord8",
      "light": "nord10"
    },
    "secondary": {
      "dark": "nord9",
      "light": "nord9"
    },
    "accent": {
      "dark": "nord7",
      "light": "nord7"
    },
    "error": {
      "dark": "nord11",
      "light": "nord11"
    },
    "warning": {
      "dark": "nord12",
      "light": "nord12"
    },
    "success": {
      "dark": "nord14",
      "light": "nord14"
    },
    "info": {
      "dark": "nord8",
      "light": "nord10"
    },
    "text": {
      "dark": "nord4",
      "light": "nord0"
    },
    "textMuted": {
      "dark": "nord3",
      "light": "nord1"
    },
    "background": {
      "dark": "nord0",
      "light": "nord6"
    },
    "backgroundPanel": {
      "dark": "nord1",
      "light": "nord5"
    },
    "backgroundElement": {
      "dark": "nord1",
      "light": "nord4"
    },
    "border": {
      "dark": "nord2",
      "light": "nord3"
    },
    "borderActive": {
      "dark": "nord3",
      "light": "nord2"
    },
    "borderSubtle": {
      "dark": "nord2",
      "light": "nord3"
    },
    "diffAdded": {
      "dark": "nord14",
      "light": "nord14"
    },
    "diffRemoved": {
      "dark": "nord11",
      "light": "nord11"
    },
    "diffContext": {
      "dark": "nord3",
      "light": "nord3"
    },
    "diffHunkHeader": {
      "dark": "nord3",
      "light": "nord3"
    },
    "diffHighlightAdded": {
      "dark": "nord14",
      "light": "nord14"
    },
    "diffHighlightRemoved": {
      "dark": "nord11",
      "light": "nord11"
    },
    "diffAddedBg": {
      "dark": "#3B4252",
      "light": "#E5E9F0"
    },
    "diffRemovedBg": {
      "dark": "#3B4252",
      "light": "#E5E9F0"
    },
    "diffContextBg": {
      "dark": "nord1",
      "light": "nord5"
    },
    "diffLineNumber": {
      "dark": "nord2",
      "light": "nord4"
    },
    "diffAddedLineNumberBg": {
      "dark": "#3B4252",
      "light": "#E5E9F0"
    },
    "diffRemovedLineNumberBg": {
      "dark": "#3B4252",
      "light": "#E5E9F0"
    },
    "markdownText": {
      "dark": "nord4",
      "light": "nord0"
    },
    "markdownHeading": {
      "dark": "nord8",
      "light": "nord10"
    },
    "markdownLink": {
      "dark": "nord9",
      "light": "nord9"
    },
    "markdownLinkText": {
      "dark": "nord7",
      "light": "nord7"
    },
    "markdownCode": {
      "dark": "nord14",
      "light": "nord14"
    },
    "markdownBlockQuote": {
      "dark": "nord3",
      "light": "nord3"
    },
    "markdownEmph": {
      "dark": "nord12",
      "light": "nord12"
    },
    "markdownStrong": {
      "dark": "nord13",
      "light": "nord13"
    },
    "markdownHorizontalRule": {
      "dark": "nord3",
      "light": "nord3"
    },
    "markdownListItem": {
      "dark": "nord8",
      "light": "nord10"
    },
    "markdownListEnumeration": {
      "dark": "nord7",
      "light": "nord7"
    },
    "markdownImage": {
      "dark": "nord9",
      "light": "nord9"
    },
    "markdownImageText": {
      "dark": "nord7",
      "light": "nord7"
    },
    "markdownCodeBlock": {
      "dark": "nord4",
      "light": "nord0"
    },
    "syntaxComment": {
      "dark": "nord3",
      "light": "nord3"
    },
    "syntaxKeyword": {
      "dark": "nord9",
      "light": "nord9"
    },
    "syntaxFunction": {
      "dark": "nord8",
      "light": "nord8"
    },
    "syntaxVariable": {
      "dark": "nord7",
      "light": "nord7"
    },
    "syntaxString": {
      "dark": "nord14",
      "light": "nord14"
    },
    "syntaxNumber": {
      "dark": "nord15",
      "light": "nord15"
    },
    "syntaxType": {
      "dark": "nord7",
      "light": "nord7"
    },
    "syntaxOperator": {
      "dark": "nord9",
      "light": "nord9"
    },
    "syntaxPunctuation": {
      "dark": "nord4",
      "light": "nord0"
    }
  }
}
```

--------------------------------

### Sessions APIs

Source: https://opencode.ai/docs/server

Comprehensive endpoints for managing sessions, including creation, retrieval, messaging, and operations like aborting or summarizing. Sessions represent chat or interaction contexts.

```APIDOC
## GET /session

### Description
List sessions.

### Method
GET

### Endpoint
/session

### Parameters
No parameters.

### Response
#### Success Response (200)
- Array of Session objects.

## GET /session/:id

### Description
Get session by ID.

### Method
GET

### Endpoint
/session/{id}

### Parameters
#### Path Parameters
- **id** (string) - Required - Session identifier.

### Response
#### Success Response (200)
- Session object.

## GET /session/:id/children

### Description
List child sessions.

### Method
GET

### Endpoint
/session/{id}/children

### Parameters
#### Path Parameters
- **id** (string) - Required - Parent session ID.

### Response
#### Success Response (200)
- Array of Session objects.

## POST /session

### Description
Create a new session.

### Method
POST

### Endpoint
/session

### Parameters
No path or query parameters.

### Request Body
- **parentID** (string) - Optional - Parent session ID.
- **title** (string) - Optional - Session title.

### Response
#### Success Response (201)
- Created Session object.

## DELETE /session/:id

### Description
Delete a session.

### Method
DELETE

### Endpoint
/session/{id}

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID to delete.

### Response
#### Success Response (200)
- No content.

## PATCH /session/:id

### Description
Update session properties.

### Method
PATCH

### Endpoint
/session/{id}

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.

### Request Body
- **title** (string) - Optional - New title.

### Response
#### Success Response (200)
- Updated Session object.

## POST /session/:id/init

### Description
Analyze app and create AGENTS.md.

### Method
POST

### Endpoint
/session/{id}/init

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.

### Request Body
- **messageID** (string) - Required - Message identifier.
- **providerID** (string) - Required - Provider ID.
- **modelID** (string) - Required - Model ID.

### Response
#### Success Response (200)
- Initialization result.

## POST /session/:id/abort

### Description
Abort a running session.

### Method
POST

### Endpoint
/session/{id}/abort

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.

### Response
#### Success Response (200)
- No content.

## POST /session/:id/share

### Description
Share a session.

### Method
POST

### Endpoint
/session/{id}/share

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.

### Response
#### Success Response (200)
- Shared Session object.

## DELETE /session/:id/share

### Description
Unshare a session.

### Method
DELETE

### Endpoint
/session/{id}/share

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.

### Response
#### Success Response (200)
- Unshared Session object.

## POST /session/:id/summarize

### Description
Summarize a session.

### Method
POST

### Endpoint
/session/{id}/summarize

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.

### Response
#### Success Response (200)
- Summary result.

## GET /session/:id/message

### Description
List messages in a session.

### Method
GET

### Endpoint
/session/{id}/message

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.

### Response
#### Success Response (200)
- Array of message objects with info and parts.

## GET /session/:id/message/:messageID

### Description
Get message details.

### Method
GET

### Endpoint
/session/{id}/message/{messageID}

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.
- **messageID** (string) - Required - Message ID.

### Response
#### Success Response (200)
- Message object with info and parts.

## POST /session/:id/message

### Description
Send a chat message to the session.

### Method
POST

### Endpoint
/session/{id}/message

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.

### Request Body
- Matches ChatInput structure. Optional noReply: true to skip AI inference.

### Response
#### Success Response (201)
- Created Message object.

## POST /session/:id/shell

### Description
Run a shell command in the session.

### Method
POST

### Endpoint
/session/{id}/shell

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.

### Request Body
- Matches CommandInput structure.

### Response
#### Success Response (201)
- Message object from command.

## POST /session/:id/revert

### Description
Revert a message.

### Method
POST

### Endpoint
/session/{id}/revert

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.

### Request Body
- **messageID** (string) - Required - ID of message to revert.

### Response
#### Success Response (200)
- Revert result.

## POST /session/:id/unrevert

### Description
Restore reverted messages.

### Method
POST

### Endpoint
/session/{id}/unrevert

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.

### Response
#### Success Response (200)
- Restore result.

## POST /session/:id/permissions/:permissionID

### Description
Respond to a permission request.

### Method
POST

### Endpoint
/session/{id}/permissions/{permissionID}

### Parameters
#### Path Parameters
- **id** (string) - Required - Session ID.
- **permissionID** (string) - Required - Permission ID.

### Request Body
- **response** (string) - Required - Response to permission.

### Response
#### Success Response (200)
- Permission response result.
```

--------------------------------

### Set OpenCode theme (JSON)

Source: https://opencode.ai/docs/config

Shows a simple theme setting in the OpenCode config. Intended to switch the UI theme by updating the theme property.

```json

```

--------------------------------

### Configure Global Tools

Source: https://opencode.ai/docs/tools

Demonstrates how to configure tools globally using the `tools` option in the configuration file.  This sets default tool states for the entire project.  It uses JSON format to configure tool enabling/disabling.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "write": false,
    "bash": false,
    "webfetch": true
  }
}
```

--------------------------------

### Configure remote MCP server

Source: https://opencode.ai/docs/mcp-servers

Demonstrates how to configure a remote MCP server with URL and authentication headers. Includes required fields like type and url, and optional fields like headers and timeout.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "my-remote-mcp": {
      "type": "remote",
      "url": "https://my-mcp-server.com",
      "enabled": true,
      "headers": {
        "Authorization": "Bearer MY_API_KEY"
      }
    }
  }
}
```

--------------------------------

### Documentation API

Source: https://opencode.ai/docs/server

Endpoint to retrieve the OpenAPI 3.1 specification for the API.

```APIDOC
## GET /doc

### Description
Retrieves the OpenAPI 3.1 specification for this API.

### Method
GET

### Endpoint
/doc

### Parameters
None

### Response
#### Success Response (200)
- **openapiSpec** (HTML) - An HTML page displaying the OpenAPI 3.1 specification.

### Response Example
(This endpoint returns an HTML page, the content would be the rendered OpenAPI documentation.)
```

--------------------------------

### Configure .npmrc for Private Registry (Config)

Source: https://opencode.ai/docs/enterprise

This configuration defines registry and authentication details for the private NPM registry in a .npmrc file. It uses environment variables for security and requires Bun or NPM support. It enables package fetching; developers must set NPM_AUTH_TOKEN and place the file in the project root.

```ini
registry=https://your-company.jfrog.io/api/npm/npm-virtual/
//your-company.jfrog.io/api/npm/npm-virtual/:_authToken=${NPM_AUTH_TOKEN}
```

--------------------------------

### Path API

Source: https://opencode.ai/docs/sdk

The Path API provides a method for retrieving information about the current path.

```APIDOC
## Path

### Description
The Path API provides a method for retrieving information about the current path.

### Method
APIDOC

### Endpoint
N/A - This is a client-side method.

### Parameters
#### Path Parameters
- N/A

#### Query Parameters
- N/A

#### Request Body
- N/A

### Get Current Path
#### Description
Retrieves the current path information.
#### Response
#### Success Response (200)
- **pathInfo** (Path) - Details of the current path.
#### Response Example
{
  "id": "path_id",
  "name": "Path Name"
}
```

--------------------------------

### Authentication API

Source: https://opencode.ai/docs/server

Endpoint for setting authentication credentials for a specific provider.

```APIDOC
## PUT /auth/:id

### Description
Sets the authentication credentials for a given provider ID. The request body must conform to the specific provider's schema.

### Method
PUT

### Endpoint
/auth/:id

### Parameters
#### Path Parameters
- **id** (string) - Required - The unique identifier for the authentication provider.

#### Request Body
- **credentials** (object) - Required - An object containing the authentication credentials, matching the provider's schema.

### Response
#### Success Response (200)
- **success** (boolean) - Indicates if the operation was successful.

### Request Example
```json
{
  "credentials": {
    "apiKey": "your_api_key_here"
  }
}
```

### Response Example
```json
{
  "success": true
}
```
```

--------------------------------

### GitHub Workflow for opencode

Source: https://opencode.ai/docs/github

This YAML file defines a GitHub Actions workflow that triggers opencode based on issue comments containing `/oc` or `/opencode`. It fetches the repository, runs the opencode action, and uses environment variables for API keys and model selection.

```YAML
name: opencode

on:
  issue_comment: 
    types: [created]

jobs:
  opencode:
    if: | 
      contains(github.event.comment.body, '/oc') ||
      contains(github.event.comment.body, '/opencode')
    runs-on: ubuntu-latest
    permissions:
      id-token: write
    steps:
      - name: Checkout repository
        uses: actions/checkout@v4
        with:
          fetch-depth: 1

      - name: Run opencode
        uses: sst/opencode/github@latest
        env:
          ANTHROPIC_API_KEY: ${{ secrets.ANTHROPIC_API_KEY }}
        with:
          model: anthropic/claude-sonnet-4-20250514
          # share: true
          # github_token: xxxx
```

--------------------------------

### Configure Global Permissions in JSON

Source: https://opencode.ai/docs/permissions

This JSON configuration sets global permissions for the edit, bash, and webfetch tools in OpenCode. It uses the 'permission' object to specify 'allow', 'ask', or 'deny' for each tool, requiring no external dependencies. Inputs are the permission levels, and outputs control tool behavior without limitations on scope.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "edit": "allow",
    "bash": "ask",
    "webfetch": "deny"
  }
}
```

--------------------------------

### Configure Permissions for OpenCode Agents

Source: https://opencode.ai/docs/agents

Permissions manage agent actions for tools like edit, bash, and webfetch, with options ask, allow, or deny. Can be set globally or per agent, including specific bash commands or glob patterns. Overrides allow finer control, and Markdown agents support permission configs.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "edit": "deny"
  }
}
```

```json
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "edit": "deny"
  },
  "agent": {
    "build": {
      "permission": {
        "edit": "ask"
      }
    }
  }
}
```

```markdown
---
description: Code review without edits
mode: subagent
permission:
  edit: deny
  bash:
    "git diff": allow
    "git log*": allow
    "*": ask
  webfetch: deny
---

Only analyze code and suggest changes.

```

```json
{
  "$schema": "https://opencode.ai/config.json",
  "agent": {
    "build": {
      "permission": {
        "bash": {
          "git push": "ask"
        }
      }
    }
  }
}
```

```json
{
  "$schema": "https://opencode.ai/config.json",
  "agent": {
    "build": {
      "permission": {
        "bash": {
          "git *": "ask"
        }
      }
    }
  }
}
```

```json
{
  "$schema": "https://opencode.ai/config.json",
  "agent": {
    "build": {
      "permission": {
        "bash": {
          "git status": "allow",
          "*": "ask"
        }
      }
    }
  }
}
```

--------------------------------

### Configure Edit Tool

Source: https://opencode.ai/docs/tools

Demonstrates enabling the `edit` tool for making precise string replacements in files. This is the tool primarily used for code modifications. Configuration is done in a JSON file.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "edit": true
  }
}
```

--------------------------------

### Configure Agent-Specific Tools

Source: https://opencode.ai/docs/tools

Illustrates overriding global tool settings for specific agents in their definitions. Agent-specific configurations take precedence over global settings. Uses JSON format.

```JSON
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "write": true,
    "bash": true
  },
  "agent": {
    "plan": {
      "tools": {
        "write": false,
        "bash": false
      }
    }
  }
}
```

--------------------------------

### Set Bash Tool Permissions Globally in JSON

Source: https://opencode.ai/docs/permissions

This configuration sets permissions for all bash commands in OpenCode by setting 'bash' to 'ask'. It ensures user approval is required for any bash operation, with no specific dependencies. The input is the permission level, controlling bash tool access comprehensively.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "bash": "ask"
  }
}
```

--------------------------------

### Reference files in command prompt

Source: https://opencode.ai/docs/commands

Shows how to include file contents in custom command prompts using the @filename syntax. The referenced file content is automatically included in the prompt when the command is executed.

```markdown
---
description: Review component
---


Review the component in @src/components/Button.tsx.
Check for performance issues and suggest improvements.

```

--------------------------------

### Configure TUI scroll speed (JSON)

Source: https://opencode.ai/docs/config

Configures TUI settings, specifically the scroll speed. Uses a JSON structure with a $schema reference. No external dependencies besides the OpenCode config schema.

```json
{\n  \"$schema\": \"https://opencode.ai/config.json\",\n  \"tui\": {\n    \"scroll_speed\": 3\n  }\n}
```

--------------------------------

### Control Edit Tool Permissions in JSON

Source: https://opencode.ai/docs/permissions

This JSON snippet configures permissions specifically for the edit tool in OpenCode. It sets the 'edit' key under 'permission' to 'ask', prompting for user approval before file edits. No dependencies are needed, and it takes the permission value as input to enforce restrictions on editing operations.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "edit": "ask"
  }
}
```

--------------------------------

### Create Tool with Plain Object

Source: https://opencode.ai/docs/custom-tools

Alternative tool definition using plain object structure with Zod imported directly. Provides same functionality as tool() helper but with more explicit control over schema definition and execution.

```typescript
import { z } from "zod"

export default {
  description: "Tool description",
  args: {
    param: z.string().describe("Parameter description"),
  },
  async execute(args, context) {
    // Tool implementation
    return "result"
  },
}
```

```javascript
import { z } from "zod"

export default {
  description: "Tool description",
  args: {
    param: z.string().describe("Parameter description"),
  },
  async execute(args, context) {
    // Tool implementation
    return "result"
  },
}
```

--------------------------------

### Configure Granular Bash Command Permissions in JSON

Source: https://opencode.ai/docs/permissions

This JSON allows fine-tuned control over specific bash commands in OpenCode by mapping command strings to permission levels. It uses a 'bash' object for commands like 'git push' as 'ask', enabling selective enforcement without external tools. Inputs include command names and levels, structuring access to terminal operations.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "bash": {
      "git push": "ask",
      "git status": "allow",
      "git diff": "allow",
      "npm run build": "allow",
      "ls": "allow",
      "pwd": "allow"
    }
  }
}
```

--------------------------------

### Access Tool Context Information

Source: https://opencode.ai/docs/custom-tools

Demonstrates how tools receive context about the current session including agent, sessionID, and messageID. Context provides session state information that tools can use for enhanced functionality.

```typescript
import { tool } from "@opencode-ai/plugin"

export default tool({
  description: "Get project information",
  args: {},
  async execute(args, context) {
    // Access context information
    const { agent, sessionID, messageID } = context
    return `Agent: ${agent}, Session: ${sessionID}, Message: ${messageID}`
  },
})
```

```javascript
import { tool } from "@opencode-ai/plugin"

export default tool({
  description: "Get project information",
  args: {},
  async execute(args, context) {
    // Access context information
    const { agent, sessionID, messageID } = context
    return `Agent: ${agent}, Session: ${sessionID}, Message: ${messageID}`
  },
})
```

--------------------------------

### Set Custom Prompt for Agent in JSON

Source: https://opencode.ai/docs/agents

This config specifies a custom system prompt file for the agent, relative to the config file location. The prompt file contains instructions tailored to the agent's purpose. Applies to global or project-specific configs.

```json
{
  "agent": {
    "review": {
      "prompt": "{file:./prompts/code-review.txt}"
    }
  }
}
```

--------------------------------

### Handle SDK Errors

Source: https://opencode.ai/docs/sdk

Implement error handling for SDK operations using a try-catch block. This allows you to gracefully manage potential issues, such as attempting to access a non-existent session.

```typescript
try {
  await client.session.get({ path: { id: "invalid-id" } })
} catch (error) {
  console.error("Failed to get session:", (error as Error).message)
}
```

--------------------------------

### Configure Agent Permissions in YAML

Source: https://opencode.ai/docs/permissions

This YAML frontmatter defines permissions for an agent in Markdown, denying edit, asking for bash, and denying webfetch. It sets a mode and description, controlling tool access without JSON. Input is the YAML structure, outputting restricted agent capabilities for code review.

```yaml
---
description: Code review without edits
mode: subagent
permission:
  edit: deny
  bash: ask
  webfetch: deny
---
```

--------------------------------

### Files APIs

Source: https://opencode.ai/docs/server

Endpoints for searching and managing files in the workspace, including text search, file finding, and status retrieval.

```APIDOC
## GET /find?pattern=<pat>

### Description
Search for text in files using a pattern.

### Method
GET

### Endpoint
/find

### Parameters
#### Query Parameters
- **pattern** (string) - Required - Search pattern.

### Response
#### Success Response (200)
- Array of match objects with path, lines, line_number, absolute_offset, submatches.

## GET /find/file?query=<q>

### Description
Find files by name.

### Method
GET

### Endpoint
/find/file

### Parameters
#### Query Parameters
- **query** (string) - Required - File name query.

### Response
#### Success Response (200)
- Array of file paths (string[]).

## GET /find/symbol?query=<q>

### Description
Find workspace symbols.

### Method
GET

### Endpoint
/find/symbol

### Parameters
#### Query Parameters
- **query** (string) - Required - Symbol query.

### Response
#### Success Response (200)
- Array of Symbol objects.

## GET /file?path=<path>

### Description
Read a file content.

### Method
GET

### Endpoint
/file

### Parameters
#### Query Parameters
- **path** (string) - Required - File path.

### Response
#### Success Response (200)
- Object with type (raw or patch) and content string.

## GET /file/status

### Description
Get status for tracked files.

### Method
GET

### Endpoint
/file/status

### Parameters
No parameters.

### Response
#### Success Response (200)
- Array of File objects.
```

--------------------------------

### Configure Code Formatters - JSON

Source: https://opencode.ai/docs/config

Sets up code formatting tools including built-in and custom formatters. Supports disabling default formatters like Prettier and configuring custom formatters with specific commands, environments, and file extensions.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "formatter": {
    "prettier": {
      "disabled": true
    },
    "custom-prettier": {
      "command": ["npx", "prettier", "--write", "$FILE"],
      "environment": {
        "NODE_ENV": "development"
      },
      "extensions": [".js", ".ts", ".jsx", ".tsx"]
    }
  }
}
```

--------------------------------

### Logging APIs

Source: https://opencode.ai/docs/server

Endpoint for writing log entries to the server. Used for logging application events programmatically.

```APIDOC
## POST /log

### Description
Write a log entry.

### Method
POST

### Endpoint
/log

### Parameters
No parameters.

### Request Body
- **service** (string) - Required - Logging service.
- **level** (string) - Required - Log level.
- **message** (string) - Required - Log message.
- **extra** (object) - Optional - Additional data.

### Response
#### Success Response (200)
- boolean indicating success.
```

--------------------------------

### Execute Shell Command in TUI - Shell

Source: https://opencode.ai/docs/tui

Prefix messages with ! to run shell commands directly in the TUI. The output of the command is added to the conversation. Dependencies include a shell environment.

```shell
!ls -la
```

--------------------------------

### Import TypeScript Definitions

Source: https://opencode.ai/docs/sdk

Import TypeScript type definitions for various opencode API elements like Session, Message, and Part. These types are generated from the server's OpenAPI specification.

```typescript
import type { Session, Message, Part } from "@opencode-ai/sdk"
```

--------------------------------

### TUI API

Source: https://opencode.ai/docs/sdk

Provides text user interface controls for managing the application's interactive terminal interface.

```APIDOC
## TUI API

### Description
The TUI (Text User Interface) API provides methods for controlling the application's terminal interface, including prompt management, dialogs, and command execution.

### Methods
- **POST** `tui.appendPrompt({ body })` - Append text to the prompt
- **GET** `tui.openHelp()` - Open the help dialog
- **GET** `tui.openSessions()` - Open the session selector
- **GET** `tui.openThemes()` - Open the theme selector
- **GET** `tui.openModels()` - Open the model selector
- **POST** `tui.submitPrompt()` - Submit the current prompt
- **GET** `tui.clearPrompt()` - Clear the prompt
- **POST** `tui.executeCommand({ body })` - Execute a command
- **POST** `tui.showToast({ body })` - Show toast notification

### Parameters
#### Path Parameters
- **text** (string) - Required - Text to append to prompt
- **command** (string) - Required - Command to execute
- **message** (string) - Required - Toast message content
- **variant** (string) - Optional - Toast variant (success, error, info, warning)

### Request Example
{
  "body": {
    "text": "Add this to prompt"
  }
}

### Response
#### Success Response (200)
- **boolean** - Operation success status

#### Response Example
{
  "success": true
}
```

--------------------------------

### Events API

Source: https://opencode.ai/docs/sdk

Provides server-sent events streaming for real-time updates and notifications from the application.

```APIDOC
## Events API

### Description
The Events API provides server-sent events streaming for receiving real-time updates, notifications, and system events from the OpenCode AI application.

### Methods
- **GET** `event.subscribe()` - Subscribe to server-sent events stream

### Parameters
None - This is a streaming endpoint that doesn't require parameters.

### Request Example
No request body required - establishes a persistent connection for event streaming.

### Response
#### Success Response (200)
- **Server-Sent Events Stream** (stream) - Continuous stream of event objects

#### Response Example
{
  "event": {
    "type": "session.created",
    "properties": {
      "id": "session-123",
      "title": "New Session",
      "timestamp": "2024-01-01T12:00:00Z"
    }
  }
}
```

--------------------------------

### Bind Keyboard Shortcut in Zed for OpenCode ACP

Source: https://opencode.ai/docs/acp

This JSON snippet defines a keyboard shortcut in Zed to open a new thread for the OpenCode agent. This facilitates quick access to OpenCode's AI coding assistance features within the editor. The shortcut 'cmd-alt-o' is mapped to the 'agent::NewExternalAgentThread' action.

```json
[
  {
    "bindings": {
      "cmd-alt-o": ["agent::NewExternalAgentThread", { "agent": "OpenCode" }]
    }
  }
]
```

--------------------------------

### Configure theme in OpenCode config

Source: https://opencode.ai/docs/themes

JSON configuration snippet to set the theme in OpenCode. Replace 'tokyonight' with any built-in or custom theme name.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "theme": "tokyonight"
}
```

--------------------------------

### POST /zen/v1/messages

Source: https://opencode.ai/docs/zen

Endpoint for Claude models including Claude Sonnet 4.5, Claude Sonnet 4, Claude Haiku 4.5, Claude Haiku 3.5, and Claude Opus 4.1. This endpoint handles chat completion requests for Anthropic models.

```APIDOC
## POST /zen/v1/messages

### Description
Endpoint for Claude models including Claude Sonnet 4.5, Claude Sonnet 4, Claude Haiku 4.5, Claude Haiku 3.5, and Claude Opus 4.1. This endpoint handles chat completion requests for Anthropic models.

### Method
POST

### Endpoint
https://opencode.ai/zen/v1/messages

### Parameters
#### Path Parameters
None

#### Query Parameters
None

#### Request Body
- **model** (string) - Required - The model identifier (e.g., claude-sonnet-4-5, claude-opus-4-1)
- **messages** (array) - Required - Array of message objects
- **max_tokens** (integer) - Required - Maximum number of tokens to generate
- **stream** (boolean) - Optional - Whether to stream the response

### Request Example
{
  "model": "claude-sonnet-4-5",
  "messages": [
    {
      "role": "user",
      "content": "Explain how neural networks work"
    }
  ],
  "max_tokens": 1000,
  "stream": false
}

### Response
#### Success Response (200)
- **id** (string) - Unique identifier for the response
- **content** (array) - Array of content blocks
- **usage** (object) - Token usage statistics

#### Response Example
{
  "id": "msg_123",
  "content": [
    {
      "type": "text",
      "text": "Neural networks are computational models inspired by..."
    }
  ],
  "usage": {
    "input_tokens": 30,
    "output_tokens": 850
  }
}
```

--------------------------------

### Check terminal truecolor support

Source: https://opencode.ai/docs/themes

Command to check if your terminal supports truecolor (24-bit color). Output should be 'truecolor' or '24bit' for full theme color accuracy.

```shell
echo $COLORTERM
```

--------------------------------

### Configure Global Model Options in JSON

Source: https://opencode.ai/docs/models

This JSON config globally sets options for specific models via their providers, such as reasoning effort for OpenAI and thinking budget for Anthropic. It depends on having the providers configured and uses the $schema for structure. Options apply globally unless overridden by agent configs.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "openai": {
      "models": {
        "gpt-5": {
          "options": {
            "reasoningEffort": "high",
            "textVerbosity": "low",
            "reasoningSummary": "auto",
            "include": ["reasoning.encrypted_content"],
          },
        },
      },
    },
    "anthropic": {
      "models": {
        "claude-sonnet-4-5-20250929": {
          "options": {
            "thinking": {
              "type": "enabled",
              "budgetTokens": 16000,
            },
          },
        },
      },
    },
  },
}
```

--------------------------------

### Enable truecolor in terminal

Source: https://opencode.ai/docs/themes

Command to enable truecolor support in your terminal by setting the COLORTERM environment variable. Add this to your shell profile for persistence.

```shell
export COLORTERM=truecolor
```

--------------------------------

### Configure Agent Description in JSON

Source: https://opencode.ai/docs/agents

This configuration option sets a description for an agent in OpenCode, outlining its purpose and usage. It requires the description to be provided as a string in the config. This is a required field to define agent functionality clearly.

```json
{
  "agent": {
    "review": {
      "description": "Reviews code for best practices and potential issues"
    }
  }
}
```

--------------------------------

### GitHub Permissions for opencode

Source: https://opencode.ai/docs/github

This configuration specifies the required permissions for the opencode GitHub Action to perform operations like creating comments, committing changes, and opening pull requests. It uses the `id-token` capability to enable authenticated actions.

```YAML
permissions:
  id-token: write
  contents: write
  pull-requests: write
  issues: write
```

--------------------------------

### Control Agent Tools in JSON

Source: https://opencode.ai/docs/agents

This config enables or disables specific tools for agents, using true or false values. Wildcards can control multiple tools, such as disabling all from an MCP server. Tools like write, bash, or edit can be managed per agent.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "write": true,
    "bash": true
  },
  "agent": {
    "plan": {
      "tools": {
        "write": false,
        "bash": false
      }
    }
  }
}
```

```json
{
  "$schema": "https://opencode.ai/config.json",
  "agent": {
    "readonly": {
      "tools": {
        "mymcp_*": false,
        "write": false,
        "edit": false
      }
    }
  }
}
```

--------------------------------

### Specify Agent for Command in JSON

Source: https://opencode.ai/docs/commands

This configuration allows specifying which agent executes a command, triggering a subagent by default. It's optional and defaults to the current agent. Use in projects like OpenCode AI where agent roles are defined.

```json
{
  "command": {
    "review": {
      "agent": "plan"
    }
  }
}
```

--------------------------------

### Agents APIs

Source: https://opencode.ai/docs/server

Endpoint for listing available agents in the opencode system. Agents represent AI or tool capabilities.

```APIDOC
## GET /agent

### Description
List all available agents.

### Method
GET

### Endpoint
/agent

### Parameters
No parameters.

### Response
#### Success Response (200)
- Array of Agent objects.
```

--------------------------------

### TUI API Endpoints

Source: https://opencode.ai/docs/server

Endpoints for interacting with the Text User Interface (TUI) of Opencode AI, including prompt management, dialogs, and command execution.

```APIDOC
## POST /tui/append-prompt

### Description
Appends text to the current prompt in the TUI.

### Method
POST

### Endpoint
/tui/append-prompt

### Parameters
#### Request Body
- **text** (string) - Required - The text to append to the prompt.

### Response
#### Success Response (200)
- **success** (boolean) - Indicates if the operation was successful.

### Request Example
```json
{
  "text": "This is the text to append."
}
```

### Response Example
```json
{
  "success": true
}
```
```

```APIDOC
## POST /tui/open-help

### Description
Opens the help dialog in the TUI.

### Method
POST

### Endpoint
/tui/open-help

### Parameters
None

### Response
#### Success Response (200)
- **success** (boolean) - Indicates if the operation was successful.

### Response Example
```json
{
  "success": true
}
```
```

```APIDOC
## POST /tui/open-sessions

### Description
Opens the session selector dialog in the TUI.

### Method
POST

### Endpoint
/tui/open-sessions

### Parameters
None

### Response
#### Success Response (200)
- **success** (boolean) - Indicates if the operation was successful.

### Response Example
```json
{
  "success": true
}
```
```

```APIDOC
## POST /tui/open-themes

### Description
Opens the theme selector dialog in the TUI.

### Method
POST

### Endpoint
/tui/open-themes

### Parameters
None

### Response
#### Success Response (200)
- **success** (boolean) - Indicates if the operation was successful.

### Response Example
```json
{
  "success": true
}
```
```

```APIDOC
## POST /tui/open-models

### Description
Opens the model selector dialog in the TUI.

### Method
POST

### Endpoint
/tui/open-models

### Parameters
None

### Response
#### Success Response (200)
- **success** (boolean) - Indicates if the operation was successful.

### Response Example
```json
{
  "success": true
}
```
```

```APIDOC
## POST /tui/submit-prompt

### Description
Submits the current prompt in the TUI.

### Method
POST

### Endpoint
/tui/submit-prompt

### Parameters
None

### Response
#### Success Response (200)
- **success** (boolean) - Indicates if the operation was successful.

### Response Example
```json
{
  "success": true
}
```
```

```APIDOC
## POST /tui/clear-prompt

### Description
Clears the current prompt in the TUI.

### Method
POST

### Endpoint
/tui/clear-prompt

### Parameters
None

### Response
#### Success Response (200)
- **success** (boolean) - Indicates if the operation was successful.

### Response Example
```json
{
  "success": true
}
```
```

```APIDOC
## POST /tui/execute-command

### Description
Executes a specific command within the TUI.

### Method
POST

### Endpoint
/tui/execute-command

### Parameters
#### Request Body
- **command** (string) - Required - The command to execute.

### Response
#### Success Response (200)
- **success** (boolean) - Indicates if the operation was successful.

### Request Example
```json
{
  "command": "git status"
}
```

### Response Example
```json
{
  "success": true
}
```
```

```APIDOC
## POST /tui/show-toast

### Description
Displays a toast notification in the TUI.

### Method
POST

### Endpoint
/tui/show-toast

### Parameters
#### Request Body
- **title** (string) - Optional - The title of the toast notification.
- **message** (string) - Required - The main message of the toast notification.
- **variant** (string) - Optional - The variant or type of the toast (e.g., 'info', 'success', 'warning', 'error').

### Response
#### Success Response (200)
- **success** (boolean) - Indicates if the operation was successful.

### Request Example
```json
{
  "title": "Update Complete",
  "message": "Your changes have been saved.",
  "variant": "success"
}
```

### Response Example
```json
{
  "success": true
}
```
```

```APIDOC
## GET /tui/control/next

### Description
Waits for and retrieves the next control request from the TUI.

### Method
GET

### Endpoint
/tui/control/next

### Parameters
None

### Response
#### Success Response (200)
- **controlRequest** (object) - An object representing the control request.

### Response Example
```json
{
  "type": "input",
  "details": {
    "prompt": "Enter your name:"
  }
}
```
```

```APIDOC
## POST /tui/control/response

### Description
Sends a response to a pending TUI control request.

### Method
POST

### Endpoint
/tui/control/response

### Parameters
#### Request Body
- **body** (object) - Required - The response data for the control request.

### Response
#### Success Response (200)
- **success** (boolean) - Indicates if the operation was successful.

### Request Example
```json
{
  "body": {
    "answer": "John Doe"
  }
}
```

### Response Example
```json
{
  "success": true
}
```
```

--------------------------------

### Configure MCP server per agent

Source: https://opencode.ai/docs/mcp-servers

Shows how to enable an MCP server for a specific agent while disabling it globally. This approach is useful when managing multiple MCP servers. The configuration disables the tool globally and then enables it for a specific agent.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "my-mcp": {
      "type": "local",
      "command": ["bun", "x", "my-mcp-command"],
      "enabled": true
    }
  },
  "tools": {
    "my-mcp*": false
  },
  "agent": {
    "my-agent": {
      "tools": {
        "my-mcp*": true
      }
    }
  }
}
```

--------------------------------

### Events API

Source: https://opencode.ai/docs/server

Endpoint for accessing a stream of server-sent events.

```APIDOC
## GET /event

### Description
Provides a server-sent events (SSE) stream. The first event received will be `server.connected`, followed by subsequent bus events.

### Method
GET

### Endpoint
/event

### Parameters
None

### Response
#### Success Response (200)
- **eventStream** (SSE) - A stream of server-sent events.

### Response Example
```sse
Event: server.connected

Data: {"message": "Connected to the event stream"}

Event: some.bus.event

Data: {"payload": "event data"}
```
```

--------------------------------

### Define Custom Models in JSON Config

Source: https://opencode.ai/docs/models

This JSON defines custom models that extend built-in ones with specific options, like high or low reasoning for GPT-5 variants. It requires the base model to be available and uses the $schema for config validation. Custom models allow tailoring options without altering global settings.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "opencode": {
      "models": {
        "gpt-5-high": {
          "id": "gpt-5",
          "options": {
            "reasoningEffort": "high",
            "textVerbosity": "low",
            "reasoningSummary": "auto",
          },
        },
        "gpt-5-low": {
          "id": "gpt-5",
          "options": {
            "reasoningEffort": "low",
            "textVerbosity": "low",
            "reasoningSummary": "auto",
          },
        },
      },
    },
  },
}
```

--------------------------------

### Use Wildcards to Deny Terraform Commands in JSON

Source: https://opencode.ai/docs/permissions

This JSON config employs wildcards to deny all Terraform commands in OpenCode by setting 'terraform *' to 'deny'. It leverages glob patterns for pattern matching, requiring no additional dependencies. The wildcard acts as input to block specific command families effectively.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "bash": {
      "terraform *": "deny"
    }
  }
}
```

--------------------------------

### Configure agent mode in JSON

Source: https://opencode.ai/docs/agents

Sets the operational mode for an agent using the mode configuration option. The mode can be set to primary, subagent, or all. This determines how the agent can be used within the system. If not specified, the mode defaults to all.

```json
{
  "agent": {
    "review": {
      "mode": "subagent"
    }
  }
}
```

--------------------------------

### Disable specific OpenCode tools (JSON)

Source: https://opencode.ai/docs/config

Shows how to disable certain OpenCode tools via the tools section in the config. Inputs are boolean flags under the tools object; outputs are applied tool availability at runtime.

```json
{\n  \"$schema\": \"https://opencode.ai/config.json\",\n  \"tools\": {\n    \"write\": false,\n    \"bash\": false\n  }\n}
```

--------------------------------

### Override Model for Command in JSON

Source: https://opencode.ai/docs/commands

This config overrides the default model for a specific command execution. It's optional and allows use of models like Anthropic's Claude. Depends on available AI models in the system.

```json
{
  "command": {
    "analyze": {
      "model": "anthropic/claude-3-5-sonnet-20241022"
    }
  }
}
```

--------------------------------

### Set Default Model in JSON Config

Source: https://opencode.ai/docs/models

This JSON configuration sets a default model for OpenCode by specifying the provider and model ID. The full ID format is provider_id/model_id, and it uses the $schema for validation. It requires no dependencies but ensures the model is available in the configured provider.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "model": "lmstudio/google/gemma-3n-e4b"
}
```

--------------------------------

### Configure Agent Temperature in JSON

Source: https://opencode.ai/docs/agents

This option controls the randomness of LLM responses for agents, with lower values for determinism and higher for creativity. Values range from 0.0 to 1.0, suitable for tasks like analysis or brainstorming. Defaults are model-specific if not specified.

```json
{
  "agent": {
    "plan": {
      "temperature": 0.1
    },
    "creative": {
      "temperature": 0.8
    }
  }
}
```

```json
{
  "agent": {
    "analyze": {
      "temperature": 0.1,
      "prompt": "{file:./prompts/analysis.txt}"
    },
    "build": {
      "temperature": 0.3
    },
    "brainstorm": {
      "temperature": 0.7,
      "prompt": "{file:./prompts/creative.txt}"
    }
  }
}
```

--------------------------------

### Configure Disabled Providers - JSON

Source: https://opencode.ai/docs/config

Disables specific AI providers from loading regardless of available credentials. Useful for preventing certain providers from appearing in model selection lists or being used even when API keys are configured.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "disabled_providers": ["openai", "gemini"]
}
```

--------------------------------

### OpenCode Configuration - Disabling TypeScript LSP Server

Source: https://opencode.ai/docs/lsp

This configuration snippet shows how to disable the built-in TypeScript LSP server. By setting `disabled` to `true` for a specific LSP server, OpenCode will not attempt to use it, even if matching file extensions are detected.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "lsp": {
    "typescript": {
      "disabled": true
    }
  }
}
```

--------------------------------

### Set EDITOR Environment Variable for OpenCode

Source: https://opencode.ai/docs/tui

Configure the EDITOR environment variable to specify the external editor for /editor and /export commands. To make permanent, add to shell profile, system properties, or PowerShell profile. Dependencies include the chosen editor.

```bash
export EDITOR=nano
export EDITOR=vim
export EDITOR="code --wait"
```

```cmd
set EDITOR=notepad
set EDITOR=code --wait
```

```powershell
$env:EDITOR = "notepad"
$env:EDITOR = "code --wait"
```

--------------------------------

### Globally disable MCP tools

Source: https://opencode.ai/docs/mcp-servers

Shows how to globally disable specific MCP tools or use glob patterns to disable multiple matching MCP tools. This configuration affects all MCP tools in the system.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "my-mcp-foo": {
      "type": "local",
      "command": ["bun", "x", "my-mcp-command-foo"]
    },
    "my-mcp-bar": {
      "type": "local",
      "command": ["bun", "x", "my-mcp-command-bar"]
    }
  },
  "tools": {
    "my-mcp*": false
  }
}
```

--------------------------------

### Disable Agent in JSON

Source: https://opencode.ai/docs/agents

Setting disable to true deactivates the agent in OpenCode configurations. This prevents the agent from executing or being used in tasks. Useful for temporarily suspending agent functionality.

```json
{
  "agent": {
    "review": {
      "disable": true
    }
  }
}
```

=== COMPLETE CONTENT === This response contains all available snippets from this library. No additional content exists. Do not make further requests.