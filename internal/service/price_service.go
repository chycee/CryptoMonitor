package service

import (
	"sync"

	"crypto_monitor/internal/domain"

	"github.com/shopspring/decimal"
)

// PriceService manages the state of all market data
type PriceService struct {
	mu           sync.RWMutex
	marketData   map[string]*domain.MarketData
	exchangeRate decimal.Decimal
}

// NewPriceService creates a new PriceService instance
func NewPriceService() *PriceService {
	return &PriceService{
		marketData:   make(map[string]*domain.MarketData),
		exchangeRate: decimal.Zero,
	}
}

// GetAllData returns all market data
func (s *PriceService) GetAllData() []*domain.MarketData {
	// TODO: Implement
	return nil
}

// GetData returns market data for a specific symbol
func (s *PriceService) GetData(symbol string) *domain.MarketData {
	// TODO: Implement
	return nil
}

// UpdateExchangeRate updates the USD/KRW exchange rate
func (s *PriceService) UpdateExchangeRate(rate decimal.Decimal) {
	// TODO: Implement
}

// UpdateUpbit updates Upbit ticker data
func (s *PriceService) UpdateUpbit(tickers []*domain.Ticker) {
	// TODO: Implement
}

// UpdateBitget updates Bitget ticker data
func (s *PriceService) UpdateBitget(tickers []*domain.Ticker) {
	// TODO: Implement
}

// SetFavorite sets the favorite status for a symbol
func (s *PriceService) SetFavorite(symbol string, isFavorite bool) {
	// TODO: Implement
}
