package connections

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"northstar/internal/features/connections/services"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

type Handlers struct {
	sessionStore      sessions.Store
	connectionService *services.ConnectionService
}

func NewHandlers(sessionStore sessions.Store, connectionService *services.ConnectionService) *Handlers {
	return &Handlers{
		sessionStore:      sessionStore,
		connectionService: connectionService,
	}
}

// ListConnections returns all database connections
func (h *Handlers) ListConnections(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	connections, err := h.connectionService.ListConnections(ctx)
	if err != nil {
		slog.Error("Failed to list connections", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(connections); err != nil {
		slog.Error("Failed to encode connections", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// GetConnection returns a single connection by ID
func (h *Handlers) GetConnection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid connection ID", http.StatusBadRequest)
		return
	}

	connection, err := h.connectionService.GetConnection(ctx, id)
	if err != nil {
		slog.Error("Failed to get connection", "id", id, "error", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(connection); err != nil {
		slog.Error("Failed to encode connection", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// CreateConnection creates a new database connection
func (h *Handlers) CreateConnection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input services.CreateConnectionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if input.Name == "" || input.Host == "" || input.Port == 0 || input.Database == "" || input.Username == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	connection, err := h.connectionService.CreateConnection(ctx, input)
	if err != nil {
		slog.Error("Failed to create connection", "error", err)
		http.Error(w, fmt.Sprintf("Failed to create connection: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(connection); err != nil {
		slog.Error("Failed to encode connection", "error", err)
	}
}

// UpdateConnection updates an existing connection
func (h *Handlers) UpdateConnection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid connection ID", http.StatusBadRequest)
		return
	}

	var input services.UpdateConnectionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the ID from the URL param
	input.ID = id

	connection, err := h.connectionService.UpdateConnection(ctx, input)
	if err != nil {
		slog.Error("Failed to update connection", "id", id, "error", err)
		http.Error(w, fmt.Sprintf("Failed to update connection: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(connection); err != nil {
		slog.Error("Failed to encode connection", "error", err)
	}
}

// DeleteConnection deletes a connection
func (h *Handlers) DeleteConnection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid connection ID", http.StatusBadRequest)
		return
	}

	if err := h.connectionService.DeleteConnection(ctx, id); err != nil {
		slog.Error("Failed to delete connection", "id", id, "error", err)
		http.Error(w, fmt.Sprintf("Failed to delete connection: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// TestConnectionHandler tests a connection and returns JSON result
func (h *Handlers) TestConnectionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input services.CreateConnectionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Test the connection using TestConnectionDirect
	if err := h.connectionService.TestConnectionDirect(ctx, input); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": fmt.Sprintf("Connection test failed: %v", err),
		})
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Connection successful!",
	})
}

// TestConnectionByID tests an existing saved connection
func (h *Handlers) TestConnectionByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid connection ID", http.StatusBadRequest)
		return
	}

	// Test the connection
	if err := h.connectionService.TestConnection(ctx, id); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": fmt.Sprintf("Connection test failed: %v", err),
		})
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Connection successful!",
	})
}
