package editor

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"northstar/internal/features/editor/components"
	"northstar/internal/features/editor/pages"
	"northstar/internal/features/editor/services"

	"github.com/gorilla/sessions"
	"github.com/starfederation/datastar-go/datastar"
)

type Handlers struct {
	sessionStore  sessions.Store
	editorService *services.EditorService
}

func NewHandlers(sessionStore sessions.Store, editorService *services.EditorService) *Handlers {
	return &Handlers{
		sessionStore:  sessionStore,
		editorService: editorService,
	}
}

func (h *Handlers) EditorPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Just render the shell - SSE will handle state updates
	if err := pages.EditorPage(pages.EditorPageData{}).Render(ctx, w); err != nil {
		slog.Error("Failed to render editor page", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) EditorSSE(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	schema := "public"
	table := "test"
	page := 1

	sse := datastar.NewSSE(w, r)

	tableData, err := h.editorService.GetTableData(ctx, schema, table, page)
	if err != nil {
		if err := sse.ConsoleError(fmt.Errorf("failed to load table data: %w", err)); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	columns := make([]components.Column, len(tableData.Columns))
	for i, col := range tableData.Columns {
		columns[i] = components.Column{
			Title: col.Name,
			Width: "150px",
		}
	}

	// c := components.DataTable(columns, tableData.Rows, tableData.Pagination, fmt.Sprintf("%s.%s", schema, table))
	c := components.DataTableV2(columns, tableData.Rows)

	// c := components.EmptyTableView()

	if err := sse.PatchElementTempl(c); err != nil {
		if err := sse.ConsoleError(err); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

}

// renderState renders the current state to the SSE connection
func (h *Handlers) renderState(ctx context.Context, sse *datastar.ServerSentEventGenerator, state services.EditorState) error {
	// If no table selected, show empty view
	if state.TableName == "" {
		sse.PatchElementTempl(components.EmptyTableView(), datastar.WithSelector("#table-container"))
		return nil
	}

	// Parse table name
	schema, table, err := h.editorService.ParseTableName(state.TableName)
	if err != nil {
		sse.PatchElementTempl(components.ErrorView(fmt.Sprintf("Invalid table name: %v", err)), datastar.WithSelector("#table-container"))
		return nil
	}

	// Load table data
	tableData, err := h.editorService.GetTableData(ctx, schema, table, state.Page)
	if err != nil {
		sse.PatchElementTempl(components.ErrorView(fmt.Sprintf("Failed to load table: %v", err)), datastar.WithSelector("#table-container"))
		return nil
	}

	// Convert domain types to component types
	columns := make([]components.Column, len(tableData.Columns))
	for i, col := range tableData.Columns {
		columns[i] = components.Column{
			Title: col.Name,
			Width: "150px",
		}
	}

	// Render the data table component
	sse.PatchElementTempl(components.DataTable(columns, tableData.Rows, tableData.Pagination, state.TableName), datastar.WithSelector("#table-container"))
	return nil
}

// LoadTable command - updates state only
