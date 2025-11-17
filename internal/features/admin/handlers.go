package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"northstar/internal/features/admin/components"
	"northstar/internal/features/admin/pages"
	"northstar/internal/features/admin/services"
	commoncomponents "northstar/internal/features/common/components"
	"northstar/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/nats-io/nats.go"
	"github.com/starfederation/datastar-go/datastar"
)

type Handlers struct {
	sessionStore      sessions.Store
	connectionService *services.ConnectionService
	nc                *nats.Conn
}

func NewHandlers(connectionService *services.ConnectionService, nc *nats.Conn) *Handlers {
	return &Handlers{
		connectionService: connectionService,
		nc:                nc,
	}
}

func (h *Handlers) AdminPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Just render the shell - SSE will handle state updates
	if err := pages.AdminPage().Render(ctx, w); err != nil {
		slog.Error("Failed to render editor page", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// SSE Handlers for Datastar

// AdminConnectionsSSE maintains a long-lived SSE connection watching for updates
func (h *Handlers) AdminConnectionsSSE(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sse := datastar.NewSSE(w, r)

	slog.Info("SSE connection established for admin connections")

	// Helper function to fetch and render connections
	// Takes a context parameter to allow using different contexts (request ctx vs background ctx)
	renderConnections := func(queryCtx context.Context) error {
		slog.Debug("renderConnections called")
		connections, err := h.connectionService.ListConnections(queryCtx)
		if err != nil {
			slog.Error("Failed to list connections in renderConnections", "error", err)
			return err
		}

		slog.Debug("Retrieved connections", "count", len(connections))

		// Convert to component connections
		componentConns := make([]components.Connection, len(connections))
		for i, conn := range connections {
			componentConns[i] = toComponentConnection(conn)
		}

		// Render the connections tab content (includes modal)
		if err := sse.PatchElementTempl(components.ConnectionsTabContent(componentConns)); err != nil {
			slog.Error("Failed to patch element in renderConnections", "error", err)
			return err
		}

		slog.Debug("Successfully patched connections UI")
		return nil
	}

	// Render initial state using request context
	if err := renderConnections(ctx); err != nil {
		sse.ConsoleError(fmt.Errorf("failed to load initial connections: %w", err))
		return
	}

	// Subscribe to NATS updates
	sub, err := h.nc.Subscribe("admin.connections.update", func(msg *nats.Msg) {
		slog.Info("NATS callback triggered - connection update received")
		// Use background context for async DB queries, not the request context
		if err := renderConnections(context.Background()); err != nil {
			slog.Error("Failed to render connections in NATS callback", "error", err)
			sse.ConsoleError(fmt.Errorf("failed to update connections: %w", err))
		} else {
			slog.Info("Successfully updated connections via NATS callback")
		}
	})
	if err != nil {
		slog.Error("Failed to subscribe to connection updates", "error", err)
		sse.ConsoleError(fmt.Errorf("failed to subscribe to updates: %w", err))
		return
	}
	defer sub.Unsubscribe()

	slog.Info("Subscribed to NATS topic: admin.connections.update")

	// Keep connection alive until client disconnects
	<-ctx.Done()
	slog.Info("SSE connection closed for admin connections")
}

// CreateConnectionSSE creates a new connection via SSE

// TestConnection tests an existing connection
func (h *Handlers) TestConnection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid connection ID", http.StatusBadRequest)
		return
	}

	sse := datastar.NewSSE(w, r)

	if err := h.connectionService.TestConnection(ctx, id); err != nil {
		msg := fmt.Sprintf("Connection test failed: %v", err)
		slog.Error(msg, "id", id)
		sse.ConsoleLog(msg)

		sse.PatchElementTempl(
			commoncomponents.Toast(msg, commoncomponents.ToastError),
			datastar.WithSelectorID("toast-container"),
			datastar.WithModeAppend(),
		)
		return
	}

	msg := "Connection test successful!"
	slog.Info(msg, "id", id)
	sse.PatchElementTempl(
		commoncomponents.Toast(msg, commoncomponents.ToastSuccess),
		datastar.WithSelectorID("toast-container"),
		datastar.WithModeAppend(),
	)
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

	// Delete the connection
	if err := h.connectionService.DeleteConnection(ctx, id); err != nil {
		slog.Error("Failed to delete connection", "id", id, "error", err)
		http.Error(w, fmt.Sprintf("Failed to delete connection: %v", err), http.StatusInternalServerError)
		return
	}

	// Notify all subscribers
	h.publishConnectionUpdate()

	w.WriteHeader(http.StatusNoContent)
}

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

	// Create the connection
	connection, err := h.connectionService.CreateConnection(ctx, input)
	if err != nil {
		slog.Error("Failed to create connection", "error", err)
		http.Error(w, fmt.Sprintf("Failed to create connection: %v", err), http.StatusInternalServerError)
		return
	}

	// Notify all subscribers
	h.publishConnectionUpdate()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(connection); err != nil {
		slog.Error("Failed to encode connection", "error", err)
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

// Helper Functions

// publishConnectionUpdate notifies all subscribers that connections have changed
func (h *Handlers) publishConnectionUpdate() {
	slog.Info("Publishing connection update to NATS", "topic", "admin.connections.update")
	// NATS is required for app startup - no nil check needed
	if err := h.nc.Publish("admin.connections.update", []byte("{}")); err != nil {
		slog.Error("Failed to publish connection update", "error", err)
		// Best-effort notification - don't fail the HTTP request
	} else {
		slog.Info("Successfully published connection update to NATS")
	}
}

// toComponentConnection converts store.Connection to components.Connection
func toComponentConnection(conn store.Connection) components.Connection {
	return components.Connection{
		ID:          conn.ID,
		Name:        conn.Name,
		Host:        conn.Host,
		Port:        conn.Port,
		Database:    conn.Database,
		Username:    conn.Username,
		SSLMode:     conn.SslMode,
		IsActive:    conn.IsActive == 1,
		Environment: "", // Not in DB schema yet
		UpdatedAt:   formatTime(conn.UpdatedAt),
	}
}

// formatTime formats a time.Time to a human-readable string
func formatTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return "just now"
	case diff < time.Hour:
		mins := int(diff.Minutes())
		if mins == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", mins)
	case diff < 24*time.Hour:
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	case diff < 7*24*time.Hour:
		days := int(diff.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	default:
		return t.Format("Jan 2, 2006")
	}
}
