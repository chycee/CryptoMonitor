package infra

import (
	"context"

	"github.com/shopspring/decimal"
)

// ExchangeRateClient fetches USD/KRW exchange rate
type ExchangeRateClient struct {
	onUpdate func(decimal.Decimal)
	rate     decimal.Decimal
}

// NewExchangeRateClient creates a new exchange rate client
func NewExchangeRateClient(onUpdate func(decimal.Decimal)) *ExchangeRateClient {
	return &ExchangeRateClient{
		onUpdate: onUpdate,
		rate:     decimal.Zero,
	}
}

// Start begins polling for exchange rate updates
func (c *ExchangeRateClient) Start(ctx context.Context) error {
	// TODO: Implement with Upbit (Dunamu) Forex API
	// Endpoint: https://quotation-api-cdn.dunamu.com/v1/forex/recent?codes=FRX.KRWUSD
	// using ctx for cancellation
	return nil
}

// Stop stops the polling
func (c *ExchangeRateClient) Stop() {
	// TODO: Implement
}

// GetRate returns the current exchange rate
func (c *ExchangeRateClient) GetRate() decimal.Decimal {
	return c.rate
}
