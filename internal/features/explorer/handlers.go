// internal/features/explorer/handlers.go
package explorer

import (
	"net/http"

	"northstar/internal/domain"
	"northstar/internal/features/explorer/components"
	"northstar/internal/features/explorer/pages"
	"northstar/internal/features/explorer/services"

	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/gorilla/sessions"
)

type Handlers struct {
	sessionStore sessions.Store
	ns           *embeddednats.Server
	s            *services.ExplorerService
}

func NewHandlers(sessionStore sessions.Store, ns *embeddednats.Server, explorerService *services.ExplorerService) *Handlers {
	return &Handlers{
		sessionStore: sessionStore,
		ns:           ns,
		s:            explorerService,
	}
}

func (h *Handlers) ExplorerPage(w http.ResponseWriter, r *http.Request) {

	schemas, err := h.s.GetSchemaOverview(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	activeSchemaName := schemas[0].Name
	activeTableName := schemas[0].Tables[0].Name
	activeTable := schemas[0].Tables[0]

	tableColumns, err := h.s.GetTableColumns(r.Context(), activeSchemaName, activeTableName)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data := toExplorerPageData(
		schemas,
		activeSchemaName,
		activeTableName,
		activeTable,
		tableColumns,
	)

	if err := pages.ExplorerPage(data).Render(r.Context(), w); err != nil {

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func toExplorerPageData(
	schemas []domain.Schema,
	activeSchemaName, activeTableName string,
	activeTable domain.Table,
	activeColumns []domain.Column,
) pages.ExplorerPageData {

	// Map Schemas for the Sidebar
	sidebarSchemas := make([]components.SchemaView, len(schemas))
	for i, s := range schemas {
		tables := make([]components.TableView, len(s.Tables))
		for j, t := range s.Tables {
			tables[j] = components.TableView{
				Table:    t,
				IsActive: s.Name == activeSchemaName && t.Name == activeTableName,
			}
		}

		sidebarSchemas[i] = components.SchemaView{
			Schema: s,
			Tables: tables,
		}
	}

	// Map Data for the Main Content Area
	mainContent := components.MainContentData{
		Breadcrumbs: []components.BreadcrumbItem{
			{Name: "Databases", Href: "/"},
			{Name: activeSchemaName},
			{Name: activeTableName},
		},
		TableDetails: activeTable,
		Columns:      activeColumns,
		// Mocked table data for display
		TableData: components.TableData{
			Headers: []string{"id", "product_name", "status"},
			Rows: [][]any{
				{1, "Gadget Pro", "active"},
				{2, "Widget Max", "inactive"},
			},
		},
	}

	return pages.ExplorerPageData{
		SidebarSchemas: sidebarSchemas,
		MainContent:    mainContent,
	}
}
