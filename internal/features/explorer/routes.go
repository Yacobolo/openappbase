package explorer

// import (
// 	"northstar/internal/features/explorer/services"
// 	"northstar/internal/store"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/gorilla/sessions"
// 	"github.com/nats-io/nats.go/jetstream"
// )

// func SetupRoutes(router chi.Router, store sessions.Store, js jetstream.JetStream, q *store.Queries) error {
// 	explorerService, err := services.NewExplorerService(q, js, store)

// 	if err != nil {
// 		return err
// 	}

// 	handlers := NewHandlers(store, explorerService)
// 	// setup api routes

// 	router.Get("/explorer", handlers.ExplorerPage)
// 	router.Get("/explorer/{schemaName}", handlers.ExplorerPage)
// 	router.Get("/explorer/{schemaName}/{tableName}", handlers.ExplorerPage)

// 	return nil
// }
