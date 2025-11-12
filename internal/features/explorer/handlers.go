// internal/features/explorer/handlers.go
package explorer

// import (
// 	"fmt"
// 	"net/http"

// 	"northstar/internal/domain"
// 	"northstar/internal/features/explorer/components"
// 	"northstar/internal/features/explorer/pages"
// 	"northstar/internal/features/explorer/services"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/gorilla/sessions"
// )

// type Handlers struct {
// 	explorerService *services.ExplorerService
// }

// func NewHandlers(sessionStore sessions.Store, explorerService *services.ExplorerService) *Handlers {
// 	return &Handlers{
// 		explorerService: explorerService,
// 	}
// }

// func (h *Handlers) ExplorerPage(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	// 1. Get the active schema and table directly from the URL.
// 	activeSchemaName := chi.URLParam(r, "schemaName")
// 	activeTableName := chi.URLParam(r, "tableName")

// 	schemas, err := h.explorerService.GetSchemaOverview(ctx)
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	if len(schemas) == 0 {
// 		// Handle the case where the database is empty.
// 		pages.ExplorerPage(pages.ExplorerPageData{}).Render(ctx, w)
// 		return
// 	}

// 	if activeSchemaName == "" {
// 		activeSchemaName = schemas[0].Name // Default to the first schema
// 	}
// 	if activeTableName == "" {
// 		// Find the first table in the now-active schema.
// 		for _, s := range schemas {
// 			if s.Name == activeSchemaName && len(s.Tables) > 0 {
// 				activeTableName = s.Tables[0].Name
// 				break
// 			}
// 		}
// 	}

// 	activeTable := findActiveTable(schemas, activeSchemaName, activeTableName)

// 	if activeTable == nil {
// 		// The requested table does not exist. Return a 404 Not Found.
// 		http.Error(w, fmt.Sprintf("Table not found: %s.%s", activeSchemaName, activeTableName), http.StatusNotFound)
// 		return
// 	}

// 	tableColumns, err := h.explorerService.GetTableColumns(r.Context(), activeSchemaName, activeTableName)

// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	data := toExplorerPageData(
// 		schemas,
// 		activeSchemaName,
// 		activeTableName,
// 		*activeTable,
// 		tableColumns,
// 	)

// 	if err := pages.ExplorerPage(data).Render(r.Context(), w); err != nil {

// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 	}
// }

// func toExplorerPageData(
// 	schemas []domain.Schema,
// 	activeSchemaName, activeTableName string,
// 	activeTable domain.Table,
// 	activeColumns []domain.Column,
// ) pages.ExplorerPageData {

// 	// Map Schemas for the Sidebar
// 	sidebarSchemas := make([]components.SchemaView, len(schemas))
// 	for i, schema := range schemas {
// 		tables := make([]components.TableView, len(schema.Tables))
// 		for j, table := range schema.Tables {
// 			tables[j] = components.TableView{
// 				Table:    table,
// 				IsActive: schema.Name == activeSchemaName && table.Name == activeTableName,
// 			}
// 		}

// 		sidebarSchemas[i] = components.SchemaView{
// 			Name:   schema.Name,
// 			Tables: tables,
// 		}
// 	}

// 	// Map Data for the Main Content Area
// 	mainContent := components.MainContentData{
// 		Breadcrumbs: []components.BreadcrumbItem{
// 			{Name: "Databases", Href: "/"},
// 			{Name: activeSchemaName},
// 			{Name: activeTableName},
// 		},
// 		TableDetails: activeTable,
// 		Columns:      activeColumns,
// 		// Mocked table data for display
// 		TableData: components.TableData{
// 			Headers: []string{"id", "product_name", "status"},
// 			Rows: [][]any{
// 				{1, "Gadget Pro", "active"},
// 				{2, "Widget Max", "inactive"},
// 			},
// 		},
// 	}

// 	return pages.ExplorerPageData{
// 		SidebarSchemas: sidebarSchemas,
// 		MainContent:    mainContent,
// 	}
// }

// func findActiveTable(schemas []domain.Schema, schemaName, tableName string) *domain.Table {
// 	for i := range schemas {
// 		if schemas[i].Name == schemaName {
// 			for j := range schemas[i].Tables {
// 				if schemas[i].Tables[j].Name == tableName {
// 					return &schemas[i].Tables[j]
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }
