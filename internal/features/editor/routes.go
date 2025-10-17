package editor

import (
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupRoutes(router chi.Router, sessionStore sessions.Store) error {
	handlers := NewHandlers(sessionStore)

	router.Get("/editor", handlers.EditorPage)

	// router.Route("/editor/increment", func(incrementRouter chi.Router) {
	// 	incrementRouter.Post("/global", handlers.IncrementGlobal)
	// 	incrementRouter.Post("/user", handlers.IncrementUser)
	// })

	return nil
}
