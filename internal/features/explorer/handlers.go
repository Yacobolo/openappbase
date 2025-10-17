package explorer

import (
	"net/http"

	"northstar/internal/features/explorer/pages"

	"github.com/gorilla/sessions"
)

type Handlers struct {
	sessionStore sessions.Store
}

func NewHandlers(sessionStore sessions.Store) *Handlers {
	return &Handlers{
		sessionStore: sessionStore,
	}
}

func (h *Handlers) ExplorerPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.ExplorerPage().Render(r.Context(), w); err != nil {

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
