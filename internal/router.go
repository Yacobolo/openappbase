package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"northstar/config"
	"northstar/internal/features/editor"
	"northstar/internal/features/explorer"
	"northstar/internal/store"
	"northstar/web/resources"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/starfederation/datastar-go/datastar"
)

func SetupRoutes(ctx context.Context, router chi.Router, sessionStore *sessions.CookieStore, js jetstream.JetStream, pool *pgxpool.Pool, q *store.Queries) (err error) {

	if config.Global.Environment == config.Dev {
		setupReload(ctx, router)
	}

	router.Handle("/static/*", resources.Handler())

	// Create NATS KV bucket for editor sessions
	_, err = js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket:  "editor_sessions",
		TTL:     24 * time.Hour,
		Storage: jetstream.MemoryStorage,
	})
	if err != nil && err != jetstream.ErrBucketExists {
		return fmt.Errorf("error creating editor_sessions KV bucket: %w", err)
	}

	if err := errors.Join(
		explorer.SetupRoutes(router, sessionStore, js, q),
		editor.SetupRoutes(router, sessionStore, js, pool, q),
	); err != nil {
		return fmt.Errorf("error setting up routes: %w", err)
	}

	return nil
}

func setupReload(ctx context.Context, router chi.Router) {
	reloadChan := make(chan struct{}, 1)
	var hotReloadOnce sync.Once
	router.Get("/reload", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)
		reload := func() { sse.ExecuteScript("window.location.reload()") }
		hotReloadOnce.Do(reload)
		select {
		case <-reloadChan:
			reload()
		case <-r.Context().Done():
		}
	})

	router.Get("/hotreload", func(w http.ResponseWriter, r *http.Request) {
		select {
		case reloadChan <- struct{}{}:
		default:
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

}
