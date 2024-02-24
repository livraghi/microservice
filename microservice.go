package microservice

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartDefaultHttpServer(ctx context.Context, opts ...ServerOption) {
	cfg := newServerConfig(opts...)

	mainCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		BaseContext: func(_ net.Listener) context.Context {
			return mainCtx
		},
	}

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		return srv.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		slog.Info("start graceful shutdown", "timeout", cfg.gracefulShutdownTimeout)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.gracefulShutdownTimeout)
		defer cancel()

		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			slog.Error("graceful shutdown failed", "timeout", cfg.gracefulShutdownTimeout, "error", err.Error())
			return err
		}
		select {
		case <-shutdownCtx.Done():
		}

		slog.Info("graceful shutdown completed", "timeout", cfg.gracefulShutdownTimeout)
		return nil
	})

	if err := g.Wait(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("http server failed", "error", err.Error())
	} else {
		slog.Info("http server stopped")
	}
}
