package editor

import (
	"northstar/internal/features/editor/services"
	"northstar/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(router chi.Router, sessionStore sessions.Store, pool *pgxpool.Pool, q *store.Queries) error {
	editorService := services.NewEditorService(q, pool)
	handlers := NewHandlers(sessionStore, editorService)

	router.Get("/editor", handlers.EditorPage)
	router.Get("/editor/load", handlers.LoadTableData)
	router.Get("/editor/change-page", handlers.ChangePage)

	return nil
}
