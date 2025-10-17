package editor

import (
	"fmt"
	"net/http"

	"northstar/internal/features/editor/pages"
	"northstar/internal/store"

	"github.com/gorilla/sessions"
	"github.com/starfederation/datastar-go/datastar"
)

type Handlers struct {
	sessionStore sessions.Store
}

func NewHandlers(sessionStore sessions.Store) *Handlers {
	return &Handlers{
		sessionStore: sessionStore,
	}
}

func (h *Handlers) EditorPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.EditorPage().Render(r.Context(), w); err != nil {

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	sse := datastar.NewSSE(w, r)
	data, err := store.GetTableData("test")

	if err != nil {
		sse.ConsoleLog(fmt.Sprintf("Error fetching table data: %v", err))
		return
	}

	for _, row := range data {
		sse.ConsoleLog(fmt.Sprintf("%+v", row))
	}
}
