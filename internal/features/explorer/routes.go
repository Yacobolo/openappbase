package explorer

import (
	"northstar/internal/features/explorer/services"
	"northstar/internal/store"

	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupRoutes(router chi.Router, sessionStore sessions.Store, ns *embeddednats.Server, q *store.Queries) error {
	explorerService := services.NewExplorerService(q)
	handlers := NewHandlers(sessionStore, ns, explorerService)

	router.Get("/explorer", handlers.ExplorerPage)

	return nil
}
