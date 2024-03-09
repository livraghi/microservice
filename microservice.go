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

	logDefaultsArgs := []any{
		slog.String("app", cfg.name),
		slog.String("version", cfg.version),
		slog.Int("revision", cfg.revision),
		slog.String("environment", cfg.environment),
		slog.Int("port", int(cfg.port)),
		slog.Duration("gracefulShutdownTimeout", cfg.gracefulShutdownTimeout),
	}

	slog.Info("starting http server", logDefaultsArgs...)

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
		slog.Info("start graceful shutdown", logDefaultsArgs...)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.gracefulShutdownTimeout)
		defer cancel()

		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			slog.Error("graceful shutdown failed", append(logDefaultsArgs, slog.Any("error", err))...)
			return err
		}
		select {
		case <-shutdownCtx.Done():
		}

		slog.Info("graceful shutdown completed", logDefaultsArgs...)
		return nil
	})

	if err := g.Wait(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("http server failed", append(logDefaultsArgs, slog.Any("error", err))...)
	} else {
		slog.Info("http server stopped", logDefaultsArgs...)
	}
}
