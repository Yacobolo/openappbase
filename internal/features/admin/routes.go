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

	// API routes
	router.Route("/api/admin", func(adminrouter chi.Router) {
		adminrouter.Route("/connections", func(connectionsRouter chi.Router) {
			// SSE endpoint for real-time updates
			connectionsRouter.Get("/", handlers.AdminConnectionsSSE)
			// RESTful CRUD endpoints
			connectionsRouter.Post("/", handlers.CreateConnection)
			connectionsRouter.Route("/{id}", func(connectionRouter chi.Router) {
				connectionRouter.Post("/test", handlers.TestConnection)
				connectionRouter.Delete("/", handlers.DeleteConnection)
			})
		})
	})

	return nil
}
