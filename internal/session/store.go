// /internal/session/store.go

package session

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/delaneyj/toolbelt"
	"github.com/gorilla/sessions"
	"github.com/nats-io/nats.go/jetstream"
)

// StateStore is a generic service for managing user state in a KV store.
// T is the type of the state struct (e.g., ExplorerState).
type StateStore[T any] struct {
	kv    jetstream.KeyValue
	store sessions.Store
}

func NewStateStore[T any](kv jetstream.KeyValue, store sessions.Store) *StateStore[T] {
	return &StateStore[T]{
		kv:    kv,
		store: store,
	}
}

// GetSessionID is a generic helper that belongs here.
func (s *StateStore[T]) GetSessionID(r *http.Request, w http.ResponseWriter) (string, error) {
	sess, err := s.store.Get(r, "connections")
	if err != nil {
		return "", fmt.Errorf("failed to get session: %w", err)
	}

	id, ok := sess.Values["id"].(string)
	if !ok || id == "" {
		id = toolbelt.NextEncodedID()
		sess.Values["id"] = id
		if err := sess.Save(r, w); err != nil {
			return "", fmt.Errorf("failed to save session: %w", err)
		}
	}
	return id, nil
}

// Get retrieves and unmarshals the state for a given session ID.
// If the state doesn't exist, it returns a new, zero-value instance of T.
func (s *StateStore[T]) Get(ctx context.Context, sessionID string) (*T, error) {
	entry, err := s.kv.Get(ctx, sessionID)
	state := new(T) // Create a new pointer to the generic type T

	if err != nil {
		if err == jetstream.ErrKeyNotFound {
			return state, nil // Return the empty state, not an error
		}
		return nil, fmt.Errorf("failed to get kv entry: %w", err)
	}

	if err := json.Unmarshal(entry.Value(), state); err != nil {
		return state, fmt.Errorf("failed to unmarshal state: %w", err)
	}

	return state, nil
}

// Save marshals and saves the state for a given session ID.
func (s *StateStore[T]) Save(ctx context.Context, sessionID string, state *T) error {
	b, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("failed to marshal state: %w", err)
	}

	if _, err := s.kv.Put(ctx, sessionID, b); err != nil {
		return fmt.Errorf("failed to save state to kv: %w", err)
	}
	return nil
}
