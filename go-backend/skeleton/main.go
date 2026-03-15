package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/adapter/handler"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/adapter/memory"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/item/service"
	"github.com/${{ (values.repoUrl | parseRepoUrl).owner }}/${{ values.name }}/internal/server"
)

func main() {
	logger, err := buildLogger(env("LOG_LEVEL", "info"))
	if err != nil {
		panic(err)
	}
	defer logger.Sync() //nolint:errcheck

	cfg := server.Config{
		ServiceName: env("SERVICE_NAME", "${{ values.name }}"),
		Port:        env("PORT", "${{ values.port }}"),
	}

	// Composition root — wire: driven adapter → service → handler
	itemRepo := memory.NewRepository()          // driven adapter  (output port impl)
	itemSvc := service.NewItemService(itemRepo) // app service     (input port impl)
	itemHandler := handler.NewHandler(itemSvc)  // driving adapter (consumes input port)

	srv, r := server.New(cfg, logger)
	r.Route("/api/v1", func(r chi.Router) {
		itemHandler.Register(r)
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		logger.Info("server starting", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil {
			stop()
		}
	}()

	<-ctx.Done()
	logger.Info("shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("shutdown error", zap.Error(err))
		os.Exit(1)
	}
}

func buildLogger(level string) (*zap.Logger, error) {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapLevel)
	return cfg.Build()
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
