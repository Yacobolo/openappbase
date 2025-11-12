package services

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"northstar/internal/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectionService handles PostgreSQL connection management
type ConnectionService struct {
	queries       *store.Queries
	encryptionKey []byte // 32 bytes for AES-256
}

// NewConnectionService creates a new connection service
func NewConnectionService(queries *store.Queries, encryptionKey string) (*ConnectionService, error) {
	// Ensure encryption key is 32 bytes for AES-256
	key := []byte(encryptionKey)
	if len(key) != 32 {
		return nil, fmt.Errorf("encryption key must be exactly 32 bytes, got %d", len(key))
	}

	return &ConnectionService{
		queries:       queries,
		encryptionKey: key,
	}, nil
}

// CreateConnectionInput holds parameters for creating a connection
type CreateConnectionInput struct {
	Name      string
	Host      string
	Port      int64
	Database  string
	Username  string
	Password  string
	SSLMode   string
	SSLConfig *string
}

// UpdateConnectionInput holds parameters for updating a connection
type UpdateConnectionInput struct {
	ID        int64
	Name      string
	Host      string
	Port      int64
	Database  string
	Username  string
	Password  string
	SSLMode   string
	SSLConfig *string
}

// ConnectionInfo represents connection details with decrypted password
type ConnectionInfo struct {
	ID        int64
	Name      string
	Host      string
	Port      int64
	Database  string
	Username  string
	Password  string
	SSLMode   string
	SSLConfig *string
	IsActive  bool
}

// CreateConnection creates a new database connection
func (s *ConnectionService) CreateConnection(ctx context.Context, input CreateConnectionInput) (*store.Connection, error) {
	// Encrypt the password
	encryptedPassword, err := s.encryptPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt password: %w", err)
	}

	// Create the connection record
	params := store.CreateConnectionParams{
		Name:              input.Name,
		Host:              input.Host,
		Port:              input.Port,
		Database:          input.Database,
		Username:          input.Username,
		EncryptedPassword: encryptedPassword,
		SslMode:           input.SSLMode,
		SslConfig:         sqlNullString(input.SSLConfig),
	}

	conn, err := s.queries.CreateConnection(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	return &conn, nil
}

// GetConnection retrieves a connection by ID
func (s *ConnectionService) GetConnection(ctx context.Context, id int64) (*store.Connection, error) {
	conn, err := s.queries.GetConnection(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}
	return &conn, nil
}

// GetConnectionInfo retrieves connection with decrypted password
func (s *ConnectionService) GetConnectionInfo(ctx context.Context, id int64) (*ConnectionInfo, error) {
	conn, err := s.GetConnection(ctx, id)
	if err != nil {
		return nil, err
	}

	password, err := s.decryptPassword(conn.EncryptedPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt password: %w", err)
	}

	var sslConfig *string
	if conn.SslConfig.Valid {
		sslConfig = &conn.SslConfig.String
	}

	return &ConnectionInfo{
		ID:        conn.ID,
		Name:      conn.Name,
		Host:      conn.Host,
		Port:      conn.Port,
		Database:  conn.Database,
		Username:  conn.Username,
		Password:  password,
		SSLMode:   conn.SslMode,
		SSLConfig: sslConfig,
		IsActive:  conn.IsActive == 1,
	}, nil
}

// ListConnections retrieves all connections
func (s *ConnectionService) ListConnections(ctx context.Context) ([]store.Connection, error) {
	conns, err := s.queries.ListConnections(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list connections: %w", err)
	}
	return conns, nil
}

// ListActiveConnections retrieves only active connections
func (s *ConnectionService) ListActiveConnections(ctx context.Context) ([]store.Connection, error) {
	conns, err := s.queries.ListActiveConnections(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list active connections: %w", err)
	}
	return conns, nil
}

// UpdateConnection updates an existing connection
func (s *ConnectionService) UpdateConnection(ctx context.Context, input UpdateConnectionInput) (*store.Connection, error) {
	// Encrypt the password
	encryptedPassword, err := s.encryptPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt password: %w", err)
	}

	params := store.UpdateConnectionParams{
		Name:              input.Name,
		Host:              input.Host,
		Port:              input.Port,
		Database:          input.Database,
		Username:          input.Username,
		EncryptedPassword: encryptedPassword,
		SslMode:           input.SSLMode,
		SslConfig:         sqlNullString(input.SSLConfig),
		ID:                input.ID,
	}

	conn, err := s.queries.UpdateConnection(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update connection: %w", err)
	}

	return &conn, nil
}

// DeleteConnection deletes a connection by ID
func (s *ConnectionService) DeleteConnection(ctx context.Context, id int64) error {
	if err := s.queries.DeleteConnection(ctx, id); err != nil {
		return fmt.Errorf("failed to delete connection: %w", err)
	}
	return nil
}

// SetConnectionStatus enables or disables a connection
func (s *ConnectionService) SetConnectionStatus(ctx context.Context, id int64, isActive bool) error {
	var status int64 = 0
	if isActive {
		status = 1
	}

	params := store.UpdateConnectionStatusParams{
		IsActive: status,
		ID:       id,
	}

	if err := s.queries.UpdateConnectionStatus(ctx, params); err != nil {
		return fmt.Errorf("failed to update connection status: %w", err)
	}
	return nil
}

// TestConnection attempts to connect to the PostgreSQL database
func (s *ConnectionService) TestConnection(ctx context.Context, id int64) error {
	connInfo, err := s.GetConnectionInfo(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get connection info: %w", err)
	}

	return s.testConnectionWithInfo(ctx, connInfo)
}

// TestConnectionDirect tests a connection without saving it first
func (s *ConnectionService) TestConnectionDirect(ctx context.Context, input CreateConnectionInput) error {
	connInfo := &ConnectionInfo{
		Name:      input.Name,
		Host:      input.Host,
		Port:      input.Port,
		Database:  input.Database,
		Username:  input.Username,
		Password:  input.Password,
		SSLMode:   input.SSLMode,
		SSLConfig: input.SSLConfig,
	}

	return s.testConnectionWithInfo(ctx, connInfo)
}

// testConnectionWithInfo performs the actual connection test
func (s *ConnectionService) testConnectionWithInfo(ctx context.Context, info *ConnectionInfo) error {
	// Build PostgreSQL connection string
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		info.Host, info.Port, info.Username, info.Password, info.Database, info.SSLMode,
	)

	// Attempt to create a connection pool
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}
	defer pool.Close()

	// Test the connection with a simple ping
	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

// GetDatabasePool creates a new pgxpool for a given connection ID
func (s *ConnectionService) GetDatabasePool(ctx context.Context, id int64) (*pgxpool.Pool, error) {
	connInfo, err := s.GetConnectionInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	if !connInfo.IsActive {
		return nil, errors.New("connection is not active")
	}

	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		connInfo.Host, connInfo.Port, connInfo.Username, connInfo.Password, connInfo.Database, connInfo.SSLMode,
	)

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	return pool, nil
}

// encryptPassword encrypts a password using AES-256-GCM
func (s *ConnectionService) encryptPassword(password string) (string, error) {
	block, err := aes.NewCipher(s.encryptionKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(password), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// decryptPassword decrypts an encrypted password
func (s *ConnectionService) decryptPassword(encryptedPassword string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedPassword)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(s.encryptionKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// sqlNullString converts a *string to sql.NullString
func sqlNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}
