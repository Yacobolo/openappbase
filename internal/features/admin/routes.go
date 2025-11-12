package admin

import (
	"northstar/config"
	"northstar/internal/features/admin/services"
	"northstar/internal/store"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router chi.Router, q *store.Queries) error {

	// 32 byte encrept key for secure cookies
	encryptionKey := config.Global.EncryptionKey

	connectionService, err := services.NewConnectionService(q, encryptionKey)

	if err != nil {
		return err
	}

	handlers := NewHandlers(connectionService)

	router.Get("/admin", handlers.AdminPage)

	// router.Route("/api", func(apirouter chi.Router) {
	// 	apirouter.Route("/admin", func(adminrouter chi.Router) {
	// 		adminrouter.Get("/", handlers.AdminSSE)
	// 	})
	// })

	return nil
}
