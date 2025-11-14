//go:build dev
// +build dev

package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// Dev-safe default encryption key (ONLY used in dev builds when ENCRYPTION_KEY is not set)
// This allows worktrees and fresh dev environments to work immediately
// Production builds will still require a proper ENCRYPTION_KEY
const devDefaultEncryptionKey = "01234567890123456789012345678901"

func Load() *Config {
	// Load .env file if it exists (ignores error if file doesn't exist)
	godotenv.Load()

	// In dev mode, provide safe defaults for required environment variables
	encryptionKey := os.Getenv("ENCRYPTION_KEY")
	if encryptionKey == "" {
		slog.Warn("ENCRYPTION_KEY not set, using dev default (NOT SAFE FOR PRODUCTION)")
		encryptionKey = devDefaultEncryptionKey
	}

	ValidateEncryptionKey(encryptionKey)

	cfg := &Config{
		Environment: Dev,
		Host:        getEnv("HOST", "0.0.0.0"),
		Port:        getEnv("PORT", "8080"),
		LogLevel: func() slog.Level {
			switch os.Getenv("LOG_LEVEL") {
			case "DEBUG":
				return slog.LevelDebug
			case "INFO":
				return slog.LevelInfo
			case "WARN":
				return slog.LevelWarn
			case "ERROR":
				return slog.LevelError
			default:
				return slog.LevelInfo
			}
		}(),
		SessionSecret: getEnv("SESSION_SECRET", "session-secret"),
		EncryptionKey: encryptionKey,
	}

	return cfg
}
