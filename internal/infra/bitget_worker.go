package infra

import (
	"context"
	"crypto_monitor/internal/domain"
)

// BitgetSpotWorker handles Bitget Spot WebSocket connection
type BitgetSpotWorker struct {
	symbols   map[string]string // unified -> instId
	onTicker  func([]*domain.Ticker)
	connected bool
}

// NewBitgetSpotWorker creates a new Bitget Spot worker
func NewBitgetSpotWorker(symbols map[string]string, onTicker func([]*domain.Ticker)) *BitgetSpotWorker {
	return &BitgetSpotWorker{
		symbols:  symbols,
		onTicker: onTicker,
	}
}

// Connect starts the WebSocket connection
func (w *BitgetSpotWorker) Connect(ctx context.Context) error {
	// TODO: Implement using ctx
	return nil
}

// Disconnect closes the connection
func (w *BitgetSpotWorker) Disconnect() {
	// TODO: Implement
}

// BitgetFuturesWorker handles Bitget Futures WebSocket connection
type BitgetFuturesWorker struct {
	symbols   map[string]string
	onTicker  func([]*domain.Ticker)
	connected bool
}

// NewBitgetFuturesWorker creates a new Bitget Futures worker
func NewBitgetFuturesWorker(symbols map[string]string, onTicker func([]*domain.Ticker)) *BitgetFuturesWorker {
	return &BitgetFuturesWorker{
		symbols:  symbols,
		onTicker: onTicker,
	}
}

// Connect starts the WebSocket connection
func (w *BitgetFuturesWorker) Connect(ctx context.Context) error {
	// TODO: Implement using ctx
	return nil
}

// Disconnect closes the connection
func (w *BitgetFuturesWorker) Disconnect() {
	// TODO: Implement
}
