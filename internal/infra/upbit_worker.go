package infra

import (
	"context"
	"crypto_monitor/internal/domain"
)

// UpbitWorker handles Upbit WebSocket connection
type UpbitWorker struct {
	symbols   []string
	onTicker  func([]*domain.Ticker)
	connected bool
}

// NewUpbitWorker creates a new Upbit worker
func NewUpbitWorker(symbols []string, onTicker func([]*domain.Ticker)) *UpbitWorker {
	return &UpbitWorker{
		symbols:  symbols,
		onTicker: onTicker,
	}
}

// Connect starts the WebSocket connection
func (w *UpbitWorker) Connect(ctx context.Context) error {
	// TODO: Implement WebSocket connection using context for cancellation
	// Use slog for structured logging
	return nil
}

// Disconnect closes the WebSocket connection
func (w *UpbitWorker) Disconnect() {
	// TODO: Implement
}

// IsConnected returns connection status
func (w *UpbitWorker) IsConnected() bool {
	return w.connected
}
