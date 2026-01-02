package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"crypto_monitor/internal/infra"
	"crypto_monitor/internal/service"
)

func main() {
	// 4대 원칙: 구조화된 로깅 (slog) 초기화
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("🚀 CryptoMonitor - Starting...")

	// Phase 1.1: Load Configuration
	cfg, err := infra.LoadConfig("configs/config.yaml")
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}
	slog.Info("✅ Configuration loaded", "app", cfg.App.Name, "v", cfg.App.Version)

	// 4대 원칙: Context 기반 생명주기 관리
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize services
	priceService := service.NewPriceService()
	slog.InfoContext(ctx, "✅ PriceService initialized", "service", priceService)

	// TODO: Initialize infrastructure (WebSocket workers) with ctx

	// TODO: Initialize UI (Main window loop)

	slog.InfoContext(ctx, "Application ready. Press Ctrl+C to exit.")
	<-ctx.Done() // Block until shutdown signal
	slog.InfoContext(ctx, "👋 CryptoMonitor - Shutting down gracefully...")
}
