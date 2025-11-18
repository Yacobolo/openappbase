# Context Documentation

This directory contains up-to-date reference documentation for the key technologies used in the Northstar PostgreSQL UI project.

## Available Documentation

### Core Technologies

- **`datastar/`** - Datastar framework for reactive SSE (Server-Sent Events) and real-time UI updates

  - SSE event handling patterns
  - Client-side reactive bindings (`data-on:*`, `data-bind`)
  - Signal filtering and request configuration
  - Server-side SSE generation in Go
  - **`toc.md`** - Table of contents with line numbers for quick navigation

- **`daisyui/`** - DaisyUI component library and Tailwind CSS integration

  - Component examples and patterns
  - Installation and configuration
  - Semantic tokens and theming
  - **`toc.md`** - Table of contents with line numbers for quick navigation

- **`nats/`** - NATS messaging system for pub/sub real-time updates

  - Client installation and connection
  - Publishing and subscribing patterns
  - Request/reply patterns
  - Fanout and service mesh examples
  - **`toc.md`** - Table of contents with line numbers for quick navigation

- **`sqlc/`** - Type-safe SQL query generator

  - Configuration and setup
  - Writing SQL queries with annotations
  - Generated Go code patterns
  - Database driver integration (PostgreSQL, SQLite, MySQL)
  - **`toc.md`** - Table of contents with line numbers for quick navigation

- **`templ/`** - Go template engine for HTML generation
  - Template syntax and components
  - Installation and tooling setup
  - Integration with Go handlers
  - Live reload configuration
  - **`toc.md`** - Table of contents with line numbers for quick navigation

### Go Package API Reference

- **`go-packages/datastar-go/`** - Datastar Go SDK API reference (SSE helpers, signal reading, compression)
- **`go-packages/chi/`** - Chi router API reference (routing, middleware, context)
- **`go-packages/pgx/`** - PostgreSQL driver API reference (connection pooling, query execution, types)
- **`go-packages/nats-go/`** - NATS messaging client API reference (publish/subscribe, JetStream, key-value)
- **`go-packages/gorilla-sessions/`** - Session management API reference (store interface, session handling)

**To regenerate Go package documentation:**
```bash
task context:godoc
```

## How to Use These Files

### Table of Contents (TOC) Files

Each subdirectory contains a `toc.md` file with:
- **Line numbers** (e.g., `L0001`, `L0042`) for precise navigation
- **Header hierarchy** showing document structure
- **Grep examples** for searching specific sections

**To regenerate all TOC files** after updating documentation:
```bash
task context:toc
```

### Quick Search with Grep Tool

When you need to find specific information quickly, use the grep tool to search across all documentation. First check the `toc.md` file to get line numbers, then read specific sections.

**Example workflow:**
1. Read `context/datastar/toc.md` to find relevant sections
2. Use line numbers to read specific sections from `context/datastar/datastar.md`
3. Use grep to search for keywords across all files

### Reading Full Documentation

For comprehensive understanding, read the entire file using the Read tool. Each file contains:

- Installation instructions
- Code examples with explanations
- Best practices from official documentation
- Common patterns used in this project

### When to Reference These Files

- **Before implementing SSE handlers** → Read `context/datastar/toc.md` then `datastar.md` and `context/go-packages/datastar-go/`
- **Before writing database queries** → Read `context/sqlc/toc.md` then `sqlc.md` and `context/go-packages/pgx/` for proper annotation syntax
- **Before creating UI components** → Read `context/templ/toc.md` then `templ.md` and `context/daisyui/daisyui.md`
- **When debugging real-time features** → Reference `context/datastar/toc.md`, `context/nats/toc.md`, and `context/go-packages/nats-go/`
- **When adding new database operations** → Reference `context/sqlc/toc.md` and `context/go-packages/pgx/` for query patterns
- **When working with routing** → Reference `context/go-packages/chi/` for router API details
- **When implementing sessions** → Reference `context/go-packages/gorilla-sessions/` for session management

## Integration in This Project

These technologies work together in Northstar:

1. **Templ** renders HTML components on the server
2. **Datastar** (frontend + `datastar-go` backend SDK) establishes SSE connections for real-time updates
3. **NATS** (`nats-go` client) broadcasts change events between services
4. **sqlc** provides type-safe database queries using **pgx** for both SQLite (app state) and PostgreSQL (target databases)
5. **Chi** router handles HTTP routing and middleware
6. **Gorilla Sessions** manages user session state

## Searching Tips for Agents

### Step 1: Check TOC Files First
Always start by reading the relevant `toc.md` file to get an overview of available sections and their line numbers. This helps you:
- Understand the document structure
- Find relevant sections quickly
- Use precise line numbers for targeted reading

### Step 2: Targeted Search
- Use **specific keywords** related to your task (e.g., "INSERT", "Subscribe", "component", "SSE")
- Search for **error messages** if debugging (e.g., "failed to", "error")
- Look for **code patterns** you want to replicate (e.g., "example", "tutorial")
- Find **configuration** details (e.g., "config", "setup", "install")

### Example Workflow
```bash
# 1. Check TOC to find SSE-related sections
Read: context/datastar/toc.md

# 2. Found "L0043 ### JavaScript SSE Event Handling Setup"
# Now read that specific section
Read: context/datastar/datastar.md (lines 43-80)

# 3. Search for more SSE examples
Grep: pattern="SSE" path="context/datastar"
```
