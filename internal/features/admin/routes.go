package admin

import (
	"northstar/config"
	"northstar/internal/features/admin/services"
	"northstar/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/nats-io/nats.go"
)

func SetupRoutes(router chi.Router, q *store.Queries, nc *nats.Conn) error {

	// 32 byte encrept key for secure cookies
	encryptionKey := config.Global.EncryptionKey

	connectionService, err := services.NewConnectionService(q, encryptionKey)

	if err != nil {
		return err
	}

	handlers := NewHandlers(connectionService, nc)

	// Page routes
	router.Get("/admin", handlers.AdminPage)

	// SSE API routes for Datastar
	router.Route("/api/admin", func(adminrouter chi.Router) {
		adminrouter.Get("/connections", handlers.AdminConnectionsSSE)
		adminrouter.Post("/connections/create", handlers.CreateConnectionSSE)
		adminrouter.Post("/connections/test", handlers.TestConnectionSSE)
		adminrouter.Post("/connections/delete", handlers.DeleteConnectionSSE)
	})

	return nil
}
