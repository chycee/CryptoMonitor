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
	// 초기 로거 (설정 로드 전용)
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	slog.Info("🚀 CryptoMonitor - Starting...")

	// Phase 1.1: 설정 로드
	cfg, err := infra.LoadConfig("configs/config.yaml")
	if err != nil {
		slog.Error("Failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	// 4대 원칙: 설정 기반 로그 레벨 적용
	var level slog.Level
	switch cfg.Logging.Level {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)

	slog.Info("✅ Configuration loaded",
		slog.String("app", cfg.App.Name),
		slog.String("version", cfg.App.Version),
		slog.String("log_level", cfg.Logging.Level),
	)

	// 4대 원칙: Context 기반 생명주기 관리
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 서비스 초기화
	priceService := service.NewPriceService()
	priceService.StartTickerProcessor(ctx)
	slog.InfoContext(ctx, "✅ PriceService initialized with ticker processor")

	// 환율 클라이언트 초기화 및 시작
	exchangeRateClient := infra.NewExchangeRateClientWithConfig(
		priceService.UpdateExchangeRate,
		cfg.API.ExchangeRate.URL,
		cfg.API.ExchangeRate.PollIntervalSec,
	)
	if err := exchangeRateClient.Start(ctx); err != nil {
		slog.Error("Failed to start exchange rate client", slog.Any("error", err))
	}
	defer exchangeRateClient.Stop()
	slog.InfoContext(ctx, "✅ ExchangeRateClient started")

	// Upbit Worker 초기화
	if len(cfg.API.Upbit.Symbols) > 0 {
		upbitWorker := infra.NewUpbitWorker(cfg.API.Upbit.Symbols, priceService.GetTickerChan())
		if err := upbitWorker.Connect(ctx); err != nil {
			slog.Error("Failed to connect Upbit", slog.Any("error", err))
		}
		defer upbitWorker.Disconnect()
		slog.InfoContext(ctx, "✅ UpbitWorker started", slog.Any("symbols", cfg.API.Upbit.Symbols))
	}

	// Bitget Worker 초기화
	if len(cfg.API.Bitget.Symbols) > 0 {
		// Bitget Spot
		bitgetSpotWorker := infra.NewBitgetSpotWorker(cfg.API.Bitget.Symbols, priceService.GetTickerChan())
		if err := bitgetSpotWorker.Connect(ctx); err != nil {
			slog.Error("Failed to connect Bitget Spot", slog.Any("error", err))
		}
		defer bitgetSpotWorker.Disconnect()
		slog.InfoContext(ctx, "✅ BitgetSpotWorker started")

		// Bitget Futures
		bitgetFuturesWorker := infra.NewBitgetFuturesWorker(cfg.API.Bitget.Symbols, priceService.GetTickerChan())
		if err := bitgetFuturesWorker.Connect(ctx); err != nil {
			slog.Error("Failed to connect Bitget Futures", slog.Any("error", err))
		}
		defer bitgetFuturesWorker.Disconnect()
		slog.InfoContext(ctx, "✅ BitgetFuturesWorker started")
	}

	// TODO: UI 초기화 (메인 윈도우 루프)

	slog.InfoContext(ctx, "🚀 Application ready. Press Ctrl+C to exit.")
	<-ctx.Done() // 종료 신호까지 대기

	slog.InfoContext(ctx, "👋 CryptoMonitor - Shutting down gracefully...")
}
