// cmd/web/main.go
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"northstar/config"
	internal "northstar/internal"
	"northstar/internal/nats"
	"northstar/internal/store"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"github.com/nats-io/nats.go/jetstream"
	"golang.org/x/sync/errgroup"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: config.Global.LogLevel,
	}))
	slog.SetDefault(logger)

	if err := run(ctx); err != nil && err != http.ErrServerClosed {
		slog.Error("error running server", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {

	database, err := store.InitDB(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	defer func() {
		if closeErr := database.Close(); closeErr != nil {
			slog.Error("error closing database", "error", closeErr)
		}
	}()

	queries := store.New(database)

	addr := fmt.Sprintf("%s:%s", config.Global.Host, config.Global.Port)
	slog.Info("server started", "addr", addr)
	defer slog.Info("server shutdown complete")

	eg, egctx := errgroup.WithContext(ctx)

	router := chi.NewMux()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	sessionStore := sessions.NewCookieStore([]byte(config.Global.SessionSecret))
	sessionStore.MaxAge(86400 * 30)
	sessionStore.Options.Path = "/"
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.Secure = false
	sessionStore.Options.SameSite = http.SameSiteLaxMode

	ns, err := nats.SetupNATS(ctx)
	if err != nil {
		return err
	}

	nc, err := ns.Client()
	if err != nil {
		return fmt.Errorf("error creating nats client: %w", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return fmt.Errorf("error creating jetstream client: %w", err)
	}

	if err := internal.SetupRoutes(egctx, router, sessionStore, js, database.DB, queries, nc); err != nil {
		return fmt.Errorf("error setting up routes: %w", err)
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
		BaseContext: func(l net.Listener) context.Context {
			return egctx
		},
		ErrorLog: slog.NewLogLogger(
			slog.Default().Handler(),
			slog.LevelError,
		),
	}

	eg.Go(func() error {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("server error: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		<-egctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		slog.Debug("shutting down server...")

		if err := srv.Shutdown(shutdownCtx); err != nil {
			slog.Error("error during shutdown", "error", err)
			return err
		}

		return nil
	})

	return eg.Wait()
}
