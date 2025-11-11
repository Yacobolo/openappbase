package editor

import (
	"context"
	"fmt"
	"northstar/internal/features/editor/services"
	"northstar/internal/session"
	"northstar/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go/jetstream"
)

func SetupRoutes(router chi.Router, sessionStore sessions.Store, js jetstream.JetStream, pool *pgxpool.Pool, q *store.Queries) error {
	// Get editor_sessions KV bucket
	kv, err := js.KeyValue(context.Background(), "editor_sessions")
	if err != nil {
		return fmt.Errorf("failed to get editor_sessions KV bucket: %w", err)
	}

	// Create StateStore for EditorState
	stateStore := session.NewStateStore[services.EditorState](kv, sessionStore)

	// Initialize editor service with state store
	editorService := services.NewEditorService(q, pool, stateStore)
	handlers := NewHandlers(sessionStore, editorService)

	router.Get("/editor", handlers.EditorPage)

	router.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Route("/editor", func(editorRouter chi.Router) {
			editorRouter.Get("/", handlers.EditorSSE)
		})
	})

	return nil
}
