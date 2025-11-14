# Project Title (e.g., "OpenPostgresSheet" or "DB-Bizview")

## One-Page Project Description

### The Problem

In modern organizations, critical business data is often scattered across countless disconnected Excel spreadsheets and CSV files. This manual process is inefficient, prone to error, and fundamentally unscalable.

While "CRUD app" builders and no-code platforms like Airtable or NocoDB offer a partial solution, they force critical limitations on serious businesses:

- **Row Limits:** They impose strict caps on the number of records, penalizing growth.
- **Limited "Bring Your Own DB" Support:** While some platforms _can_ connect to an existing PostgreSQL database, they fail to provide the necessary granular security—like table and column-level permissions—making this feature unusable in a secure enterprise environment.
- **Expensive Enterprise Features:** Essential needs like granular Role-Based Access Control (RBAC), Row-Level Security, and Single Sign-On (SSO) integration are often missing entirely or locked behind expensive, usage-based subscription plans.

There is no open, scalable solution that empowers business users to manage their data directly within their company's existing PostgreSQL infrastructure _securely_.

### The Solution

This application is an open-source, self-hosted platform that instantly generates a powerful, spreadsheet-like CRUD application on top of **any existing PostgreSQL database**.

It acts as a secure data management portal, allowing non-technical business users to safely view, update, and upload data, while giving administrators full control over access and security.

The application **does not pollute your database**. It maintains all of its own configuration—roles, permissions, and selected tables—in a separate, self-contained SQLite database, leaving your core PostgreSQL schema clean.

### Core Features

- **Connect to Any PostgreSQL:** Instantly connect to an existing database. The application introspects the schema and requires no changes to your tables.
- **Selective Table Exposure:** Administrators can securely select _which_ tables and columns are exposed to business users, keeping sensitive data hidden.
- **Powerful Spreadsheet UI:** A fast, intuitive, spreadsheet-like interface for viewing, editing, and filtering data.
- **Real-time Updates:** Changes made by one user are broadcast in real-time to all other users viewing the same data, ensuring everyone is working from the latest version.
- **Enterprise-Grade Security:**
  - **Role-Based Access Control (RBAC):** Create roles (e.g., "Admin," "Editor," "Viewer") and define granular permissions.
  - **Row-Level Security (RLS):** Define rules to restrict row access (e.g., "Sales managers can only see data for their own region").
  - **SSO Integration:** Authenticate users through your existing identity provider (OIDC, SAML).
- **Easy Data Onboarding:**
  - **Excel & CSV Upload:** Easily upload spreadsheet data directly into the selected database table, with validation and mapping.
- **Clean & Isolated:** The application's core logic is stored in its own SQLite database. It adds **zero tables, columns, or logic** to your target PostgreSQL database.

### Architecture Overview

The application uses a "Targeted Ping" architecture to achieve scalable, real-time updates without high server load.

1. **PostgreSQL Triggers:** `FOR EACH ROW` triggers on monitored tables send a rich JSON payload (containing table name, operation type, and a shared key like `project_id`) to a single PostgreSQL `NOTIFY` channel.
2. **Go Bridge Service:** A single, persistent Go service `LISTEN`s to this channel, parses the JSON, and republishes the event to a dynamic, targeted NATS subject (e.g., `todos_updates.project_abc`).
3. **NATS Pub/Sub:** Serves as a lightweight, scalable message bus for these targeted "ping" events.
4. **Go SSE Handlers:** When a user views a dataset (e.g., "project_abc"), their SSE handler `Subscribe`s to the specific NATS subject (`todos_updates.project_abc`). When a ping is received, it re-queries the database for fresh data, ensuring only relevant clients are updated.
5. **SQLite Database:** All "single-player" application state, such as user column preferences, saved filters, and RBAC rules, is stored persistently in the app's local SQLite database.

### Tech Stack

- **Backend:** Go (Golang)
- **Frontend:** Datastar & Templ
- **Frontend Styling:** Tailwind CSS
- **Application Database:** SQLite (for storing roles, user preferences, and app state)
- **Database Migrations:** Goose (for managing schema changes)
- **Real-time Messaging:** Embedded NATS (for high-performance Pub/Sub)
- **SSO / Authentication:** markbates/goth

### Use Cases

1. **Empower Business Users:** Give your finance, operations, or marketing teams a secure way to manage their own "lookup" data (e.g., product lists, pricing tiers, user-profiles) without filing IT tickets.
2. **Replace Scattered Spreadsheets:** Consolidate dozens of disconnected Excel files into a single, secure, version-controlled database with a user-friendly interface.
3. **Safe Admin Interface:** Provide a safe, access-controlled "admin panel" for an existing application without building a custom one from scratch.

_This project empowers organizations to scale their data management by finally connecting their business users directly and securely to their central database._

---

## Development Roadmap

- [ ] **Phase 1:** Admin Panel - Database Connections 🚧
- [ ] **Phase 2:** Table & Column Exposure Management (read-only)
- [ ] **Phase 3:** Role & Permission System (RBAC)
- [ ] **Phase 4:** User Management & Authentication
- [ ] **Phase 5:** Row-Level Security (RLS)
- [ ] **Phase 6:** Saved Filters & User Preferences
- [ ] **Phase 7:** Excel/CSV Upload
- [ ] **Phase 8:** Real-time Updates (LISTEN/NOTIFY + SSE)

---

## Development Setup

### Database Migrations

This project uses [Goose](https://github.com/pressly/goose) for database migrations.

**Available commands:**

```bash
# Apply all pending migrations
task db:migrate:up

# Rollback the last migration
task db:migrate:down

# Check migration status
task db:migrate:status

# Create a new migration (provide name as argument)
task db:migrate:create -- add_new_feature

# Reset database (WARNING: deletes all data)
task db:reset
```

**Migration workflow:**

1. Create a new migration: `task db:migrate:create -- descriptive_name`
2. Edit the generated file in `internal/store/migrations/`
3. Apply the migration: `task db:migrate:up`
4. If you need to use triggers or other multi-line statements, wrap them in `-- +goose StatementBegin` and `-- +goose StatementEnd` directives

**Note:** Migrations are stored in `internal/store/migrations/` and applied to `data/app.db`
