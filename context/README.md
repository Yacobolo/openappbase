# Context Documentation

This directory contains up-to-date reference documentation for the key technologies used in the Northstar PostgreSQL UI project.

## Available Documentation

### Core Technologies

- **`datastar.md`** - Datastar framework for reactive SSE (Server-Sent Events) and real-time UI updates

  - SSE event handling patterns
  - Client-side reactive bindings (`data-on:*`, `data-bind`)
  - Signal filtering and request configuration
  - Server-side SSE generation in Go

- **`nats.md`** - NATS messaging system for pub/sub real-time updates

  - Client installation and connection
  - Publishing and subscribing patterns
  - Request/reply patterns
  - Fanout and service mesh examples

- **`sqlc.md`** - Type-safe SQL query generator

  - Configuration and setup
  - Writing SQL queries with annotations
  - Generated Go code patterns
  - Database driver integration (PostgreSQL, SQLite, MySQL)

- **`templ.md`** - Go template engine for HTML generation
  - Template syntax and components
  - Installation and tooling setup
  - Integration with Go handlers
  - Live reload configuration

## How to Use These Files

### Quick Search with Grep Tool

When you need to find specific information quickly, use the grep tool to search across all documentation

### Reading Full Documentation

For comprehensive understanding, read the entire file using the Read tool. Each file contains:

- Installation instructions
- Code examples with explanations
- Best practices from official documentation
- Common patterns used in this project

### When to Reference These Files

- **Before implementing SSE handlers** → Read `datastar.md` and `nats.md`
- **Before writing database queries** → Read `sqlc.md` for proper annotation syntax
- **Before creating UI components** → Read `templ.md` for template syntax
- **When debugging real-time features** → Reference `datastar.md` and `nats.md`
- **When adding new database operations** → Reference `sqlc.md` for query patterns

## Integration in This Project

These technologies work together in Northstar:

1. **Templ** renders HTML components on the server
2. **Datastar** establishes SSE connections for real-time updates
3. **NATS** broadcasts change events between services
4. **sqlc** provides type-safe database queries for both SQLite (app state) and PostgreSQL (target databases)

## Searching Tips for Agents

- Use **specific keywords** related to your task (e.g., "INSERT", "Subscribe", "component", "SSE")
- Search for **error messages** if debugging (e.g., "failed to", "error")
- Look for **code patterns** you want to replicate (e.g., "example", "tutorial")
- Find **configuration** details (e.g., "config", "setup", "install")
