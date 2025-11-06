package editor

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

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

	// Get table name and page from query parameters
	tableName := r.URL.Query().Get("table")
	pageStr := r.URL.Query().Get("page")

	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	// If no table specified, show empty editor with input
	if tableName == "" {
		if err := pages.EditorPage(pages.EditorPageData{}).Render(ctx, w); err != nil {
			slog.Error("Failed to render editor page", "error", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	// Parse and load table data
	schema, table, err := h.editorService.ParseTableName(tableName)
	if err != nil {
		if err := pages.EditorPage(pages.EditorPageData{
			TableName: tableName,
			Error:     err.Error(),
		}).Render(ctx, w); err != nil {
			slog.Error("Failed to render editor page with error", "error", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	tableData, err := h.editorService.GetTableData(ctx, schema, table, page)
	if err != nil {
		if err := pages.EditorPage(pages.EditorPageData{
			TableName: tableName,
			Error:     fmt.Sprintf("Failed to load table data: %v", err),
		}).Render(ctx, w); err != nil {
			slog.Error("Failed to render editor page with error", "error", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	// Convert domain types to component types
	columns := make([]components.Column, len(tableData.Columns))
	for i, col := range tableData.Columns {
		columns[i] = components.Column{
			Title: col.Name,
			Width: "150px", // Default width
		}
	}

	pageData := pages.EditorPageData{
		TableName:  tableName,
		Columns:    columns,
		Rows:       tableData.Rows,
		Pagination: tableData.Pagination,
	}

	if err := pages.EditorPage(pageData).Render(ctx, w); err != nil {
		slog.Error("Failed to render editor page", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) LoadTableData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sse := datastar.NewSSE(w, r)

	tableName := r.URL.Query().Get("table")
	pageStr := r.URL.Query().Get("page")

	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if tableName == "" {
		sse.ConsoleError(fmt.Errorf("table name is required"))
		return
	}

	schema, table, err := h.editorService.ParseTableName(tableName)
	if err != nil {
		sse.ConsoleError(fmt.Errorf("invalid table name: %v", err))
		return
	}

	tableData, err := h.editorService.GetTableData(ctx, schema, table, page)
	if err != nil {
		sse.ConsoleError(fmt.Errorf("failed to load table: %v", err))
		return
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
	sse.PatchElementTempl(components.DataTable(columns, tableData.Rows, tableData.Pagination, tableName), datastar.WithSelector("#table-container"))
}

func (h *Handlers) ChangePage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sse := datastar.NewSSE(w, r)

	tableName := r.URL.Query().Get("table")
	pageStr := r.URL.Query().Get("page")

	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if tableName == "" {
		sse.ConsoleError(fmt.Errorf("table name is required"))
		return
	}

	schema, table, err := h.editorService.ParseTableName(tableName)
	if err != nil {
		sse.ConsoleError(fmt.Errorf("invalid table name: %v", err))
		return
	}

	tableData, err := h.editorService.GetTableData(ctx, schema, table, page)
	if err != nil {
		sse.ConsoleError(fmt.Errorf("failed to load page: %v", err))
		return
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
	sse.PatchElementTempl(components.DataTable(columns, tableData.Rows, tableData.Pagination, tableName), datastar.WithSelector("#table-container"))
}
