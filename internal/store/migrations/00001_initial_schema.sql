-- +goose Up
-- PostgresUI SQLite Schema
-- This database stores all application configuration and state
-- It does NOT modify the target PostgreSQL databases

-- Database connections
CREATE TABLE connections (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    host TEXT NOT NULL,
    port INTEGER NOT NULL DEFAULT 5432,
    database TEXT NOT NULL,
    username TEXT NOT NULL,
    encrypted_password TEXT NOT NULL,
    ssl_mode TEXT NOT NULL DEFAULT 'prefer',
    ssl_config TEXT,
    is_active INTEGER NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_connections_active ON connections(is_active);

-- Tables exposed from PostgreSQL databases
CREATE TABLE exposed_tables (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    connection_id INTEGER NOT NULL,
    schema_name TEXT NOT NULL DEFAULT 'public',
    table_name TEXT NOT NULL,
    display_name TEXT NOT NULL,
    description TEXT,
    is_active INTEGER NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (connection_id) REFERENCES connections(id) ON DELETE CASCADE,
    UNIQUE(connection_id, schema_name, table_name)
);

CREATE INDEX idx_exposed_tables_connection ON exposed_tables(connection_id);
CREATE INDEX idx_exposed_tables_active ON exposed_tables(is_active);

-- Columns from exposed tables
CREATE TABLE exposed_columns (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    exposed_table_id INTEGER NOT NULL,
    column_name TEXT NOT NULL,
    display_name TEXT NOT NULL,
    data_type TEXT NOT NULL,
    is_visible INTEGER NOT NULL DEFAULT 1,
    is_editable INTEGER NOT NULL DEFAULT 1,
    display_order INTEGER NOT NULL DEFAULT 0,
    validation_rules TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (exposed_table_id) REFERENCES exposed_tables(id) ON DELETE CASCADE,
    UNIQUE(exposed_table_id, column_name)
);

CREATE INDEX idx_exposed_columns_table ON exposed_columns(exposed_table_id);
CREATE INDEX idx_exposed_columns_visible ON exposed_columns(is_visible);

-- User accounts
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    display_name TEXT NOT NULL,
    hashed_password TEXT,
    sso_provider TEXT,
    sso_id TEXT,
    is_active INTEGER NOT NULL DEFAULT 1,
    is_admin INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_login_at DATETIME
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_sso ON users(sso_provider, sso_id);
CREATE INDEX idx_users_active ON users(is_active);

-- RBAC roles
CREATE TABLE roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    is_system_role INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_roles_system ON roles(is_system_role);

-- User-to-Role mapping (many-to-many)
CREATE TABLE user_roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    UNIQUE(user_id, role_id)
);

CREATE INDEX idx_user_roles_user ON user_roles(user_id);
CREATE INDEX idx_user_roles_role ON user_roles(role_id);

-- Granular permissions for roles on tables
CREATE TABLE permissions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    role_id INTEGER NOT NULL,
    exposed_table_id INTEGER NOT NULL,
    can_view INTEGER NOT NULL DEFAULT 0,
    can_create INTEGER NOT NULL DEFAULT 0,
    can_update INTEGER NOT NULL DEFAULT 0,
    can_delete INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (exposed_table_id) REFERENCES exposed_tables(id) ON DELETE CASCADE,
    UNIQUE(role_id, exposed_table_id)
);

CREATE INDEX idx_permissions_role ON permissions(role_id);
CREATE INDEX idx_permissions_table ON permissions(exposed_table_id);

-- Row-Level Security rules
CREATE TABLE row_level_security_rules (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    exposed_table_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    rule_name TEXT NOT NULL,
    filter_expression TEXT NOT NULL,
    is_active INTEGER NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (exposed_table_id) REFERENCES exposed_tables(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

CREATE INDEX idx_rls_rules_table ON row_level_security_rules(exposed_table_id);
CREATE INDEX idx_rls_rules_role ON row_level_security_rules(role_id);
CREATE INDEX idx_rls_rules_active ON row_level_security_rules(is_active);

-- User preferences (column widths, sort order, etc.)
CREATE TABLE user_preferences (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    exposed_table_id INTEGER,
    preference_key TEXT NOT NULL,
    preference_value TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (exposed_table_id) REFERENCES exposed_tables(id) ON DELETE CASCADE,
    UNIQUE(user_id, exposed_table_id, preference_key)
);

CREATE INDEX idx_user_prefs_user ON user_preferences(user_id);
CREATE INDEX idx_user_prefs_table ON user_preferences(exposed_table_id);

-- Saved filters
CREATE TABLE saved_filters (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    exposed_table_id INTEGER NOT NULL,
    filter_name TEXT NOT NULL,
    filter_config TEXT NOT NULL,
    is_public INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (exposed_table_id) REFERENCES exposed_tables(id) ON DELETE CASCADE
);

CREATE INDEX idx_saved_filters_user ON saved_filters(user_id);
CREATE INDEX idx_saved_filters_table ON saved_filters(exposed_table_id);
CREATE INDEX idx_saved_filters_public ON saved_filters(is_public);

-- Audit logs for compliance
CREATE TABLE audit_logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    exposed_table_id INTEGER,
    operation_type TEXT NOT NULL,
    row_identifier TEXT,
    changed_data TEXT,
    ip_address TEXT,
    user_agent TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (exposed_table_id) REFERENCES exposed_tables(id) ON DELETE SET NULL
);

CREATE INDEX idx_audit_logs_user ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_table ON audit_logs(exposed_table_id);
CREATE INDEX idx_audit_logs_created ON audit_logs(created_at);
CREATE INDEX idx_audit_logs_operation ON audit_logs(operation_type);

-- SSO provider configurations
CREATE TABLE sso_providers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    provider_name TEXT NOT NULL UNIQUE,
    provider_type TEXT NOT NULL,
    client_id TEXT NOT NULL,
    encrypted_client_secret TEXT NOT NULL,
    auth_url TEXT,
    token_url TEXT,
    user_info_url TEXT,
    config TEXT,
    is_active INTEGER NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_sso_providers_active ON sso_providers(is_active);

-- Excel/CSV upload jobs
CREATE TABLE upload_jobs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    exposed_table_id INTEGER NOT NULL,
    filename TEXT NOT NULL,
    row_count INTEGER NOT NULL DEFAULT 0,
    status TEXT NOT NULL DEFAULT 'pending',
    error_message TEXT,
    mapping_config TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (exposed_table_id) REFERENCES exposed_tables(id) ON DELETE CASCADE
);

CREATE INDEX idx_upload_jobs_user ON upload_jobs(user_id);
CREATE INDEX idx_upload_jobs_table ON upload_jobs(exposed_table_id);
CREATE INDEX idx_upload_jobs_status ON upload_jobs(status);
CREATE INDEX idx_upload_jobs_created ON upload_jobs(created_at);

-- User sessions
CREATE TABLE sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    session_token TEXT NOT NULL UNIQUE,
    ip_address TEXT,
    user_agent TEXT,
    expires_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_sessions_token ON sessions(session_token);
CREATE INDEX idx_sessions_user ON sessions(user_id);
CREATE INDEX idx_sessions_expires ON sessions(expires_at);

-- Trigger to update 'updated_at' timestamps
-- +goose StatementBegin
CREATE TRIGGER update_connections_timestamp 
AFTER UPDATE ON connections
FOR EACH ROW
BEGIN
    UPDATE connections SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_exposed_tables_timestamp 
AFTER UPDATE ON exposed_tables
FOR EACH ROW
BEGIN
    UPDATE exposed_tables SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_exposed_columns_timestamp 
AFTER UPDATE ON exposed_columns
FOR EACH ROW
BEGIN
    UPDATE exposed_columns SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_users_timestamp 
AFTER UPDATE ON users
FOR EACH ROW
BEGIN
    UPDATE users SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_roles_timestamp 
AFTER UPDATE ON roles
FOR EACH ROW
BEGIN
    UPDATE roles SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_permissions_timestamp 
AFTER UPDATE ON permissions
FOR EACH ROW
BEGIN
    UPDATE permissions SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_rls_rules_timestamp 
AFTER UPDATE ON row_level_security_rules
FOR EACH ROW
BEGIN
    UPDATE row_level_security_rules SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_user_preferences_timestamp 
AFTER UPDATE ON user_preferences
FOR EACH ROW
BEGIN
    UPDATE user_preferences SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_saved_filters_timestamp 
AFTER UPDATE ON saved_filters
FOR EACH ROW
BEGIN
    UPDATE saved_filters SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_sso_providers_timestamp 
AFTER UPDATE ON sso_providers
FOR EACH ROW
BEGIN
    UPDATE sso_providers SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS update_sso_providers_timestamp;
DROP TRIGGER IF EXISTS update_saved_filters_timestamp;
DROP TRIGGER IF EXISTS update_user_preferences_timestamp;
DROP TRIGGER IF EXISTS update_rls_rules_timestamp;
DROP TRIGGER IF EXISTS update_permissions_timestamp;
DROP TRIGGER IF EXISTS update_roles_timestamp;
DROP TRIGGER IF EXISTS update_users_timestamp;
DROP TRIGGER IF EXISTS update_exposed_columns_timestamp;
DROP TRIGGER IF EXISTS update_exposed_tables_timestamp;
DROP TRIGGER IF EXISTS update_connections_timestamp;

DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS upload_jobs;
DROP TABLE IF EXISTS sso_providers;
DROP TABLE IF EXISTS audit_logs;
DROP TABLE IF EXISTS saved_filters;
DROP TABLE IF EXISTS user_preferences;
DROP TABLE IF EXISTS row_level_security_rules;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS exposed_columns;
DROP TABLE IF EXISTS exposed_tables;
DROP TABLE IF EXISTS connections;
